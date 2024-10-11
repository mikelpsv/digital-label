package usecase

import (
	"github.com/mikelpsv/digital-label/pkg/model"
	"github.com/mikelpsv/digital-label/pkg/repositories"
)

type Service struct {
	repository *repositories.ServiceRepository
}

func NewService(repository *repositories.ServiceRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetLink(keyLink string) (*model.LinkData, error) {
	res := model.LinkData{}
	data, err := s.repository.GetLink(keyLink)
	if err != nil {
		return nil, err
	}
	res.FromDbo(data)
	return &res, nil
}

func (s *Service) WriteData(data *model.LinkData) error {
	return s.repository.WriteData(data.ToDbo())
}
