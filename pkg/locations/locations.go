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

type ZonesByProductID struct {
	ProductID *string   `json:"productId"`
	ZoneIDs   []*string `json:"zonesIds"`
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
	FindByProduct(ID *string) ([]*Zone, error)
	FindByCountry(ID *string) ([]*Zone, error)
	ProductsIDsByZone(ID *string) ([]*string, error)
	Remove(ID *string) error
	Store(*Zone) error
	Update(ID *string, zone *Zone) error
}

type ZonesByProductIDRepository interface {
	Store(*ZonesByProductID) error
	Remove(productID *string) error
	Find(productID *string) (*ZonesByProductID, error)
	FindByZone(ID *string) ([]*ZonesByProductID, error)
}
