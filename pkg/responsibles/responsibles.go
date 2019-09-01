package responsibles

type User struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	// Those are the category where this guy is an EXPERT
	ProductIDs []*string `json:"productIds"`
}

type UserRepository interface {
	Find(ID *string) (*User, error)
	All() (*[]User, error)
	FindMany(ids []*string) ([]*User, error)
	FindByProductID(id *string) ([]*User, error)
	Store(*User) error
	Remove(ID *string) error
}
