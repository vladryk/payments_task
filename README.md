# API for Stripe
 API with the required methods for handling payments in a (general) business

Methods that were implemented:
 - create a customer
 - retrieve a customer
 - retrieve all customers
 - create a product
 - retrieve a product
 - retrieve all products
 - given a customer and a list of products, assign an appropriate payment to the customer


# Example of using (you can find more in main.go):

```
key = "sk_test_rD7TfW7yhmpSh5BF6XVqrsLp"
client := payments.NewClient(key)
customer, err := client.CreateCustomer(testingEmail, "test_desc") // Create customer
```