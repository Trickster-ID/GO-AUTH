package entity

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Token 	  string `gorm:"-" json:"token,omitempty"`
	Prov_id   int
	City_id   int
	Dis_id    int
	Subdis_id int
	Avatar	  string
	Create_at time.Time
	Update_at time.Time
}

type Getcompleteuser struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Token 	  string `gorm:"-" json:"token,omitempty"`
	Prov_id   int
	Prov_name string
	City_id   int
	City_name string
	Dis_id    int
	Dis_name  string
	Subdis_id int
	Subdis_name string
	Avatar	  string
	Create_at time.Time
	Update_at time.Time
}