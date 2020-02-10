//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/medis/go-expressive/src/App/Handlers"
)

func InitialisePingHandler() *Handlers.PingHandler {
	wire.Build(Handlers.NewPingHandler)
	return &Handlers.PingHandler{}
}
