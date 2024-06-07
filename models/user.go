package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:unique`
	Password string `json:"password"`
	// Role string `json:"role"`
	// Token string `json:"token"`
	// IsActive bool `json:"is_active"`
	// IsDeleted bool `json:"is_deleted"`
	// IsVerified bool `json:"is_verified"`
	// IsAdmin bool `json:"is_admin"`
	// IsSuperAdmin bool `json:"is_super_admin"`
	// IsStaff bool `json:"is_staff"`
	// IsCustomer bool `json:"is_customer"`
	// IsVendor bool `json:"is_vendor"`
}