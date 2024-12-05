package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/umeh-promise/simple_bank/db/sqlc"
)

// Server serves all HTTP request
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server iinstance and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts/:id", server.updateAccount)

	server.router = router
	return server
}

// Starts the HTTP server on the specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) *gin.H {
	return &gin.H{"error": err.Error()}
}
