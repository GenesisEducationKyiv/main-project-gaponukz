package unit

type message struct {
	to    string
	title string
	body  string
}

type mockNotifier struct {
	notify func(message)
}

func (mn mockNotifier) Notify(to string, title, body string) error {
	mn.notify(message{to: to, title: title, body: body})
	return nil
}

func NewMockNotifier(notify func(message)) mockNotifier {
	return mockNotifier{notify: notify}
}
