package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/alejo-lapix/related-products-go/pkg/groups"
)

type QuoteService struct {
	userRepository    responsibles.UserRepository
	groupRepository   groups.GroupRepository
	productRepository products.ProductRepository
}

func (service QuoteService) NewQuote(primaryProduct *products.Product, amount *float64, customer *Customer, location *Location) (*Quote, error) {
	relatedProducts := &RelatedProducts{PrimaryProduct: &ProductRelation{
		Product: primaryProduct,
		Amount:  amount,
	}}
	notification := &Notificated{Sellers: location.Zone.Sellers}
	group, err := service.groupRepository.FindByProduct(primaryProduct.ID)

	if err != nil {
		return nil, err
	}

	relatedProducts.AssociatedProducts = make([]*ProductRelation, len(group.Associations))

	for _, association := range group.Associations {
		relationAmount := *amount * *association.Ratio * *association.Product.UnitOfMeasurement.Quantity

		relatedProducts.AssociatedProducts = append(relatedProducts.AssociatedProducts, &ProductRelation{
			Product: association.Product,
			Amount:  &relationAmount,
		})
	}

	notification.Experts, err = service.userRepository.FindByProductIDs([]*string{primaryProduct.ID})

	if err != nil {
		return nil, err
	}

	return NewQuote(customer, location, relatedProducts, notification), nil
}
