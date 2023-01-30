package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nitinkumar91296/simplebank/db/sqlc"
)

// Server serves HTTP requests for banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer reated a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
