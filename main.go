package main

import (
	_ "ecommerce-go/docs"
	"ecommerce-go/internal/core/service"
	"ecommerce-go/internal/infrastructure/environmentconfig"
	"ecommerce-go/internal/infrastructure/ginhandler/auth"
	"ecommerce-go/internal/infrastructure/ginhandler/handler"
	"ecommerce-go/internal/infrastructure/ginhandler/router"
	"ecommerce-go/internal/infrastructure/jobs"
	"ecommerce-go/internal/infrastructure/repository"
	gorm_repo "ecommerce-go/internal/infrastructure/repository/gorm-repo"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Ecommerce API
// @version         0.1
// @description     API for ecommerce

// @host      localhost:8080
// @BasePath  /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	configService := environmentconfig.NewConfigService()
	envVariables := configService.GetDomainEnv()

	db, err := gorm_repo.GetGormConnection(envVariables)
	if err != nil {
		return
	}

	repo := repository.NewRepository(db)

	serv := service.NewEcommerceService(repo)

	ginRouter := gin.Default()

	ginRouter.Use(auth.TokenMiddleware(envVariables))

	transactionController := handler.NewTransactionHandler(serv)

	go jobs.ScheduleAllJobs(serv, envVariables)

	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginServer := router.NewHTTPServer(ginRouter, transactionController)

	ginServer.RegisterRouter()

	if err := ginServer.Run("8080"); err != nil {
		panic(err.Error())
	}

	/*	var products []domain.Product
		db.Find(&products)
		db.First(&product, 1)
		fmt.Printf("%+v\n", product)
		fmt.Printf("%+v\n", products)
	*/

	//db.Create(&domain.Product{})
}
