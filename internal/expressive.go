package expressive

import (
	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/medis/go-expressive/config"
	"github.com/medis/go-expressive/internal/Template"
	app_middleware "github.com/medis/go-expressive/src/Middlewares"
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
	Template *Template.Template
}

func NewExpressive() *Expressive {
	expressive := &Expressive{}

	template, err := Template.NewTemplate()
	if err != nil {
		log.Fatalln(err)
	}
	expressive.Template = template
	expressive.Logger = initLoggers()
	expressive.Router = chi.NewRouter()

	expressive.Config = config.Load(template)

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
		chi_middleware.RequestID,
		chi_middleware.RealIP,
		chi_middleware.Logger,
		chi_middleware.Recoverer,
		chi_middleware.Timeout(60*time.Second),
		chi_middleware.GetHead,
		app_middleware.SecureHeaders,
	)
}
