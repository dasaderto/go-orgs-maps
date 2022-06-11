package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "go-gr-maps/docs"
	"go-gr-maps/pkg"
	"go-gr-maps/pkg/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func dbConnect() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DBNAME"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	return db, err
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := dbConnect()
	if err != nil {
		logrus.Fatalf("Failed db connect: %s", err.Error())
		return
	}

	seeder := pkg.NewSeeder(db)
	seeder.SeedAll()

	var router = api.NewRouter(db)
	router.InitRouter()

	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("PORT"), router.Router); err != nil {
			logrus.Fatalf("error occured while running api server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	err = db.Close()
	if err != nil {
		logrus.Fatalf("Failed connection closing: %s", err.Error())
		return
	}
	logrus.Print("DB connection Closed")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
