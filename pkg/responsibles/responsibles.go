package responsibles

type User struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`

	// Those are the categories where this guy is an EXPERT
	CategoryIDs []*string `json:"categoryIds"`
	ZoneIDs     []*string `json:"zoneIds"`
}

type UserRepository interface {
	Find(ID *string) (*User, error)
	All() ([]*User, error)
	FindMany(ids []*string) ([]*User, error)
	FindByCategoryAndZone(categoryID, zoneID *string) ([]*User, error)
	Store(*User) error
	Remove(ID *string) error
}
