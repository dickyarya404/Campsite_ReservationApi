package routes

import (
	"campsite_reservation/controller/admin"
	"campsite_reservation/controller/user"
	"campsite_reservation/middleware"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	UserController  *user.UserController
	AdminController *admin.AdminController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwtConfig = echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:   "cookie:JwtToken",
	}

	// User
	user := e.Group("/api/user")
	user.POST("/register", r.UserController.Register)
	user.POST("/login", r.UserController.Login)
	user.Use(echojwt.WithConfig(jwtConfig), middleware.User)
	user.PUT("/change-password", r.UserController.ChangePassword)

	// Admin
	admin := e.Group("/api/admin")
	admin.POST("/login", r.AdminController.Login)
	admin.Use(echojwt.WithConfig(jwtConfig), middleware.Admin)
	admin.GET("", r.AdminController.GetAll)
	admin.GET("/:id", r.AdminController.GetByID)
	admin.PUT("/reset-password", r.AdminController.ResetPassword)
	admin.PUT("/change-password", r.AdminController.ChangePassword)
	admin.GET("", r.AdminController.GetAll)
	admin.GET("/:id", r.AdminController.GetByID)
	admin.DELETE("user/:id/delete", r.AdminController.DeleteAccount)
	admin.PUT("user/:id/change-password", r.AdminController.ChangePassword)
}
