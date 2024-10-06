package api

import (
	"fmt"

	db "bitemap/db/sqlc"
	"bitemap/token"
	"bitemap/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// serve HTTP requests for banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Change this to your allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}
	router.Use(cors.New(config))
	router.GET("/", server.healthCheck)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.Use(cors.New(config))
	authRoutes.GET("/restaurants", server.listRestaurants)
	authRoutes.GET("/restaurants/cuisines", server.getRestaurantCuisines)
	authRoutes.POST("restaurants/review", server.addReview)
	authRoutes.GET("/restaurants/filter", server.getRestaurantsByFilter)
	authRoutes.GET("/restaurants/review/:id", server.getRestaurantReviews)

	fmt.Println(authRoutes)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
