package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/superjantung/bankita-api/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/accounts", server.createAccount)
		api.GET("/accounts/:id", server.getAccount)
		api.GET("/accounts", server.listAccount)
	}

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
