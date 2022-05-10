package controller

import (
	"JwtAuth/helper"
	"JwtAuth/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DistrictsController interface {
	DisGetAll(ctx *gin.Context)
	DisGetById(ctx *gin.Context)
	DisGetByParent(ctx *gin.Context)
	DisGetByContain(ctx *gin.Context)
}

type districtsController struct{
	disRepo repository.DistrictsRepo
}

func NewDistrictsController(DisRepo repository.DistrictsRepo) DistrictsController{
	return &districtsController{
		disRepo: DisRepo,
	}
}

func(rp *districtsController) DisGetAll(ctx *gin.Context){
	res, err := rp.disRepo.SelectAll()
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Districts",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *districtsController) DisGetById(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("id"),10,16)
	if errConv != nil{
		responseErr := helper.BuildErrorResponse("Fail when convert Id",errConv.Error(),helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	res, err := rp.disRepo.SelectById(int16(id))
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Districts",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *districtsController) DisGetByParent(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("parent"),10,16)
	if errConv != nil{
		responseErr := helper.BuildErrorResponse("Fail when convert Id parent",errConv.Error(),helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	res, err := rp.disRepo.SelectByParent(int16(id))
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Districts",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *districtsController) DisGetByContain(ctx *gin.Context){
	contain := ctx.Param("contain")
	res, err := rp.disRepo.SelectByContain(contain)
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select by contain",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}