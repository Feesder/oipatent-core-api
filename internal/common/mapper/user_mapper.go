package mapper

import (
	"server/internal/common/dto"
	"server/internal/common/entity"
)

func MapUserEntityToUserDto(entity *entity.User) *dto.UserDto {
	return &dto.UserDto{
		Id:        entity.Id,
		Username:  entity.Username,
		Firstname: entity.Firstname,
		Lastname:  entity.Lastname,
		Email:     entity.Email,
		Role:      entity.Role,
	}
}
