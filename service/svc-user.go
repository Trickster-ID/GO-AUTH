package service

import (
	"JwtAuth/dto"
	"JwtAuth/entity"
	"JwtAuth/helper"
	"JwtAuth/repository"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) (interface{}, error)
	CreateUser(userdto dto.RegisterPostDTO, filename string) (entity.Getcompleteuser, error)
	UpdateProfile(userdto dto.UserPutDto, filename string) (entity.Getcompleteuser, error)
	FindByEmail(email string) (entity.Getcompleteuser, error)
	FindProfile(UserID string) (entity.Getcompleteuser, error)
	IsDuplicateEmail(email string, id int) bool
}

type authService struct{userRepo repository.UserRepo}

func NewAuthService(userRep repository.UserRepo) AuthService{
	return &authService{
		userRepo: userRep,
	}
}

func (svc *authService) VerifyCredential(email string, password string) (interface{}, error) {
	resfind, errfind := svc.userRepo.VerifyCredential(email)
	if errfind != nil{
		return false, errfind
	}
	if v, ok := resfind.(entity.Getcompleteuser); ok {
		comparedPass := comparePassword(v.Password, password)
		if v.Email == email && comparedPass{
			return resfind, nil
		}
		return false, errors.New("fail to compare password")
	}
	return false, errfind
}

func (svc *authService) CreateUser(userdto dto.RegisterPostDTO, filename string) (entity.Getcompleteuser, error){
	var user = entity.User{}
	entUserC := entity.Getcompleteuser{}
	user.Name = userdto.Name
	user.Email = userdto.Email
	reshas, errhas := hashAndSalt([]byte(userdto.Password))
	if errhas != nil{
		return entUserC, errhas
	}
	user.Password = string(reshas)
	user.Avatar = filename
	user.Create_at = time.Now()
	user.Update_at = time.Now()
	return svc.userRepo.InsertUser(user)
}

func (svc *authService) UpdateProfile(d dto.UserPutDto, filename string) (entity.Getcompleteuser, error) {
	u := entity.User{}
	result := entity.Getcompleteuser{}
	//get current data
	cd, errcurdata := svc.userRepo.ProfileUser(strconv.Itoa(d.Id))
	if errcurdata != nil{
		return result, errcurdata
	}
	//fill to entity
	u.Id = cd.Id
	u.Name = helper.Ifelse(d.Name, cd.Name).(string)
	u.Email = helper.Ifelse(d.Email, cd.Email).(string)
	u.Prov_id = helper.Ifelse(d.Prov_id, cd.Prov_id).(int)
	u.City_id = helper.Ifelse(d.City_id, cd.City_id).(int)
	u.Dis_id = helper.Ifelse(d.Dis_id, cd.Dis_id).(int)
	u.Subdis_id = helper.Ifelse(d.Subdis_id, cd.Subdis_id).(int)
	u.Avatar = helper.Ifelse(filename, cd.Avatar).(string)
	u.Update_at = time.Now()
	u.Create_at = cd.Create_at

	//hash password
	if d.Password != ""{
		reshas, errhas := hashAndSalt([]byte(d.Password))
		if errhas != nil{
			return result, errhas
		}
		u.Password = string(reshas)
	}else{
		u.Password = cd.Password
	}
	return svc.userRepo.UpdateUser(u)
}

func (svc *authService) FindByEmail(email string) (entity.Getcompleteuser, error){
	return svc.userRepo.FindByEmail(email)
}

func (svc *authService) FindProfile(UserID string) (entity.Getcompleteuser, error) {
	return svc.userRepo.ProfileUser(UserID)
}

func (svc *authService) IsDuplicateEmail(email string, id int) bool{
	res := svc.userRepo.IsDuplicateEmail(email, id)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPass string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPass))
	return err == nil
}

func hashAndSalt(pwd []byte) ([]byte, error){
	return bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
}