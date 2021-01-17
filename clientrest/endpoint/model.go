package endpoint

type (
	CreateOrder struct {
		UserID int `json:"userId"`
	}

	AddProductToOrder struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}
)
