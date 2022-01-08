package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/raihanlh/go-kubernetes/src/config"
	"github.com/raihanlh/go-kubernetes/src/controller"
	"github.com/raihanlh/go-kubernetes/src/repository"
	"github.com/raihanlh/go-kubernetes/src/service"
	"github.com/spf13/viper"
)

func main() {
	var db *sql.DB

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", configuration.Db.Host, configuration.Db.User, configuration.Db.Password, configuration.Db.Name)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	bulletinRepo := repository.NewPostgresqlBulletinRepository(db)
	bulletinService := service.NewBulletinService(bulletinRepo)
	bulletinController := controller.NewBulletinController(bulletinService)

	app := fiber.New()
	_, err = db.Query(configuration.Db.Migration)
	if err != nil {
		log.Println("failed to run migration", err.Error())
	}

	controller.RouteAll(app, bulletinController)

	log.Println("server is running...")

	log.Fatal(app.Listen(":3000"))
}
