package notifier

import "fmt"

type logger interface {
	Info(string)
	Warning(string)
}

type notifier interface {
	Notify(to string, title, body string) error
}

type decorator struct {
	n notifier
	l logger
}

func NewLoggingDecorator(n notifier, l logger) decorator {
	return decorator{n, l}
}

func (d *decorator) Notify(to, title, body string) error {
	err := d.n.Notify(to, title, body)

	if err != nil {
		d.l.Warning(fmt.Sprintf("Failed to send message to %s: %v", to, err))

	} else {
		d.l.Info(fmt.Sprintf("Successfully sent message to %s", to))
	}

	return err
}
