package repository

import (
	"JwtAuth/entity"
	"strings"

	"gorm.io/gorm"
)

type ProvincesRepo interface {
	SelectAll() ([]entity.Provinces, error)
	SelectById(id int8) (entity.Provinces, error)
	SelectByContain(condition string) ([]entity.Provinces, error)
}

type provincesConn struct{
	connection *gorm.DB
}

func NewProvincesRepo(db *gorm.DB) ProvincesRepo{
	return &provincesConn{
		connection: db,
	}
}

func (db *provincesConn) SelectAll() ([]entity.Provinces, error){
	var entProvs []entity.Provinces
	err := db.connection.Find(&entProvs).Error
	return entProvs, err
}

func (db *provincesConn) SelectById(id int8) (entity.Provinces, error){
	var entProv entity.Provinces
	err := db.connection.Where("prov_id = ?", id).Find(&entProv).Error
	return entProv, err
}

func(db *provincesConn) SelectByContain(condition string) ([]entity.Provinces, error){
	var entProvs []entity.Provinces
	newCond := "%"+strings.ToUpper(condition)+"%"
	err := db.connection.Where("prov_name LIKE ?", newCond).Find(&entProvs).Error
	return entProvs, err
}