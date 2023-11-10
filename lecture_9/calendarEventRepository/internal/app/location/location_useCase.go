package location

type LocationUseCase interface {
	GetAllLocations() ([]*Location, error)
	GetLocationByID(id int) (*Location, error)
	CreateLocation(location *Location) (*Location, error)
	UpdateLocation(location *Location) (*Location, error)
	DeleteLocation(id int) error
}

type locationUseCase struct {
	locationRepo LocationRepository
}

func NewLocationUseCase(locationRepo LocationRepository) LocationUseCase {
	return &locationUseCase{
		locationRepo: locationRepo,
	}
}

func (uc *locationUseCase) GetAllLocations() ([]*Location, error) {
	return uc.locationRepo.GetAllLocations()
}

func (uc *locationUseCase) GetLocationByID(id int) (*Location, error) {
	return uc.locationRepo.GetLocationByID(id)
}

func (uc *locationUseCase) CreateLocation(location *Location) (*Location, error) {
	return uc.locationRepo.CreateLocation(location)
}

func (uc *locationUseCase) UpdateLocation(location *Location) (*Location, error) {
	return uc.locationRepo.UpdateLocation(location)
}

func (uc *locationUseCase) DeleteLocation(id int) error {
	return uc.locationRepo.DeleteLocation(id)
}
