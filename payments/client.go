package payments

import (
	"github.com/vladryk/payments_task/payments/types"
	"github.com/vladryk/payments_task/payments/backends/stripe"
)

type (
	CustomersClient interface {
		CreateCustomer(email, desc string) (*types.Customer, error)
		GetCustomer(id string) (*types.Customer, error)
		ListCustomers() []*types.Customer
	}

	ProductsClient interface {
		ListProducts() []*types.Product
		GetProduct(id string) (*types.Product, error)
		CreateProduct(n, d string, l []string, s *bool) (*types.Product, error)
	}
)

type Client interface {
	CustomersClient
	ProductsClient

	CreateOrder(customerEmail, customerToken string, orderItems []types.OrderItem) (*types.Order, error)
}

func NewClient(key string) Client{
	return stripe.NewClient(key)
}

