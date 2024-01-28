package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"interview/pkg/core/common"
	"interview/pkg/interface/handler"
	"net/http"
)

// API is the interface for the API server.
type API interface {
	// Start starts the API server
	Start() error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	SetSessionKye()
}

// apiServer is the implementation of the API interface for gin.
type apiServer struct {
	port     int
	gin      *gin.Engine
	Services common.Services
	handler  *handler.Handler
	mode     string
}

// Start starts the API server.
func (s *apiServer) Start() error {
	return s.gin.Run(fmt.Sprintf(":%d", s.port))
}

type ServerConfig struct {
	Port     int
	GinMode  string
	Services common.Services
}

// NewServer creates a new API server.
func NewServer(c *ServerConfig) API {

	ginHandler := gin.Default()

	server := &apiServer{
		port:     c.Port,
		gin:      ginHandler,
		Services: c.Services,
		mode:     c.GinMode,
		handler:  handler.NewHandler(c.Services),
	}

	server.AddCors()
	server.SetSessionKye()
	server.RegisterHealthCheck()
	server.RegisterRoutes()

	return server
}

func (s *apiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.gin.ServeHTTP(w, r)
}

// RegisterHealthCheck registers a health check endpoint.
func (s *apiServer) RegisterHealthCheck() {
	s.gin.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

// RegisterHealthCheck registers a health check endpoint.
func (s *apiServer) RegisterRoutes() {
	apiGroup := s.gin.Group("")
	if s.mode == "test" {
		apiGroup.Use(TestSession)
	}

	apiGroup.GET("/", s.handler.TaxController.ShowAddItemForm)
	apiGroup.POST("/add-item", s.handler.TaxController.AddItem)
	apiGroup.GET("/remove-cart-item", s.handler.TaxController.DeleteCartItem)
}

func (s *apiServer) AddCors() {
	s.gin.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	})
}

func (s *apiServer) SetSessionKye() {
	store := cookie.NewStore([]byte("d74d44f8-e8e6-44d0-83e5-7a5ceaf1303e"))
	store.Options(sessions.Options{
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: 0,
	})
	s.gin.Use(sessions.Sessions(common.IceSessionIdKey, store))

}

func TestSession(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get(common.IceSessionIdKey)
	if sessionID == nil {
		session.Set(common.IceSessionIdKey, uuid.NewString())
		err := session.Save()
		if err != nil {
			return
		}
	}
	c.Next()
}
