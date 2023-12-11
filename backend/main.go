package main

import (
	"context"
	"errors"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	// timeouts
	readTimeout     = 10 * time.Second
	writeTimeout    = 10 * time.Second
	shutdownTimeout = 15 * time.Second
)

type Server struct {
	Router     *mux.Router
	DB         *sqlx.DB
	HttpServer *http.Server
	Logger     *slog.Logger
}

func (s *Server) initDB() error {
	var err error
	s.DB, err = sqlx.Connect("mysql", "root:@tcp(localhost:3306)/imdb2")
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) initHttpServer(port string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	s.HttpServer = &http.Server{
		Addr:           port,
		Handler:        handlers.CORS(headersOk, originsOk, methodsOk)(s.Router),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
}

func NewServer(port string) (*Server, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	slog.SetDefault(logger)
	s := &Server{
		Router: mux.NewRouter(),
		Logger: logger,
	}
	err := s.initDB()
	if err != nil {
		return nil, err
	}
	s.initRoutes()
	s.initHttpServer(port)
	return s, nil
}

func main() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	portFlag := flag.String("port", "8000", "port to run http server on")
	flag.Parse()
	go func() {
		osCall := <-interruptChan
		log.Printf("received syscall: %+v\n", osCall)
		cancel()
	}()

	if err := serve(ctx, ":"+*portFlag); err != nil {
		log.Printf("failed to serve due to err %+v\n", err)
	}
}

// serve serves http requests on port `port` and gracefully shuts down the server after waiting for `shutdownTimeout`
func serve(ctx context.Context, port string) error {
	s, err := NewServer(port)
	if err != nil {
		return err
	}
	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve error %+v\n", err)
		}
	}()

	log.Printf("Server has started on port: %+v\n", port)
	<-ctx.Done()
	log.Println("gracefully exiting server")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer func() {
		cancel()
	}()

	err = s.HttpServer.Shutdown(ctxShutdown)
	if err != nil {
		log.Fatalf("server shutdown failed %+v\n", err)
	}

	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}
	return err
}
