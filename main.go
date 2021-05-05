package main

import (
	"fmt"
	"log"
	"net/http"
	"svc-users-go/config"

	UserHandlerPackage "svc-users-go/module/v1/user/presenter"
	UserRepoPackage "svc-users-go/module/v1/user/repository"
	UserUseCasePackage "svc-users-go/module/v1/user/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "svc-users-go/docs/echosimple"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API Users Golang POC - VerifyMyAge Company
// @ versão 1.0
// @description Swagger documentation API for user.
// @termsOfService http://swagger.io/terms/

// @ contact.name Isaque Souza Ramos
// @ contact.url https://www.linkedin.com/in/isaquesouzaramos/
// @ contact.email isaquesouzaramos@gmail.com

// @ license.name Apache 2.0
// @ license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
func main() {
	fmt.Println("Server is running :)")

	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	echoRouter := echo.New()
	// Routes
	echoRouter.GET("/swagger/*", echoSwagger.WrapHandler)

	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// User modules
	userRepo := UserRepoPackage.NewUserRepository(mongoConnection)
	userUseCase := UserUseCasePackage.NewUserUseCase(userRepo)
	UserHandlerPackage.NewUserHandler(echoRouter, userUseCase)

	//Configuration of logger
	echoRouter.Use(middleware.Logger())
	echoRouter.Logger.Fatal(echoRouter.Start(":8000"))
}
