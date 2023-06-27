package mocks

type Message struct {
	To    string
	Title string
	Body  string
}

type mockNotifier struct {
	notify func(Message)
}

func (mn mockNotifier) Notify(to string, title, body string) error {
	mn.notify(Message{To: to, Title: title, Body: body})
	return nil
}

func NewMockNotifier(notify func(Message)) mockNotifier {
	return mockNotifier{notify: notify}
}
