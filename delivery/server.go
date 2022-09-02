package delivery

import (
	"github.com/wisnuragaprawida/golang-rest-api/config"
	"github.com/wisnuragaprawida/golang-rest-api/delivery/controller"
	"github.com/wisnuragaprawida/golang-rest-api/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initControllers() {
	controller.NewProductController(a.engine, a.useCaseManager.CreateProductUseCase(), a.useCaseManager.ListProductUseCase())

	controller.NewUserController(a.engine, a.useCaseManager.LoginUserUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)

	repoManager := manager.NewRepositoryManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := appConfig.Url
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}
