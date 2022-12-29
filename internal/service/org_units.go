package service

import (
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/storage"
)

type OrgUnits interface {
	CreateOrgUnit(unit *entities.OrgUnit) error
	GetOrgUnits() (*[]string, error)
	DeleteOrgUnit(name string) error
	UpdateOrgUnit(unit *entities.OrgUnit) error
}

type OrgUnitsService struct {
	repo storage.OrgUnits
}

func NewOrgUnitsService(repo storage.OrgUnits) *OrgUnitsService {
	return &OrgUnitsService{repo: repo}
}

func (s *OrgUnitsService) CreateOrgUnit(unit *entities.OrgUnit) error {
	return s.repo.CreateOrgUnit(unit)
}

func (s *OrgUnitsService) GetOrgUnits() (*[]string, error) {
	return s.repo.GetOrgUnits()
}

func (s *OrgUnitsService) DeleteOrgUnit(name string) error {
	return s.repo.DeleteOrgUnit(name)
}

func (s *OrgUnitsService) UpdateOrgUnit(unit *entities.OrgUnit) error {
	return s.repo.UpdateOrgUnit(unit)
}
