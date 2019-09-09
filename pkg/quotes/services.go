package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	"github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/alejo-lapix/related-products-go/pkg/groups"
)

type QuoteService struct {
	UserRepository    responsibles.UserRepository
	GroupRepository   groups.GroupRepository
	ProductRepository products.ProductRepository
	QuoteRepository   QuoteRepository
}

func (service QuoteService) NewQuote(primaryProduct *products.Product, amount *float64, customer *Customer, zone *locations.Zone) (*Quote, error) {
	relatedProducts := &RelatedProducts{PrimaryProduct: &ProductRelation{
		Product: primaryProduct,
		Amount:  amount,
	}}

	notification := &Notificated{}
	group, err := service.GroupRepository.FindByProduct(primaryProduct.ID)

	if err != nil {
		return nil, err
	}

	productIDs := make([]*string, len(group.Associations))
	productIDsMap := make(map[string]*products.Product, len(group.Associations))

	for index, association := range group.Associations {
		productIDsMap[*association.ProductID] = nil
		productIDs[index] = group.Associations[index].ProductID
	}

	productList, err := service.ProductRepository.FindMany(productIDs)

	if err != nil {
		return nil, err
	}

	for _, product := range productList {
		productIDsMap[*product.ID] = product
	}

	relatedProducts.AssociatedProducts = make([]*ProductRelation, len(group.Associations))

	for index, association := range group.Associations {
		product := productIDsMap[*association.ProductID]
		relationAmount := *amount * *association.Ratio

		relatedProducts.AssociatedProducts[index] = &ProductRelation{
			Product: product,
			Amount:  &relationAmount,
		}
	}

	notification.Experts, err = service.UserRepository.FindByProductID(primaryProduct.ID)

	if err != nil {
		return nil, err
	}

	if len(zone.SellersIDs) > 0 {
		notification.Sellers, err = service.UserRepository.FindMany(zone.SellersIDs)

		if err != nil {
			return nil, err
		}
	}

	return NewQuote(customer, zone, relatedProducts, notification), nil
}

func (service QuoteService) StoreNewQuote(primaryProduct *products.Product, amount *float64, customer *Customer, zone *locations.Zone) (*Quote, error) {
	quote, err := service.NewQuote(primaryProduct, amount, customer, zone)

	if err != nil {
		return nil, err
	}

	err = service.QuoteRepository.Store(quote)

	if err != nil {
		return nil, err
	}

	return quote, nil
}
