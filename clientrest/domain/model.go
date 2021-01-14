package domain

type (
	CreateOrder struct {
		UserID string `json:"userId"`
	}

	AddProductToOrder struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}

	Product struct {
		ProductID string
		Price     string
		Name      string
	}

	Order struct {
		ID       int
		UserID   int
		Products []Product
	}
)
