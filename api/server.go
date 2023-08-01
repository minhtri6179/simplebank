package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/minhtri67/simplebank/db/sqlc"
	"github.com/minhtri67/simplebank/token"
	"github.com/minhtri67/simplebank/util"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	store      sqlc.Store
	router     *gin.Engine
}

func NewServer(config util.Config, store sqlc.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetric)
	if err != nil {
		return nil, fmt.Errorf("cannot create token")
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)
	server.router = router

}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
