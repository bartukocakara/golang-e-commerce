package route

import (
	"go-jwt/handler"
	"net/http"

	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to JWT System Deneme")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:user", userHandler.GetUser)
	}

	productProtectedRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productProtectedRoutes.GET("/", productHandler.GetAllProducts)
		productProtectedRoutes.GET("/:product", productHandler.GetProduct)
	}




	return r.Run(address)

}