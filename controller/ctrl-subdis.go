package controller

import (
	"JwtAuth/helper"
	"JwtAuth/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubdistrictsController interface {
	SubdisGetAll(ctx *gin.Context)
	SubdisGetById(ctx *gin.Context)
	SubdisGetByParent(ctx *gin.Context)
	SubdisGetByContain(ctx *gin.Context)
}

type subdistrictsController struct{
	subdisRepo repository.SubdistrictsRepo
}

func NewSubdistrictsController(SubdisRepo repository.SubdistrictsRepo) SubdistrictsController{
	return &subdistrictsController{
		subdisRepo: SubdisRepo,
	}
}

func(rp *subdistrictsController) SubdisGetAll(ctx *gin.Context){
	res, err := rp.subdisRepo.SelectAll()
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Subdistricts",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *subdistrictsController) SubdisGetById(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("id"),10,32)
	if errConv != nil{
		responseErr := helper.BuildErrorResponse("Fail when convert Id",errConv.Error(),helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	res, err := rp.subdisRepo.SelectById(int(id))
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Subdistricts",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *subdistrictsController) SubdisGetByParent(ctx *gin.Context){
	id, errConv := strconv.ParseInt(ctx.Param("parent"),10,16)
	if errConv != nil{
		responseErr := helper.BuildErrorResponse("Fail when convert Id parent",errConv.Error(),helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	res, err := rp.subdisRepo.SelectByParent(int16(id))
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select all Subdistricts",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}

func(rp *subdistrictsController) SubdisGetByContain(ctx *gin.Context){
	contain := ctx.Param("contain")
	res, err := rp.subdisRepo.SelectByContain(contain)
	if err != nil{
		responseErr := helper.BuildErrorResponse("Fail when select by contain",err.Error(),res)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	responseSuccess := helper.BuildResponse(res)
	ctx.JSON(http.StatusOK, responseSuccess)
}