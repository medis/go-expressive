package expressive

import (
	"github.com/medis/go-expressive/config"
	"log"
	"os"
)

type Logger struct {
	AccessLog log.Logger
	AppLog    log.Logger
}

type Expressive struct {
	*Logger
	Config *config.Config
}

func NewExpressive() *Expressive {
	// Load config.
	config := config.Load()

	return &Expressive{
		Logger: initLoggers(),
		Config: config,
	}
}

func initLoggers() *Logger {
	return &Logger{
		AccessLog: *log.New(os.Stdout, "", log.Ldate|log.Ltime),
		AppLog:    *log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
	}

}
