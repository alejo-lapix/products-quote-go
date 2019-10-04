package quotes

import (
	"github.com/alejo-lapix/products-go/pkg/categories"
	"github.com/alejo-lapix/products-go/pkg/products"
	"github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/alejo-lapix/related-products-go/pkg/groups"
)

type QuoteService struct {
	UserRepository     responsibles.UserRepository
	GroupRepository    groups.GroupRepository
	QuoteRepository    QuoteRepository
	ProductRepository  products.ProductRepository
	CategoryRepository categories.CategoryRepository
}

// NewPrivateQuote creates a "contact" quote. This quote does not have zone o notificated users.
func (service QuoteService) NewPrivateQuote(primaryProduct *products.Product, amount *float64, customer *Customer) (*Quote, error) {
	relatedProducts, err := service.relatedProducts(primaryProduct, amount)

	if err != nil {
		return nil, err
	}

	return NewPrivateQuote(customer, relatedProducts), nil
}

// NewQuote Creates a new Quote from a given product, amount, customer and zone.
// The related products are taken from the primary's associated products. The
// product's main category give us the people who is going to receive the notifications.
func (service QuoteService) NewQuote(primaryProduct *products.Product, amount *float64, customer *Customer, zone *locations.Zone) (*Quote, error) {
	notification := &Notificated{}
	relatedProducts, err := service.relatedProducts(primaryProduct, amount)

	if err != nil {
		return nil, err
	}

	parentCategory, err := service.CategoryRepository.FindMainCategory(primaryProduct.CategoryID)

	if err != nil {
		return nil, err
	}

	notification.Experts, err = service.UserRepository.FindByCategoryAndZone(parentCategory.ID, zone.ID)

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

// StoreNewQuote creates a new quote and persist it in the given repository.
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

// TODO: move the logic that retrieve the products to the group repository.
// relatedProducts get the products that are associated with the give product
func (service *QuoteService) relatedProducts(primaryProduct *products.Product, amount *float64) (*RelatedProducts, error) {
	relatedProducts := &RelatedProducts{PrimaryProduct: &ProductRelation{
		Product: primaryProduct,
		Amount:  amount,
	}}
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

	return relatedProducts, nil
}
