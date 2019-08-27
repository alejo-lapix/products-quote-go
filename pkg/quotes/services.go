package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	"github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/alejo-lapix/related-products-go/pkg/groups"
)

type QuoteService struct {
	userRepository    responsibles.UserRepository
	groupRepository   groups.GroupRepository
	productRepository products.ProductRepository
	quoteRepository   QuoteRepository
}

func (service QuoteService) NewQuote(primaryProduct *products.Product, amount *float64, customer *Customer, zone *locations.Zone) (*Quote, error) {
	relatedProducts := &RelatedProducts{PrimaryProduct: &ProductRelation{
		Product: primaryProduct,
		Amount:  amount,
	}}

	notification := &Notificated{}
	group, err := service.groupRepository.FindByProduct(primaryProduct.ID)

	if err != nil {
		return nil, err
	}

	productIDs := make([]*string, len(group.Associations))
	productIDsMap := make(map[string]*products.Product, len(group.Associations))

	for _, association := range group.Associations {
		productIDsMap[*association.ProductID] = nil
		productIDs = append(productIDs, association.ProductID)
	}

	productList, err := service.productRepository.FindMany(productIDs)

	if err != nil {
		return nil, err
	}

	for _, product := range productList {
		productIDsMap[*product.ID] = product
	}

	relatedProducts.AssociatedProducts = make([]*ProductRelation, len(group.Associations))

	for _, association := range group.Associations {
		relationAmount := *amount * *association.Ratio

		relatedProducts.AssociatedProducts = append(relatedProducts.AssociatedProducts, &ProductRelation{
			Product: productIDsMap[*association.ProductID],
			Amount:  &relationAmount,
		})
	}

	notification.Experts, err = service.userRepository.FindByProductIDs([]*string{primaryProduct.ID})

	if err != nil {
		return nil, err
	}

	notification.Sellers, err = service.userRepository.FindMany(zone.SellersIDs)

	if err != nil {
		return nil, err
	}

	return NewQuote(customer, zone, relatedProducts, notification), nil
}

func (service QuoteService) StoreNewQuote(primaryProduct *products.Product, amount *float64, customer *Customer, zone *locations.Zone) (*Quote, error) {
	quote, err := service.NewQuote(primaryProduct, amount, customer, zone)

	if err != nil {
		return nil, err
	}

	err = service.quoteRepository.Store(quote)

	if err != nil {
		return nil, err
	}

	return quote, nil
}
