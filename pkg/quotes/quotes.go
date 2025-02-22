package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	loc "github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/google/uuid"
	"time"
)

type QuoteType string

const (
	PUBLIC  QuoteType = "public"
	PRIVATE QuoteType = "private"
)

type Customer struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

type RelatedProducts struct {
	PrimaryProduct     *ProductRelation   `json:"primaryProduct"`
	AssociatedProducts []*ProductRelation `json:"associatedProducts"`
}

type ProductRelation struct {
	Product *products.Product `json:"product"`
	Amount  *float64          `json:"amount"`
}

func (relation *ProductRelation) Total() float64 {
	return *relation.Product.Price * *relation.Amount
}

type Notificated struct {
	Experts []*responsibles.User `json:"experts"`
	Sellers []*responsibles.User `json:"sellers"`
}

type Quote struct {
	ID              *string          `json:"id"`
	Customer        *Customer        `json:"customer"`
	Zone            *loc.Zone        `json:"zone"`
	RelatedProducts *RelatedProducts `json:"relatedProducts"`
	Type            QuoteType        `json:"type"`
	Notificated     *Notificated     `json:"notificated"`
	CreatedAt       *string          `json:"createdAt"`
}

func (quote *Quote) Total() float64 {
	total := quote.RelatedProducts.PrimaryProduct.Total()

	for _, relatedProduct := range quote.RelatedProducts.AssociatedProducts {
		total += relatedProduct.Total()
	}

	return total
}

func (quote *Quote) NotificationEmails() []*string {
	emails := make([]*string, len(quote.Notificated.Sellers)+len(quote.Notificated.Experts))

	for index, _ := range quote.Notificated.Sellers {
		emails[index] = quote.Notificated.Sellers[index].Email
	}

	for index, _ := range quote.Notificated.Experts {
		emails[len(quote.Notificated.Sellers)+index] = quote.Notificated.Experts[index].Email
	}

	return emails
}

type QuoteRepository interface {
	Store(*Quote) error
	Find(ID *string) (*Quote, error)
	Paginate(year, month, quoteType, lastPage *string) (list []*Quote, nextKey *string, err error)
	All() ([]*Quote, error)
}

func NewQuote(customer *Customer, zone *loc.Zone, relatedProducts *RelatedProducts, notificated *Notificated) *Quote {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)

	return &Quote{
		ID:              &id,
		Customer:        customer,
		Zone:            zone,
		Type:            PUBLIC,
		RelatedProducts: relatedProducts,
		Notificated:     notificated,
		CreatedAt:       &createdAt,
	}
}

func NewPrivateQuote(customer *Customer, relatedProducts *RelatedProducts) *Quote {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)

	return &Quote{
		ID:              &id,
		Customer:        customer,
		Type:            PRIVATE,
		RelatedProducts: relatedProducts,
		CreatedAt:       &createdAt,
	}
}
