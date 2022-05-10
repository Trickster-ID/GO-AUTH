package controller

import (
	"JwtAuth/helper"
	"JwtAuth/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProvController interface {
	ProvGetAll(ctx *gin.Context)
	ProvGetById(ctx *gin.Context)
	ProvGetByContain(ctx *gin.Context)
}

type provController struct {
	provRepo repository.ProvincesRepo
}

func NewProvController(ProvRepo repository.ProvincesRepo) ProvController{
	return &provController{provRepo: ProvRepo}
}

func (rp *provController) ProvGetAll(ctx *gin.Context){
	res, err := rp.provRepo.SelectAll()
	if err != nil{
		responseError := helper.BuildErrorResponse("Fail to get all Provinces",err.Error(), res)
		ctx.JSON(http.StatusBadRequest, responseError)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func (rp *provController) ProvGetById(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("id"),10,8)
	if errConv != nil{
		responseError := helper.BuildErrorResponse("Fail to convert parameter",errConv.Error(), helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseError)
	}
	res, err := rp.provRepo.SelectById(int8(id))
	if err != nil{
		responseError := helper.BuildErrorResponse("Fail to get all Provinces",err.Error(), res)
		ctx.JSON(http.StatusBadRequest, responseError)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *provController) ProvGetByContain(ctx *gin.Context){
	contain := ctx.Param("contain")
	res, err := rp.provRepo.SelectByContain(contain)
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select by contain",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}