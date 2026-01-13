package models

import (
	"time"
)

// User model untuk tabel users
// @Description Model data user
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Name      string    `json:"name" gorm:"not null" example:"John Doe"`
	Email     string    `json:"email" gorm:"unique;not null" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-13T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-13T12:00:00Z"`
}

// CreateUserRequest untuk request body saat create user
// @Description Request body untuk membuat user baru
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
}

// UpdateUserRequest untuk request body saat update user
// @Description Request body untuk update user
type UpdateUserRequest struct {
	Name  string `json:"name" example:"John Updated"`
	Email string `json:"email" example:"john.updated@example.com"`
}
