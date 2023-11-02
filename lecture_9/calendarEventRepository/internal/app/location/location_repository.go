package location

type LocationRepositoryImpl struct {
	locations map[int]*Location
	lastID    int
}

func NewLocationRepository() LocationRepository {
	return &LocationRepositoryImpl{
		locations: make(map[int]*Location),
		lastID:    0,
	}
}

func (r *LocationRepositoryImpl) GetAllLocations() ([]*Location, error) {
	locations := make([]*Location, 0, len(r.locations))
	for _, location := range r.locations {
		locations = append(locations, location)
	}
	return locations, nil
}

func (r *LocationRepositoryImpl) GetLocationByID(id int) (*Location, error) {
	location, exists := r.locations[id]
	if !exists {
		return nil, ErrLocationNotFound
	}
	return location, nil
}

func (r *LocationRepositoryImpl) CreateLocation(location *Location) (*Location, error) {
	r.lastID++
	location.ID = r.lastID
	r.locations[location.ID] = location
	return location, nil
}

func (r *LocationRepositoryImpl) UpdateLocation(location *Location) (*Location, error) {
	_, exists := r.locations[location.ID]
	if !exists {
		return nil, ErrLocationNotFound
	}
	r.locations[location.ID] = location
	return location, nil
}

func (r *LocationRepositoryImpl) DeleteLocation(id int) error {
	_, exists := r.locations[id]
	if !exists {
		return ErrLocationNotFound
	}
	delete(r.locations, id)
	return nil
}
