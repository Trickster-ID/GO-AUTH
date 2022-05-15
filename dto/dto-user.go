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
	Name      string `json:"name,omitempty" form:"name"`
	Email     string `json:"email,omitempty" form:"email" binding:"required" validate:"email"`
	Password  string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
	Prov_id   int	`json:"prov_id,omitempty" form:"prov_id"`
	City_id   int	`json:"city_id,omitempty" form:"city_id"`
	Dis_id    int	`json:"dis_id,omitempty" form:"dis_id"`
	Subdis_id int	`json:"subdis_id,omitempty" form:"subdis_id"`
	Avatar    *multipart.FileHeader `json:"avatar,omitempty" form:"avatar"`
}