package app

import (
	"context"
	"github.com/alex-dev-master/fibonacci.git/intrernal/handler"
	"github.com/alex-dev-master/fibonacci.git/intrernal/server"
	"github.com/alex-dev-master/fibonacci.git/intrernal/service"
	"github.com/alex-dev-master/fibonacci.git/pkg/cache"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	ctx := context.Background()

	rdbCache := cache.NewRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASS"), 0, ctx).RunRedis()
	services := service.NewService(rdbCache)
	handlers := handler.NewHandler(services)

	srv := new(server.Bootstrap)
	srvGrpc := new(server.BootstrapGrpc)

	go func() {
		if err := srv.Run(viper.GetString("http.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	go func() {
		if err := srvGrpc.RunRpc(services); err != nil {
			logrus.Fatalf("error occured while running grpc server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	srvGrpc.Shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
