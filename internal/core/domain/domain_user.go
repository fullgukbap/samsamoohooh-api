package domain

import "gorm.io/gorm"

type RoleType string

const (
	Admin  RoleType = "ADMIN"
	Member RoleType = "MEMBER"
	Guest  RoleType = "GUEST"
)

type User struct {
	gorm.Model
	Name       string
	Resolution string
	Role       RoleType
	Sub        string
}
