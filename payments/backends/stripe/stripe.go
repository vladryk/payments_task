package stripe

import (
	"github.com/stripe/stripe-go/sku"
	"github.com/jinzhu/copier"
	"github.com/stripe/stripe-go/order"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/product"
	"github.com/pkg/errors"
	"github.com/vladryk/payments_task/payments/types"
)

type Client struct {}

func NewClient(key string) Client {
	stripe.Key = key
	return Client{}
}

func (a Client) CreateCustomer(email, desc string) (*types.Customer, error) {
	customerParams := &stripe.CustomerParams{
		Desc: desc,
		Email: email,
	}
	c, err := customer.New(customerParams)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't create user with %s email", email)
	}
	customerS := types.Customer{}
	copier.Copy(&customerS, &c)
	return &customerS, nil
}

func (a Client) GetCustomer(id string) (*types.Customer, error) {
	c, err := customer.Get(id, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "User with %s id not found", id)
	}
	customerS := types.Customer{}
	copier.Copy(&customerS, &c)
	return &customerS, nil
}

func (a Client) ListCustomers() []*types.Customer {
	customersList := []*types.Customer{}
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("", "", "")
	i := customer.List(params)

	for i.Next() {
		customerS := types.Customer{}
		c := i.Customer()
		copier.Copy(&customerS, &c)
		customersList = append(customersList, &customerS)
	}
	return customersList
}

func (a Client) ListProducts() []*types.Product {
	productsList := []*types.Product{}
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("", "", "")
	i := product.List(params)
	for i.Next() {
		productS := types.Product{}
		c := i.Product()
		copier.Copy(&productS, &c)
		productsList = append(productsList, &productS)
	}
	return productsList
}

func (a Client) GetProduct(id string) (*types.Product, error) {
	c, err := product.Get(id)
	if err != nil {
		return nil, errors.Wrapf(err, "Product with %s id not found", id)
	}
	productS := types.Product{}
	copier.Copy(&productS, &c)
	return &productS, nil
}

func (a Client) CreateProduct(n, d string, l []string, s *bool) (*types.Product, error) {
	productParams := &stripe.ProductParams{
		Name: n,
		Desc: d,
		Attrs: l,
		Shippable: s,
	}
	c, err := product.New(productParams)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't create product with %s name", n)
	}
	productS := types.Product{}
	copier.Copy(&productS, &c)
	return &productS, nil
}

func (a Client) createSKU (productId string, price, quantity int64) (*types.SKU, error){
	skuParams := &stripe.SKUParams{
		Product: productId,
		Price: price,
		Currency: "usd",
		Inventory: stripe.Inventory{
			Type: "finite",
			Quantity: quantity,
		},
	}
	s, err := sku.New(skuParams)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't create SKU with  for product %s ", productId)
	}
	skuS := types.SKU{}
	copier.Copy(&skuS, &s)
	return &skuS, nil
}

// convertToStripeOrderItems is a helper which simply converts internal order items to Stripe's order items.
func convertToStripeOrderItems(orderItems []types.OrderItem) (result []*stripe.OrderItemParams) {
	for _, v := range orderItems {
		item := &stripe.OrderItemParams{}
		copier.Copy(item, &v)
		item.Type = "sku"
		result = append(result, item)
	}
	return
}

func (a Client) CreateOrder(customerEmail, customerToken string, orderItems []types.OrderItem) (*types.Order, error)  {
	// Convert the data to applicable Stripe type.
	items := convertToStripeOrderItems(orderItems)
	o, err := order.New(&stripe.OrderParams{
		Currency: currency.USD,
		Email: customerEmail,
		Items: items,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "Can't create oder for %s", customerEmail)
	}

	// Use token obtained with Stripe.js
	orderPayParams := &stripe.OrderPayParams{Source: &stripe.SourceParams{Token: customerToken}}
	o, err = order.Pay(
		o.ID,
		orderPayParams,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't create payment for %s order", o.ID)
	}

	// Convert Stripe data back to internal representation.
	orderS := types.Order{}
	copier.Copy(&orderS, &o)
	return &orderS, nil
}

