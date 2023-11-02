package location

type LocationRepository interface {
	GetAllLocations() ([]*Location, error)
	GetLocationByID(id int) (*Location, error)
	CreateLocation(location *Location) (*Location, error)
	UpdateLocation(location *Location) (*Location, error)
	DeleteLocation(id int) error
}
