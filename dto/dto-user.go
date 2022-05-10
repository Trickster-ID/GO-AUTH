package dto

import "mime/multipart"

type UserPostDto struct {
	Name      string
	Email     string `json:"email" form:"email" binding:"required" validate:"email"`
	Password  string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
	Prov_id   int
	City_id   int
	Dis_id    int
	Subdis_id int
	Avatar    *multipart.FileHeader
}

type UserPutDto struct {
	Id        int `json:"omitempty"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email" binding:"required" validate:"email"`
	Password  string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
	Prov_id   int	`json:"prov_id" form:"prov_id"`
	City_id   int	`json:"city_id" form:"city_id"`
	Dis_id    int	`json:"dis_id" form:"dis_id"`
	Subdis_id int	`json:"subdis_id" form:"subdis_id"`
	Avatar    *multipart.FileHeader `json:"avatar" form:"avatar"`
}