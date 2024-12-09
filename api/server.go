package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/umeh-promise/simple_bank/db/sqlc"
	"github.com/umeh-promise/simple_bank/token"
	"github.com/umeh-promise/simple_bank/util"
)

// Server serves all HTTP request
type Server struct {
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server iinstance and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMarker(config.TokenSecret)
	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/login", server.loginUser)

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts/:id", server.updateAccount)

	server.router = router
}

// Starts the HTTP server on the specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) *gin.H {
	return &gin.H{"error": err.Error()}
}
