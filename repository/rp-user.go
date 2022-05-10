package repository

import (
	"JwtAuth/entity"
	"strconv"

	"gorm.io/gorm"
)

type UserRepo interface {
	ProfileUser(id string) (entity.Getcompleteuser, error)
	FindByEmail(email string) (entity.Getcompleteuser, error)
	VerifyCredential(email string) (interface{}, error)
	IsDuplicateEmail(email string) (tx *gorm.DB)
	InsertUser(user entity.User) (entity.Getcompleteuser, error)
	UpdateUser(user entity.User) (entity.Getcompleteuser, error)
}

type userConn struct{
	connection *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userConn{
		connection: db,
	}
}

func (db *userConn) ProfileUser(id string) (entity.Getcompleteuser, error){
	var gcuser entity.Getcompleteuser
	err := db.connection.Find(&gcuser, id).Error
	return gcuser, err
}

func (db *userConn) FindByEmail(email string) (entity.Getcompleteuser, error){
	var gcuser entity.Getcompleteuser
	err := db.connection.Where("email = ?",email).Take(&gcuser).Error
	return gcuser, err
}

func (db *userConn) VerifyCredential(email string) (interface{}, error){
	var user entity.Getcompleteuser
	err := db.connection.Where("email = ?",email).Take(&user).Error
	if err != nil{
		return nil, err
	}
	return user, err
}	

func (db *userConn) IsDuplicateEmail(email string) (tx *gorm.DB){
	var user entity.User
	return db.connection.Where("email = ?",email).Take(&user)
}	

func (db *userConn) InsertUser(user entity.User) (entity.Getcompleteuser, error){
	var nilentity entity.Getcompleteuser
	err := db.connection.Create(&user).Error
	if err != nil{
		return nilentity, err
	}
	resget, errget := db.ProfileUser(strconv.Itoa(user.Id))
	return resget, errget
}

func (db *userConn) UpdateUser(user entity.User) (entity.Getcompleteuser, error){
	var nilentity entity.Getcompleteuser
	err := db.connection.Save(&user).Error
	if err != nil{
		return nilentity, err
	}
	resget, errget := db.ProfileUser(strconv.Itoa(user.Id))
	return resget, errget
}