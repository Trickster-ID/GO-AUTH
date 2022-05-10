package repository

import (
	"JwtAuth/entity"
	"strings"

	"gorm.io/gorm"
)

type CitiesRepo interface {
	SelectAll() ([]entity.Cities, error)
	SelectById(id int16) (entity.Cities, error)
	SelectByParent(parent int8) ([]entity.Cities, error)
	SelectByContain(condition string) ([]entity.Cities, error)
}

type citiesConn struct{
	connection *gorm.DB
}

func NewCitiesRepo(db *gorm.DB) CitiesRepo{
	return &citiesConn{
		connection: db,
	}
}

var entCities []entity.Cities

func(db *citiesConn) SelectAll() ([]entity.Cities, error){
	err := db.connection.Find(&entCities).Error
	return entCities, err
}

func(db *citiesConn) SelectById(id int16) (entity.Cities, error){
	var entCity entity.Cities
	err := db.connection.Find(&entCity, id).Error
	return entCity, err
}

func(db *citiesConn) SelectByParent(parent int8) ([]entity.Cities, error){
	err := db.connection.Where("prov_id = ?", parent).Find(&entCities).Error
	return entCities, err
}

func(db *citiesConn) SelectByContain(condition string) ([]entity.Cities, error){
	newCond := "%"+strings.ToUpper(condition)+"%"
	err := db.connection.Where("city_name LIKE ?", newCond).Find(&entCities).Error
	return entCities, err
}