package types


type Customer struct {
	ID            string                   `json:"id"`
	Live          bool                     `json:"livemode"`
	Created       int64                    `json:"created"`
	Balance       int64                    `json:"account_balance"`
	Delinquent    bool                     `json:"delinquent"`
	Desc          string                   `json:"description"`
	Email         string                   `json:"email"`
	Meta          map[string]string        `json:"metadata"`
	Deleted       bool                     `json:"deleted"`
	BusinessVatID string                   `json:"business_vat_id"`
}

type Product struct {
	ID                string             `json:"id"`
	Created           int64              `json:"created"`
	Updated           int64              `json:"updated"`
	Live              bool               `json:"livemode"`
	Active            bool               `json:"active"`
	Name              string             `json:"name"`
	Caption           string             `json:"caption"`
	Desc              string             `json:"description"`
	Attrs             []string           `json:"attributes"`
	Shippable         bool               `json:"shippable"`
	Images            []string           `json:"images"`
	Meta              map[string]string  `json:"metadata"`
	URL               string             `json:"url"`
	DeactivateOn      []string           `json:"deactivate_on"`
}

type Order struct {
	ID                     string            `json:"id"`
	Amount                 int64             `json:"amount"`
	AmountReturned         int64             `json:"amount_returned"`
	Application            string            `json:"application"`
	ApplicationFee         int64             `json:"application_fee"`
	Created                int64             `json:"created"`
	Customer               Customer          `json:"customer"`
	Email                  string            `json:"email"`
	Items                  []OrderItem       `json:"items"`
	Live                   bool              `json:"livemode"`
	Meta                   map[string]string `json:"metadata"`
	SelectedShippingMethod *string           `json:"selected_shipping_method"`
	Updated                int64             `json:"updated"`
}

type OrderItem struct {
	Amount      int64              `json:"amount"`
	Description string             `json:"description"`
	Parent      string             `json:"parent"`
	Quantity    int64              `json:"quantity"`
}

type SKU struct {
	ID                string             `json:"id"`
	Created           int64              `json:"created"`
	Updated           int64              `json:"updated"`
	Live              bool               `json:"livemode"`
	Active            bool               `json:"active"`
	Desc              string             `json:"description"`
	Attrs             map[string]string  `json:"attributes"`
	Price             int64              `json:"price"`
	Currency          string             `json:"currency"`
	Image             string             `json:"image"`
	Product           Product            `json:"product"`
	Meta              map[string]string  `json:"metadata"`
}


