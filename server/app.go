package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/dm/chat-x-back/auth"
	authhttp "github.com/dm/chat-x-back/auth/delivery/http"
	authpg "github.com/dm/chat-x-back/auth/repository/postgres"
	authusecase "github.com/dm/chat-x-back/auth/usecase"

	"github.com/dm/chat-x-back/database"
)

// App ...
type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
}

// NewApp ...
func NewApp() *App {

	userRepo := authpg.NewUserRepository(database.DB)

	return &App{
		authUC: authusecase.NewAuthUseCase(
			userRepo, viper.GetString("hashSalt"),
			[]byte(viper.GetString("signinKey")),
			viper.GetDuration("tokenTtl"),
		),
	}
}

// Start ...
func (a *App) Start() error {

	r := gin.Default()

	authhttp.RegisterHTTPEndpoints(r, a.authUC)

	a.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("url"), viper.GetString("port")),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)

}
