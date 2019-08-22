package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	"github.com/alejo-lapix/related-products-go/pkg/groups"
	"products-quote-go/pkg/responsibles"
)

type QuoteService struct {
	userRepository    responsibles.UserRepository
	groupRepository   groups.GroupRepository
	productRepository products.ProductRepository
}

func (service QuoteService) NewQuote(primaryProduct *products.Product, customer *Customer, location *Location) (*Quote, error) {
	relatedProducts := &RelatedProducts{PrimaryProduct: primaryProduct}
	noticated := &Notificated{Sellers: location.Zone.Sellers}
	group, err := service.groupRepository.FindByProduct(primaryProduct.ID)

	if err != nil {
		return nil, err
	}

	relatedProducts.AssociatedProducts = make([]*products.Product, len(group.Associations))

	for _, association := range group.Associations {
		relatedProducts.AssociatedProducts = append(relatedProducts.AssociatedProducts, association.Product)
	}

	noticated.Experts, err = service.userRepository.FindByProductIDs([]*string{primaryProduct.ID})

	if err != nil {
		return nil, err
	}

	return NewQuote(customer, location, relatedProducts, noticated), nil
}
