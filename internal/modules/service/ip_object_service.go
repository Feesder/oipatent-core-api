package service

import (
	"server/internal/common/dto"
	"server/internal/common/entity"
	"server/internal/common/mapper"
	"server/internal/config"
	"server/internal/modules/repository"
)

type IpObjectService struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewIpOjbectService(repository *repository.Repository, cfg *config.Config) *IpObjectService {
	return &IpObjectService{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *IpObjectService) CreateIpObject(ipObject dto.CreateIpObjectDto) (*dto.IpObjectDto, error) {
	ipObjectEntity := &entity.IpObject{
		UserId:       ipObject.UserId,
		Title:        ipObject.Title,
		Description:  ipObject.Description,
		Jurisdiction: ipObject.Jurisdiction,
		PatentType:   ipObject.PatentType,
	}

	if err := s.repository.CreateIpObject(ipObjectEntity); err != nil {
		return nil, err
	}

	return mapper.MapIpObjectEntityToIpObjectDto(ipObjectEntity), nil
}

func (s *IpObjectService) GetIpObjectsByUserId(userId string) ([]*dto.IpObjectDto, error) {
	ipObjectEntities, err := s.repository.GetIpObjectsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return mapper.MapIpObjectEntitiesToIpObjectDtos(ipObjectEntities), nil
}
