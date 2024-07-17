package services

import (
	"multipleParam_git/models"
	"multipleParam_git/repositories"
)

type ServicePort interface {
	GetUniversalInfoServices2(catalog []string) ([]models.UniversalInfo, error)
}

type serviceAdapter struct {
	r repositories.RepositoryPort
}

func NewServiceAdapter(r repositories.RepositoryPort) ServicePort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) GetUniversalInfoServices2(catalog []string) ([]models.UniversalInfo, error) {
	data, err := s.r.GetUniversalInfoRepositories2(catalog)
	if err != nil {
		return nil, err
	}
	return data, nil
}
