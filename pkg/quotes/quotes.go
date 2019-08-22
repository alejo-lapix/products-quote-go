package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	"github.com/google/uuid"
	"products-quote-go/pkg/locations"
	"products-quote-go/pkg/responsibles"
	"time"
)

type Customer struct {
	Name  *string
	Email *string
	Phone *string
}

type Location struct {
	Country *locations.Country
	Zone    *locations.Zone
}

type RelatedProducts struct {
	PrimaryProduct     *products.Product
	AssociatedProducts []*products.Product
}

type Notificated struct {
	Experts []*responsibles.User
	Sellers []*responsibles.User
}

type Quote struct {
	ID              *string
	Customer        *Customer
	Location        *Location
	RelatedProducts *relatedProducts
	Notificated     *Notificated
	CreatedAt       *string
}

type QuoteRepository interface {
	Store(*Quote) error
}

func NewQuote(customer *Customer, location *Location, relatedProducts *RelatedProducts, notificated *Notificated) *Quote {
	id := uuid.New.String()
	createdAt := time.Now().Format(time.RFC3339)

	return &Quote{
		ID:              &id,
		Customer:        customer,
		Location:        location,
		RelatedProducts: relatedProducts,
		Notificated:     notificated,
		CreatedAt:       &createdAt,
	}
}
