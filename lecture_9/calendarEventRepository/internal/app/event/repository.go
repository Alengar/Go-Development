package event

type EventRepositoryImpl struct {
	events map[int]*Event
	lastID int
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{
		events: make(map[int]*Event),
		lastID: 0,
	}
}

func (r *EventRepositoryImpl) GetAllEvents() ([]*Event, error) {
	events := make([]*Event, 0, len(r.events))
	for _, event := range r.events {
		events = append(events, event)
	}
	return events, nil
}

func (r *EventRepositoryImpl) GetEventByID(id int) (*Event, error) {
	event, exists := r.events[id]
	if !exists {
		return nil, ErrEventNotFound
	}
	return event, nil
}

func (r *EventRepositoryImpl) CreateEvent(event *Event) (*Event, error) {
	r.lastID++
	event.ID = r.lastID
	r.events[event.ID] = event
	return event, nil
}

func (r *EventRepositoryImpl) UpdateEvent(event *Event) (*Event, error) {
	_, exists := r.events[event.ID]
	if !exists {
		return nil, ErrEventNotFound
	}
	r.events[event.ID] = event
	return event, nil
}

func (r *EventRepositoryImpl) DeleteEvent(id int) error {
	_, exists := r.events[id]
	if !exists {
		return ErrEventNotFound
	}
	delete(r.events, id)
	return nil
}
