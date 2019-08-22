package responsibles

type User struct {
	ID    *string
	Name  *string
	Email *string
	// Those are the category where this guy is an EXPERT
	ProductIDs []*string
}

type UserRepository interface {
	Find(ID *string) (*User, error)
	FindMany(ids []*string) ([]*User, error)
	FindByProductIDs(ids []*string) ([]*User, error)
	Store(*User) error
	Remove(ID *string) error
}
