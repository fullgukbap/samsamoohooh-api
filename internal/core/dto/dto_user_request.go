package dto

import "samsamoohooh-mini-api/internal/core/domain"

type CreateUserRequest struct {
	Name       string            `json:"name" validate:"gte=6,lte=18"`
	Resolution string            `json:"resolution"`
	Role       domain.RoleType   `json:"role" validate:"oneof= ADMIN MEMBER GUEST"`
	Sub        string            `json:"sub"`
	Soical     domain.SocialType `json:"social" validate:"oneof=  GOOGLE KAKAO APPLE"`
}

func (dto CreateUserRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       dto.Name,
		Resolution: dto.Resolution,
		Role:       dto.Role,
		Sub:        dto.Sub,
		Social:     dto.Soical,
	}
}

type UpdateUserRequest struct {
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
}

func (dto UpdateUserRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       dto.Name,
		Resolution: dto.Resolution,
	}
}
