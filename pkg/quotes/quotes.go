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
	Zone            *loc.Zone
	RelatedProducts *RelatedProducts
	Notificated     *Notificated
	CreatedAt       *string
}

func (quote *Quote) NotificationEmails() []*string {
	emails := make([]*string, len(quote.Notificated.Sellers)+len(quote.Notificated.Experts))

	for _, seller := range quote.Notificated.Sellers {
		emails = append(emails, seller.Email)
	}

	for _, expert := range quote.Notificated.Experts {
		emails = append(emails, expert.Email)
	}

	return emails
}

type QuoteRepository interface {
	Store(*Quote) error
}

func NewQuote(customer *Customer, zone *loc.Zone, relatedProducts *RelatedProducts, notificated *Notificated) *Quote {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)

	return &Quote{
		ID:              &id,
		Customer:        customer,
		Zone:            zone,
		RelatedProducts: relatedProducts,
		Notificated:     notificated,
		CreatedAt:       &createdAt,
	}
}
