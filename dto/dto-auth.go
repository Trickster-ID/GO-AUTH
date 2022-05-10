package dto

import "mime/multipart"

type LoginPostDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
}

type PasswordDto struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
}

type RegisterPostDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required"`
	Avatar   *multipart.FileHeader `json:"avatar" form:"avatar"`
	Prov_id  int `json:"prov_id" form:"prov_id"`
	City_id  int `json:"city_id" form:"city_id"`
	Dis_id   int `json:"dis_id" form:"dis_id"`
	Subdis_id int `json:"subdis_id" form:"subdis_id"`
} 