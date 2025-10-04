package service

import (
	"absensi-api/dto"
	"absensi-api/internal/domain"
	"context"
)

type employeeService struct {
	repository domain.EmployeeRepository
}

func NewEmployee(repository domain.EmployeeRepository) domain.EmployeeService {
	return &employeeService{
		repository: repository,
	}
}

func (s *employeeService) GetAllEmployee(ctx context.Context) ([]domain.EmployeeResponse, error) {
	data := []domain.EmployeeResponse{}

	result, err := s.repository.FindAll(ctx)
	if err != nil {
		return []domain.EmployeeResponse{}, err
	}

	for _, v := range result {
		data = append(data, domain.EmployeeResponse{
			ID:          v.ID,
			EmployeeID:  v.EmployeeID,
			Departement: v.Departement,
			Name:        v.Name,
			Address:     v.Address,
		})
	}
	return data, nil
}

func (s *employeeService) DetailEmployee(ctx context.Context, id int) (domain.EmployeeResponse, error) {
	result, err := s.repository.Find(ctx, id)
	if err != nil {
		return domain.EmployeeResponse{}, err
	}
	return domain.EmployeeResponse{
		ID:          result.ID,
		EmployeeID:  result.EmployeeID,
		Departement: result.Departement,
		Name:        result.Name,
		Address:     result.Address,
	}, nil
}

func (s *employeeService) CreateEmployee(ctx context.Context, req dto.CreateEmployeeRequest) error {
	data := domain.Employee{
		EmployeeID:    req.EmployeeID,
		DepartementID: req.DepartementID,
		Name:          req.Name,
		Address:       req.Address,
	}

	return s.repository.Save(ctx, data)
}

func (s *employeeService) UpdateEmployee(ctx context.Context, req dto.UpdateEmployeeRequest) error {
	data := domain.Employee{
		ID:            req.ID,
		EmployeeID:    req.EmployeeID,
		DepartementID: req.DepartementID,
		Name:          req.Name,
		Address:       req.Address,
	}

	return s.repository.Update(ctx, data)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
