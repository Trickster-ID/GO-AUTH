package repository

import (
	"JwtAuth/entity"
	"strings"

	"gorm.io/gorm"
)

type DistrictsRepo interface {
	SelectAll() ([]entity.Districts, error)
	SelectById(id int16) (entity.Districts, error)
	SelectByParent(parent int16) ([]entity.Districts, error)
	SelectByContain(condition string) ([]entity.Districts, error)
}

type districtsConn struct{
	connection *gorm.DB
}

func NewDistrictsRepo(db *gorm.DB) DistrictsRepo{
	return &districtsConn{
		connection: db,
	}
}

var entDistricts []entity.Districts

func(db *districtsConn) SelectAll() ([]entity.Districts, error){
	err := db.connection.Find(&entDistricts).Error
	return entDistricts, err
}

func(db *districtsConn) SelectById(id int16) (entity.Districts, error){
	var entDistrict entity.Districts
	err := db.connection.Find(&entDistrict, id).Error
	return entDistrict, err
}

func(db *districtsConn) SelectByParent(parent int16) ([]entity.Districts, error){
	err := db.connection.Where("city_id = ?", parent).Find(&entDistricts).Error
	return entDistricts, err
}

func(db *districtsConn) SelectByContain(condition string) ([]entity.Districts, error){
	newCond := "%"+strings.ToUpper(condition)+"%"
	err := db.connection.Where("dis_name LIKE ?", newCond).Find(&entDistricts).Error
	return entDistricts, err
}