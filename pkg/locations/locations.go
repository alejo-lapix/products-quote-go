package locations

import (
	"products-quote-go/pkg/responsibles"
)

type Country struct {
	ID   *string
	Name *string
}

type Zone struct {
	ID      *string
	Name    *string
	Country *Country
	Sellers []*responsibles.User
}

type ZonesByProductID struct {
	ProductID *string
	Zones     []*Zone
}

type CountryRepository interface {
	Find(ID *string) (*Country, error)
	Remove(ID *string) error
	Store(*Country) error
	Update(ID *string, country *Country) error
}

type ZoneRepository interface {
	Find(ID *string) (*Zone, error)
	Remove(ID *string) error
	Store(*Zone) error
	Update(ID *string, zone *Zone) error
}

type LocationRepository struct {
	CountryRepository CountryRepository
	ZoneRepository    ZoneRepository
}

func (repository *LocationRepository) FindCountry(ID *string) (*Country, error) {
	return repository.CountryRepository.Find(ID)
}

func (repository *LocationRepository) RemoveCountry(ID *string) error {
	return repository.CountryRepository.Remove(ID)
}

func (repository *LocationRepository) StoreCountry(country *Country) error {
	return repository.CountryRepository.Store(country)
}

func (repository *LocationRepository) UpdateCountry(ID *string, country *Country) error {
	return repository.CountryRepository.Update(ID, country)
}

func (repository *LocationRepository) FindZone(ID *string) (*Zone, error) {
	return repository.ZoneRepository.Find(ID)
}

func (repository *LocationRepository) RemoveZone(ID *string) error {
	return repository.ZoneRepository.Remove(ID)
}

func (repository *LocationRepository) StoreZone(zone *Zone) error {
	return repository.ZoneRepository.Store(zone)
}

func (repository *LocationRepository) UpdateZone(ID *string, zone *Zone) error {
	return repository.ZoneRepository.Update(ID, zone)
}
