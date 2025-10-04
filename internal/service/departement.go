package service

import (
	"absensi-api/dto"
	"absensi-api/internal/domain"
	"context"
	"time"
)

type departementService struct {
	repository domain.DepartementRepository
}

func NewDepartement(repository domain.DepartementRepository) domain.DepartementService {
	return &departementService{
		repository: repository,
	}
}

func (s *departementService) GetAllDepartement(ctx context.Context) ([]domain.Departement, error) {
	data := []domain.Departement{}

	result, err := s.repository.FindAll(ctx)
	if err != nil {
		return []domain.Departement{}, err
	}

	for _, v := range result {
		data = append(data, domain.Departement{
			ID:              v.ID,
			DepartementName: v.DepartementName,
			MaxClockInTime:  v.MaxClockInTime,
			MaxClockInOut:   v.MaxClockInOut,
		})
	}
	return data, nil
}

func (s *departementService) DetailDepartement(ctx context.Context, id int) (domain.Departement, error) {
	result, err := s.repository.Find(ctx, id)
	if err != nil {
		return domain.Departement{}, err
	}
	return domain.Departement{
		ID:              result.ID,
		DepartementName: result.DepartementName,
		MaxClockInTime:  result.MaxClockInTime,
		MaxClockInOut:   result.MaxClockInOut,
	}, nil
}

func (s *departementService) CreateDepartement(ctx context.Context, req dto.CreateDepartementRequest) error {
	// Parse Time
	maxClockInTime, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+req.MaxClockInTime)
	if err != nil {
		return err
	}
	maxClockInOut, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+req.MaxClockInOut)
	if err != nil {
		return err
	}

	data := domain.Departement{
		DepartementName: req.DepartementName,
		MaxClockInTime:  maxClockInTime,
		MaxClockInOut:   maxClockInOut,
	}

	return s.repository.Save(ctx, data)
}

func (s *departementService) UpdateDepartement(ctx context.Context, req dto.UpdateDepartementRequest) error {
	// Parse Time
	maxClockInTime, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+req.MaxClockInTime)
	if err != nil {
		return err
	}
	maxClockInOut, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+req.MaxClockInOut)
	if err != nil {
		return err
	}

	data := domain.Departement{
		ID:              req.ID,
		DepartementName: req.DepartementName,
		MaxClockInTime:  maxClockInTime,
		MaxClockInOut:   maxClockInOut,
	}

	return s.repository.Update(ctx, data)
}

func (s *departementService) DeleteDepartement(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
