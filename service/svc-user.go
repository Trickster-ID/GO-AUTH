package service

import (
	"JwtAuth/dto"
	"JwtAuth/entity"
	"JwtAuth/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) (interface{}, error)
	CreateUser(userdto dto.RegisterPostDTO, filename string) (entity.Getcompleteuser, error)
	UpdateProfile(userdto dto.UserPutDto, filename string) (entity.Getcompleteuser, error)
	FindByEmail(email string) (entity.Getcompleteuser, error)
	FindProfile(UserID string) (entity.Getcompleteuser, error)
	IsDuplicateEmail(email string) bool
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
	entUserC := entity.Getcompleteuser{}
	u.Id = d.Id
	u.Name = d.Name
	u.Email = d.Email
	reshas, errhas := hashAndSalt([]byte(u.Password))
	if errhas != nil{
		return entUserC, errhas
	}
	u.Password = string(reshas)
	u.Prov_id = d.Prov_id
	u.City_id = d.City_id
	u.Dis_id = d.Dis_id
	u.Subdis_id = d.Subdis_id
	u.Avatar = filename
	u.Update_at = time.Now()
	return svc.userRepo.UpdateUser(u)
}

func (svc *authService) FindByEmail(email string) (entity.Getcompleteuser, error){
	return svc.userRepo.FindByEmail(email)
}

func (svc *authService) FindProfile(UserID string) (entity.Getcompleteuser, error) {
	return svc.userRepo.ProfileUser(UserID)
}

func (svc *authService) IsDuplicateEmail(email string) bool{
	res := svc.userRepo.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPass string) bool{
	// res := true
	// byteHash := []byte(hashedPwd)
	// err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	// if err != nil{
	// 	fmt.Println("fail when comparing hash and password!!")
	// 	res = false
	// }
	// return res, err
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPass))
	return err == nil
}

func hashAndSalt(pwd []byte) ([]byte, error){
	return bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
}