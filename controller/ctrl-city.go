package controller

import (
	"JwtAuth/helper"
	"JwtAuth/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CitiesController interface {
	CitGetAll(ctx *gin.Context)
	CitGetById(ctx *gin.Context)
	CitGetByParent(ctx *gin.Context)
	CitGetByContain(ctx *gin.Context)
}

type citiesController struct{
	citRepo repository.CitiesRepo
}

func NewCitiesController(CitRepo repository.CitiesRepo) CitiesController{
	return &citiesController{
		citRepo: CitRepo,
	}
}

func(rp *citiesController) CitGetAll(ctx *gin.Context){
	res, err := rp.citRepo.SelectAll()
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Cities",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *citiesController) CitGetById(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("id"),10,16)
	if errConv != nil{
		responseErr := helper.BuildErrorResponse("Fail when convert Id",errConv.Error(),helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	res, err := rp.citRepo.SelectById(int16(id))
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Cities",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *citiesController) CitGetByParent(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("parent"),10,8)
	if errConv != nil{
		responseErr := helper.BuildErrorResponse("Fail when convert Id parent",errConv.Error(),helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	res, err := rp.citRepo.SelectByParent(int8(id))
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Cities",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *citiesController) CitGetByContain(ctx *gin.Context){
	contain := ctx.Param("contain")
	res, err := rp.citRepo.SelectByContain(contain)
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select by contain",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}