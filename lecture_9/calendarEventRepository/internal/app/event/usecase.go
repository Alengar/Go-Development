package event

type EventUseCase interface {
	GetAllEvents() ([]*Event, error)
	GetEventByID(id int) (*Event, error)
	CreateEvent(event *Event) (*Event, error)
	UpdateEvent(event *Event) (*Event, error)
	DeleteEvent(id int) error
}

type eventUseCase struct {
	eventRepo EventRepository
}

func NewEventUseCase(eventRepo EventRepository) EventUseCase {
	return &eventUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *eventUseCase) GetAllEvents() ([]*Event, error) {
	return uc.eventRepo.GetAllEvents()
}

func (uc *eventUseCase) GetEventByID(id int) (*Event, error) {
	return uc.eventRepo.GetEventByID(id)
}

func (uc *eventUseCase) CreateEvent(event *Event) (*Event, error) {
	return uc.eventRepo.CreateEvent(event)
}

func (uc *eventUseCase) UpdateEvent(event *Event) (*Event, error) {
	return uc.eventRepo.UpdateEvent(event)
}

func (uc *eventUseCase) DeleteEvent(id int) error {
	return uc.eventRepo.DeleteEvent(id)
}
