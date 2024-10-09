package router

import (
	"ecommerce-go/internal/infrastructure/ginhandler/handler"
	"github.com/gin-gonic/gin"
)

type HTTPServer interface {
	RegisterRouter()
	Run(port string) error
}

type httpServer struct {
	router       *gin.Engine
	trController *handler.TransactionHandler
}

func NewHTTPServer(router *gin.Engine,
	transactionCtrl *handler.TransactionHandler,
) HTTPServer {
	return &httpServer{
		router:       router,
		trController: transactionCtrl,
	}
}

func (srv *httpServer) RegisterRouter() {
	basePath := "/api"
	router := srv.router.Group(basePath)

	router.GET("/products", srv.trController.GetProducts)

	router.POST("/products", srv.trController.PostProduct)

}

func (srv *httpServer) Run(port string) error {
	return srv.router.Run(":" + port)
}
