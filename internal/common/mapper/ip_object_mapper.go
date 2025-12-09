package mapper

import (
	"server/internal/common/dto"
	"server/internal/common/entity"
)

func MapIpObjectEntitiesToIpObjectDtos(entities []*entity.IpObject) []*dto.IpObjectDto {
	dtos := make([]*dto.IpObjectDto, 0)

	for _, e := range entities {
		dtos = append(dtos, MapIpObjectEntityToIpObjectDto(e))
	}

	return dtos
}

func MapIpObjectEntityToIpObjectDto(entity *entity.IpObject) *dto.IpObjectDto {
	return &dto.IpObjectDto{
		Id:           entity.Id,
		UserId:       entity.UserId,
		Title:        entity.Title,
		Description:  entity.Description,
		Jurisdiction: entity.Jurisdiction,
		PatentType:   entity.PatentType,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
