package service

import (
	"absensi-api/dto"
	"absensi-api/internal/config"
	"absensi-api/internal/domain"
	"absensi-api/pkg/utils"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type attendanceService struct {
	repository                  domain.AttendanceRepository
	attendanceHistoryRepository domain.AttendanceHistoryRepository
	employeeRepository          domain.EmployeeRepository
}

func NewAttendance(
	repository domain.AttendanceRepository,
	attendanceHistoryRepository domain.AttendanceHistoryRepository,
	employeeRepository domain.EmployeeRepository,
) domain.AttendanceService {
	return &attendanceService{
		repository:                  repository,
		attendanceHistoryRepository: attendanceHistoryRepository,
		employeeRepository:          employeeRepository,
	}
}

var configuration = config.Load()

func (s *attendanceService) GetAllAttendance(ctx context.Context) ([]domain.AttendanceResponse, error) {
	data := []domain.AttendanceResponse{}

	result, err := s.repository.FindAll(ctx)
	if err != nil {
		return []domain.AttendanceResponse{}, err
	}

	for _, v := range result {
		data = append(data, domain.AttendanceResponse{
			ID:           v.ID,
			AttendanceID: v.AttendanceID,
			Employee: domain.EmployeeResponse{
				ID:         v.Employee.ID,
				EmployeeID: v.EmployeeID,
				Name:       v.Employee.Name,
				Address:    v.Employee.Address,
			},
			ClockIn:   v.ClockIn,
			ClockOut:  v.ClockOut,
			CreatedAt: v.CreatedAt,
		})
	}

	return data, err
}

func (s *attendanceService) FilterAttendance(ctx context.Context, departement_id *int, date *string) ([]domain.AttendanceResponse, error) {
	data := []domain.AttendanceResponse{}

	result, err := s.repository.FindByFilter(ctx, departement_id, date)
	if err != nil {
		return []domain.AttendanceResponse{}, err
	}

	for _, v := range result {
		data = append(data, domain.AttendanceResponse{
			ID:           v.ID,
			AttendanceID: v.AttendanceID,
			Employee: domain.EmployeeResponse{
				ID:          v.Employee.ID,
				EmployeeID:  v.EmployeeID,
				Name:        v.Employee.Name,
				Address:     v.Employee.Address,
				Departement: v.Employee.Departement,
			},
			ClockIn:   v.ClockIn,
			ClockOut:  v.ClockOut,
			CreatedAt: v.CreatedAt,
		})
	}

	return data, nil
}

func (s *attendanceService) GetAttendanceHistory(ctx context.Context, attendance_id string) ([]domain.AttendanceHistoryResponse, error) {
	data := []domain.AttendanceHistoryResponse{}

	result, err := s.attendanceHistoryRepository.FindByAttendanceID(ctx, attendance_id)
	if err != nil {
		return data, err
	}

	for _, v := range result {
		data = append(data, domain.AttendanceHistoryResponse{
			ID:             v.ID,
			Employee:       v.Employee,
			Attendance:     v.Attendance,
			DateAttendance: v.DateAttendance,
			AttendanceType: v.AttendanceType,
			Description:    v.Description,
		})
	}

	return data, nil
}

func (s *attendanceService) ClockIn(ctx context.Context, req dto.AttendanceRequest) error {
	var attendanceID string

	// Find Employee
	employee, err := s.employeeRepository.FindByEmployeeID(ctx, req.EmployeeID)
	if err != nil {
		return errors.New("employee not found")
	}

	// Find Clock IN
	attendance, err := s.repository.FindAttendanceToday(ctx, req.EmployeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create Attendance
			randomAttendanceID, err := utils.RandomNumber(8)
			attendanceID = randomAttendanceID
			if err != nil {
				return err
			}

			data := domain.Attendance{
				AttendanceID: randomAttendanceID,
				EmployeeID:   employee.EmployeeID,
				ClockIn:      time.Now(),
			}

			if err := s.repository.Save(ctx, data); err != nil {
				return err
			}
		}
	} else {
		attendanceID = attendance.AttendanceID
	}

	// Find History Check IN
	history, err := s.attendanceHistoryRepository.FindLastByType(ctx, attendance.AttendanceID, req.EmployeeID)
	if err == nil && history.AttendanceType == 1 {
		return errors.New("you have already clocked in")
	}

	// Check Max Clock In Time
	loc, _ := time.LoadLocation("Asia/Jakarta")
	today := time.Now().In(loc).Format("2006-01-02")
	maxTimeStr := today + " " + employee.Departement.MaxClockInTime.Format("15:04")
	maxTime, err := time.ParseInLocation("2006-01-02 15:04", maxTimeStr, loc)
	if err != nil {
		return err
	}

	if time.Now().After(maxTime) {
		return errors.New("max time error")
	}

	// Create Attendance History
	attendanceHistory := domain.AttendanceHistory{
		EmployeeID:     employee.EmployeeID,
		AttendanceID:   attendanceID,
		DateAttendance: time.Now(),
		AttendanceType: 1,
		Description:    req.Description,
	}

	if err := s.attendanceHistoryRepository.Save(ctx, attendanceHistory); err != nil {
		return err
	}

	return nil
}

func (s *attendanceService) ClockOut(ctx context.Context, req dto.AttendanceRequest) error {
	// Find Employee
	employee, err := s.employeeRepository.FindByEmployeeID(ctx, req.EmployeeID)
	if err != nil {
		return errors.New("employee not found")
	}

	// Find Attendance
	attendance, _ := s.repository.FindAttendanceToday(ctx, req.EmployeeID)

	// Check Max Clock In Out
	loc, _ := time.LoadLocation(configuration.Server.Timezone)
	today := time.Now().In(loc).Format("2006-01-02")
	maxTimeStr := today + " " + employee.Departement.MaxClockInOut.Format("15:04")
	maxTime, err := time.ParseInLocation("2006-01-02 15:04", maxTimeStr, loc)
	if err != nil {
		return err
	}

	if time.Now().After(maxTime) {
		return errors.New("max time error")
	}

	// Find Clock Out
	history, err := s.attendanceHistoryRepository.FindLastByType(ctx, attendance.AttendanceID, req.EmployeeID)
	if err == nil && history.AttendanceType == 0 {
		return errors.New("you have already clocked out")
	}

	// Update Attendance
	now := time.Now()
	data := domain.Attendance{
		ID:       attendance.ID,
		ClockOut: &now,
	}

	if err := s.repository.Update(ctx, data); err != nil {
		return err
	}

	// Create Attendance History
	attendanceHistory := domain.AttendanceHistory{
		EmployeeID:     employee.EmployeeID,
		AttendanceID:   attendance.AttendanceID,
		DateAttendance: time.Now(),
		AttendanceType: 0,
		Description:    req.Description,
	}

	if err := s.attendanceHistoryRepository.Save(ctx, attendanceHistory); err != nil {
		return err
	}

	return nil
}
