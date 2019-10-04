package locations

type Country struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type Zone struct {
	ID         *string   `json:"id"`
	Name       *string   `json:"name"`
	CountryID  *string   `json:"countryId"`
	SellersIDs []*string `json:"sellersIds"`
}

type CountryRepository interface {
	Find(ID *string) (*Country, error)
	All() ([]*Country, error)
	Remove(ID *string) error
	Store(*Country) error
	Update(ID *string, country *Country) error
}

type ZoneRepository interface {
	Find(ID *string) (*Zone, error)
	FindMany(ids []*string) ([]*Zone, error)
	FindByCountry(ID *string) ([]*Zone, error)
	Remove(ID *string) error
	Store(*Zone) error
	Update(ID *string, zone *Zone) error
}