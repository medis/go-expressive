package expressive

import "github.com/medis/go-expressive/config"

type Expressive struct {
	Config *config.Config
}

func NewExpressive() *Expressive {
	// Load config.
	config := config.Load()

	return &Expressive{
		Config: config,
	}
}
