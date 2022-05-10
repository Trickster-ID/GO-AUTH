package controller

import (
	"JwtAuth/dto"
	"JwtAuth/entity"
	"JwtAuth/helper"
	"JwtAuth/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(cx *gin.Context)
	Register(cx *gin.Context)
	GetProfile(cx *gin.Context)
	PutProfile(cx *gin.Context)
}

type authController struct{
	auth service.AuthService
	jwt service.JWTService
}

func NewAuthController(newauth service.AuthService, newjwt service.JWTService) AuthController{
	return &authController{
		auth: newauth,
		jwt: newjwt,
	}
}

func (c *authController) Login(cx *gin.Context){
	var loginDTO dto.LoginPostDTO
	SBerr := cx.ShouldBind(&loginDTO)
	if SBerr != nil{
		response := helper.BuildErrorResponse("failed to process request", SBerr.Error(), helper.EmptyObject{})
		cx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} 
	authresult, err := c.auth.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if err != nil{
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail when verify credentials", err.Error(), helper.EmptyObject{}))
		return
	}
	if v, ok := authresult.(entity.Getcompleteuser); ok{
		v.Token = c.jwt.GenerateToken(strconv.Itoa(v.Id))
		cx.JSON(http.StatusOK, helper.BuildResponse(v))
		return
	}
	respons := helper.BuildErrorResponse("Failed!", "check your credential", helper.EmptyObject{})
	cx.AbortWithStatusJSON(http.StatusUnauthorized, respons)
}

func (c *authController) Register(cx *gin.Context){
	var registerDTO dto.RegisterPostDTO
	SBerr := cx.ShouldBind(&registerDTO)
	if SBerr != nil{
		response := helper.BuildErrorResponse("failed to process request", SBerr.Error(), helper.EmptyObject{})
		cx.JSON(http.StatusConflict, response)
		return
	}
	if !(c.auth.IsDuplicateEmail(registerDTO.Email)){
		respons := helper.BuildErrorResponse("failed to register", "email is already exist", helper.EmptyObject{})
		cx.AbortWithStatusJSON(http.StatusUnauthorized, respons)
		return
	}
	filename, errSave:= helper.SaveAvatar(registerDTO.Avatar, cx)
	if errSave != nil{
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to save file", errSave.Error(), registerDTO.Avatar))
		return
	}
	createdUser, err := c.auth.CreateUser(registerDTO, filename)
	if err != nil{
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail when create user", err.Error(), createdUser))
		return
	}
	token := c.jwt.GenerateToken(strconv.FormatUint(uint64(createdUser.Id), 10))
	createdUser.Token = token
	response := helper.BuildResponse(createdUser)
	cx.JSON(http.StatusOK, response)
}

func (c *authController) GetProfile(cx *gin.Context){
	authHeader := cx.GetHeader("Authorization")
	token, errToken := c.jwt.ValidateToken(authHeader)
	if errToken != nil{
		//panic(errToken.Error())
		cx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse("you are not authorized", errToken.Error(), helper.EmptyObject{}))
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userid := fmt.Sprintf("%v", claims["user_id"])
	res, err := c.auth.FindProfile(userid)
	if err != nil{
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail when find profile", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse(res))
}

func (c *authController) PutProfile(cx *gin.Context){
	var userUpdateDTO dto.UserPutDto
	errDTO := cx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("Fail to update when binding body", errDTO.Error(), helper.EmptyObject{}))
		return
	}
	authHeader := cx.GetHeader("Authorization")
	token, errToken := c.jwt.ValidateToken(authHeader)
	if errToken != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to update when validate token", errToken.Error(), helper.EmptyObject{}))
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id, errPrs := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if errPrs != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to update when parse id to int", errPrs.Error(), id))
		return
	}
	userUpdateDTO.Id = int(id)
	filename, errSave := helper.SaveAvatar(userUpdateDTO.Avatar, cx)
	if errSave != nil{
		cx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to save file", errSave.Error(), userUpdateDTO.Avatar))
		return
	}
	if !(c.auth.IsDuplicateEmail(userUpdateDTO.Email)){
		cx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse("failed to update", "email is already exist", helper.EmptyObject{}))
		return
	}
	if userUpdateDTO.Id == 0{
		cx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse("failed to update", "id is null", helper.EmptyObject{}))
		return
	}
	u, errServ := c.auth.UpdateProfile(userUpdateDTO, filename)
	if errServ != nil{
		panic(errServ.Error())
	}
	cx.JSON(http.StatusOK, helper.BuildResponse(u))
}