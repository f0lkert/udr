package processor

import (
	"github.com/f0lkert/udr/internal/database"
	"github.com/f0lkert/udr/pkg/app"
)

type Processor struct {
	app.App
	database.DbConnector
}

func NewProcessor(udr app.App) *Processor {
	return &Processor{
		App:         udr,
		DbConnector: database.NewDbConnector(udr.Config().Configuration.DbConnectorType),
	}
}
