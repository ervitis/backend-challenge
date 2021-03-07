package model

type (
	Product struct {
		ProductID string
		Price     float64
		Name      string
		Quantity  int
	}

	Order struct {
		ID       int64
		UserID   int64
		Products []Product
	}
)
