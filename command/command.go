package command

import (
	"backapper-client/params"
)

type Command interface {
	Run(p *params.Params) error
	Type() string
}
