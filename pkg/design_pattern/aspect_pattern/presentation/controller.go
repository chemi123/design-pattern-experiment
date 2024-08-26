package presentation

type Controller interface {
}

type controller struct {
}

var _ Controller = &controller{}

func NewController() Controller {
	return &controller{}
}
