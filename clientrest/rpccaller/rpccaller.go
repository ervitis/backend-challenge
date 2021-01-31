package rpccaller

import (
	"github.com/ervitis/logme"
	"google.golang.org/grpc"
)

type (
	rpcCall struct {
		conn *grpc.ClientConn
		log logme.Loggerme
	}

	RpcCaller interface {
		CreateOrder()
		CheckOutOrder()
		AddToBasket()
		GetTotalAmount()
		RemoveAll()
	}
)

func New() RpcCaller {
	return nil
}