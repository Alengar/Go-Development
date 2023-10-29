package event

type EventRepository interface {
	GetAllEvents() ([]*Event, error)
	GetEventByID(id int) (*Event, error)
	CreateEvent(event *Event) (*Event, error)
	UpdateEvent(event *Event) (*Event, error)
	DeleteEvent(id int) error
}
