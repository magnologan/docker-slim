package init

import (
	"github.com/docker-slim/docker-slim/pkg/app/master/commands/help"
)

func init() {
	help.RegisterCommand()
}
