package quotes

type Quote struct {
	ID          *string
	Name        *string
	Email       *string
	Phone       *string
	Country     *Country
	Zone        *Zone
	Notificated []*User
}

type QuoteRepository interface {
	Store(*Quote) error
}
