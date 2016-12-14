package notify

type Client interface {
	Send(msg string) error
}
