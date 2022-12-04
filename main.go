package main

import (
	"JwtAuth/config"
	"JwtAuth/controller"
	"JwtAuth/middleware"
	"JwtAuth/repository"
	"JwtAuth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	//DB
	db *gorm.DB = config.SetUpDatabaseConnection()
	//Repository
	rpUser   repository.UserRepo         = repository.NewUserRepo(db)
	rpProv   repository.ProvincesRepo    = repository.NewProvincesRepo(db)
	rpCity   repository.CitiesRepo       = repository.NewCitiesRepo(db)
	rpDis    repository.DistrictsRepo    = repository.NewDistrictsRepo(db)
	rpSubDis repository.SubdistrictsRepo = repository.NewSubdistrictsRepo(db)
	//Service
	svcUser service.AuthService = service.NewAuthService(rpUser)
	svcJwt  service.JWTService  = service.NewJWTService()
	//Controller
	ctrlUser   controller.AuthController         = controller.NewAuthController(svcUser, svcJwt)
	ctrlProv   controller.ProvController         = controller.NewProvController(rpProv)
	ctrlCity   controller.CitiesController       = controller.NewCitiesController(rpCity)
	ctrlDis    controller.DistrictsController    = controller.NewDistrictsController(rpDis)
	ctrlSubDis controller.SubdistrictsController = controller.NewSubdistrictsController(rpSubDis)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/login", ctrlUser.Login)
		apiGroup.POST("/register", ctrlUser.Register)

		apiGroup.GET("/provinsi", ctrlProv.ProvGetAll)
		apiGroup.GET("/provinsi/:id", ctrlProv.ProvGetById)
		apiGroup.GET("/provinsi/c/:contain", ctrlProv.ProvGetByContain)

		apiGroup.GET("/kota", ctrlCity.CitGetAll)
		apiGroup.GET("/kota/:id", ctrlCity.CitGetById)
		apiGroup.GET("/kota/p/:parent", ctrlCity.CitGetByParent)
		apiGroup.GET("/kota/c/:contain", ctrlCity.CitGetByContain)

		apiGroup.GET("/kecamatan", ctrlDis.DisGetAll)
		apiGroup.GET("/kecamatan/:id", ctrlDis.DisGetById)
		apiGroup.GET("/kecamatan/p/:parent", ctrlDis.DisGetByParent)
		apiGroup.GET("/kecamatan/c/:contain", ctrlDis.DisGetByContain)

		apiGroup.GET("/desa", ctrlSubDis.SubdisGetAll)
		apiGroup.GET("/desa/:id", ctrlSubDis.SubdisGetById)
		apiGroup.GET("/desa/p/:parent", ctrlSubDis.SubdisGetByParent)
		apiGroup.GET("/desa/c/:contain", ctrlSubDis.SubdisGetByContain)
		authreq := apiGroup.Group("", middleware.AuthJWT(svcJwt))
		{
			authreq.GET("/profile", ctrlUser.GetProfile)
			authreq.PUT("/profile", ctrlUser.PutProfile)
		}
	}

	r.Run(":8888")
}
