package event

type Event struct {
	ID       int
	Title    string
	Date     string
	Location string
}

type Repository interface {
	CreateEvent(event *Event) (*Event, error)
	GetEvents() ([]*Event, error)
}

type UseCase interface {
	CreateEvent(title, date, location string) (*Event, error)
	GetEvents() ([]*Event, error)
}
