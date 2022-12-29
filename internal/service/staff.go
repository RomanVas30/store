package service

import (
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/storage"
)

type Staff interface {
	CreateStaffer(staffer *entities.Staffer) error
	GetStaff() (*[]entities.ShortStaffer, error)
	DeleteStaffer(snils string) error
	SearchStaff(fio string, snils string) (*[]entities.Staffer, error)
	UpdateStaffer(staffer *entities.Staffer, addPosts []entities.StafferPost, deletePosts []entities.StafferPost) error
}

type StaffService struct {
	repo storage.Staff
}

func NewStaffService(repo storage.Staff) *StaffService {
	return &StaffService{repo: repo}
}

func (s *StaffService) CreateStaffer(staffer *entities.Staffer) error {
	return s.repo.CreateStaffer(staffer)
}

func (s *StaffService) GetStaff() (*[]entities.ShortStaffer, error) {
	return s.repo.GetStaff()
}

func (s *StaffService) DeleteStaffer(snils string) error {
	return s.repo.DeleteStaffer(snils)
}

func (s *StaffService) SearchStaff(fio string, snils string) (*[]entities.Staffer, error) {
	return s.repo.SearchStaff(fio, snils)
}

func (s *StaffService) UpdateStaffer(
	staffer *entities.Staffer,
	addPosts []entities.StafferPost,
	deletePosts []entities.StafferPost,
) error {
	return s.repo.UpdateStaffer(staffer, addPosts, deletePosts)
}
