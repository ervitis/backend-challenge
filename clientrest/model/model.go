package model

type (
	Product struct {
		ProductID string
		Price     string
		Name      string
		Quantity  int
	}

	Order struct {
		ID       int
		UserID   int
		Products []Product
	}
)
