package app

import (
	"context"
	"github.com/alex-dev-master/fibonacci.git/intrernal/handler"
	"github.com/alex-dev-master/fibonacci.git/intrernal/server"
	"github.com/alex-dev-master/fibonacci.git/intrernal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Run()  {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(server.Bootstrap)

	go func () {
		if err := srv.Run(viper.GetString("http.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("Chat Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}


}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}