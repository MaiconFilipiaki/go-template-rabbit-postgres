package http

import (
	"github.com/gin-gonic/gin"
	v1 "golangNetHttp/http/api/v1"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "9000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	v1.HandlersV1Http(s.server)

	log.Fatal(s.server.Run(":" + s.port))
}
