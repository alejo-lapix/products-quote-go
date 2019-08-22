package locations

import "products-quote-go/pkg/responsibles"

type Country struct {
	ID   *string
	Name *string
}

type Zone struct {
	ID        *string
	Name      *string
	CountryId *string
	SellersID []*string
	sellers   []*responsibles.User
}

func (zone *Zone) Sellers() ([]*responsibles.User, error) {
	return zone.sellers, nil
}

type ZonesByProduct struct {
	ID        *string
	ProductID *string
	ZonesID   []*string
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
