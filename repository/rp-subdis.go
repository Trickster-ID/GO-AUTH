package repository

import (
	"JwtAuth/entity"
	"strings"

	"gorm.io/gorm"
)

type SubdistrictsRepo interface {
	SelectAll() ([]entity.Subdistricts, error)
	SelectById(id int) (entity.Subdistricts, error)
	SelectByParent(parent int16) ([]entity.Subdistricts, error)
	SelectByContain(condition string) ([]entity.Subdistricts, error)
}

type subdistrictsConn struct{
	connection *gorm.DB
}

func NewSubdistrictsRepo(db *gorm.DB) SubdistrictsRepo{
	return &subdistrictsConn{
		connection: db,
	}
}

var entSubdistricts []entity.Subdistricts

func(db *subdistrictsConn) SelectAll() ([]entity.Subdistricts, error){
	err := db.connection.Find(&entSubdistricts).Error
	return entSubdistricts, err
}

func(db *subdistrictsConn) SelectById(id int) (entity.Subdistricts, error){
	var entSubDis entity.Subdistricts
	err := db.connection.Find(&entSubDis, id).Error
	return entSubDis, err
}

func(db *subdistrictsConn) SelectByParent(parent int16) ([]entity.Subdistricts, error){
	err := db.connection.Where("dis_id = ?", parent).Find(&entSubdistricts).Error
	return entSubdistricts, err
}

func(db *subdistrictsConn) SelectByContain(condition string) ([]entity.Subdistricts, error){
	newCond := "%"+strings.ToUpper(condition)+"%"
	err := db.connection.Where("subdis_name LIKE ?", newCond).Find(&entSubdistricts).Error
	return entSubdistricts, err
}