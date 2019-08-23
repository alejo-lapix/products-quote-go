package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	loc "github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	Name  *string
	Email *string
	Phone *string
}

type Location struct {
	Country *loc.Country
	Zone    *loc.Zone
}

type RelatedProducts struct {
	PrimaryProduct     *ProductRelation
	AssociatedProducts []*ProductRelation
}

type ProductRelation struct {
	Product *products.Product
	Amount  *float64
}

type Notificated struct {
	Experts []*responsibles.User
	Sellers []*responsibles.User
}

type Quote struct {
	ID              *string
	Customer        *Customer
	Location        *Location
	RelatedProducts *RelatedProducts
	Notificated     *Notificated
	CreatedAt       *string
}

type QuoteRepository interface {
	Store(*Quote) error
}

func NewQuote(customer *Customer, location *Location, relatedProducts *RelatedProducts, notificated *Notificated) *Quote {
	id := uuid.New().String()
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
