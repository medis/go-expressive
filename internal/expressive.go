package expressive

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/medis/go-expressive/config"
	"log"
	"os"
	"time"
)

type Logger struct {
	AccessLog log.Logger
	AppLog    log.Logger
}

type Expressive struct {
	*Logger
	Router *chi.Mux
	Config *config.Config
}

func NewExpressive() *Expressive {
	expressive := &Expressive{}
	expressive.Config = config.Load()
	expressive.Logger = initLoggers()
	expressive.Router = chi.NewRouter()

	expressive.registerGlobalMiddlewares()
	expressive.registerRoutes()

	return expressive
}

// Initialise loggers.
func initLoggers() *Logger {
	return &Logger{
		AccessLog: *log.New(os.Stdout, "", log.Ldate|log.Ltime),
		AppLog:    *log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Register routes.
func (e *Expressive) registerRoutes() {
	for _, module := range e.Config.Modules {
		for _, route := range module.GetRoutes() {
			e.Router.With(route.Middlewares...).Method(route.Method, route.Path, route.Handler)
		}
	}
}

// Register global middlewares that apply to all routes.
func (e *Expressive) registerGlobalMiddlewares() {
	e.Router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
		middleware.GetHead,
	)
}
