package exceptions

import "fmt"

type (
	OrderNotFoundErr struct{
		msg string
	}
)

func OrderNotFound(orderID int) OrderNotFoundErr {
	return OrderNotFoundErr{msg: fmt.Sprintf("order %d not found", orderID)}
}

func (e OrderNotFoundErr) Error() string {
	return e.msg
}
