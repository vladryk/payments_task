package main

import (
	"github.com/vladryk/payments_task/payments"
	"github.com/pkg/errors"
	"time"
	"fmt"
	"log"
)

const (
	testingEmailTemplate = "r@r%v.ru"
	testingProductTemplate = "test_product_%v"

	key = "sk_test_rD7TfW7yhmpSh5BF6XVqrsLp"
)

var (
	testingEmail string
	testingProduct string
)

func init() {
	timestamp := time.Now().Unix()
	testingEmail = fmt.Sprintf(testingEmailTemplate, timestamp)
	testingProduct = fmt.Sprintf(testingProductTemplate, timestamp)
}

// General library usage examples.
func main()  {
	client := payments.NewClient(key)

	customer, err := client.CreateCustomer(testingEmail, "test_desc")
	if err != nil {
		errors.Wrapf(err, "Can't create customer with %s email", testingEmail)
	}
	log.Println(customer)

	s := false
	products, err := client.CreateProduct(testingProduct, "test_desc", []string{}, &s)
	if err != nil {
		errors.Wrapf(err, "Can't create product")
	}
	log.Println(products)

	customers := client.ListCustomers()
	log.Println(customers)
}
