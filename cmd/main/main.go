package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	server "github.com/wintersakura/golang-vue-todo"
	"github.com/wintersakura/golang-vue-todo/pkg/handler"
	"github.com/wintersakura/golang-vue-todo/pkg/repository"
	"github.com/wintersakura/golang-vue-todo/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	s := new(server.Server)

	go func() {
		if err := s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Print("Todo API Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Todo API Shutting Down")

	if err := s.ShutDown(context.Background()); err != nil {
		logrus.Error(err)
	}
	if err := db.Close(); err != nil {
		logrus.Error(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
