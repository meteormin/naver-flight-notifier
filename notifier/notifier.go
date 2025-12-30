package notifier

type Notifier interface {
	Send(n Notification) error
}

type Notification interface {
	Query() map[string]any
	Body() []byte
}
