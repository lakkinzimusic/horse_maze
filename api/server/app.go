package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/lakkinzimusic/horse_maze/api/auth"
	authhttp "github.com/lakkinzimusic/horse_maze/api/auth/handler"
	authsql "github.com/lakkinzimusic/horse_maze/api/auth/repository"
	authusecase "github.com/lakkinzimusic/horse_maze/api/auth/usecase"
	"github.com/lakkinzimusic/horse_maze/api/db"
)

//App struct
type App struct {
	httpServer  *http.Server
	authUseCase auth.UseCase
}

//NewApp func
func NewApp() *App {
	db := db.InitDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	userRepo := authsql.NewUserRepository(db)

	return &App{
		authUseCase: authusecase.NewAuthUseCase(userRepo),
	}
}

//Run func
func (a *App) Run(port string) error {
	// Init gin handler
	router := mux.NewRouter()
	authhttp.RegisterHTTPEndpoints(router, a.authUseCase)
	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
