package command

type Command interface {
	Run(r *Request) error
	Type() string
}
