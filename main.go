package main

import (
	"campsite_reservation/config"
	user_cl "campsite_reservation/controller/user"
	user_rp "campsite_reservation/drivers/mysql/user"
	user_uc "campsite_reservation/usecase/user"

	"campsite_reservation/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.InitConfigMySQL()
	// DB := mysql.ConnectDB(config.InitConfigMySQL())

	// gcs_credentials := os.Getenv("GCS_CREDENTIALS")
	// gmaps_api_key := os.Getenv("GMAPS_API_KEY")
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDBMysql(cfg)
	config.InitMigrationMysql(db)

	userRepo := user_rp.NewUserRepo(db)
	userUsecase := user_uc.NewUserUseCase(userRepo)
	userController := user_cl.NewUserController(userUsecase)

	routes := routes.RouteController{

		UserController: userController,
		// AdminController:                 AdminController,

	}

	routes.InitRoute(e)

	e.Start(":3000")
}
