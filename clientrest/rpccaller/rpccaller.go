package rpccaller

import (
	"context"
	"github.com/ervitis/backend-challenge/clientrest/model"
	protopb "github.com/ervitis/backend-challenge/proto"
	"github.com/ervitis/logme"
	"google.golang.org/grpc"
	"time"
)

type (
	rpcCall struct {
		conn         *grpc.ClientConn
		basketClient protopb.BasketClient
		log          logme.Loggerme
		ctx          context.Context
		cancel       context.CancelFunc
		serverAddr   string
	}

	RpcCaller interface {
		Open() error
		Close() error
		CreateOrder(int) (int, error)
		CheckOutOrder(int) error
		AddToBasket(int, []model.Product) error
		GetTotalAmount(int) (float32, error)
		RemoveAll(int) error
	}
)

func New(serverAddr string, logger logme.Loggerme) RpcCaller {
	return &rpcCall{
		log:        logger,
		serverAddr: serverAddr,
	}
}

func (rpc *rpcCall) Open() error {
	rpc.ctx, rpc.cancel = context.WithTimeout(context.Background(), time.Second*5000)
	if rpc.conn == nil {
		conn, err := grpc.DialContext(rpc.ctx, rpc.serverAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		rpc.conn = conn
	}
	if rpc.basketClient == nil {
		rpc.basketClient = protopb.NewBasketClient(rpc.conn)
	}
	return nil
}

func (rpc *rpcCall) Close() error {
	return rpc.conn.Close()
}

func (rpc *rpcCall) CreateOrder(userID int) (int, error) {
	order := &protopb.NewOrderByUser{UserID: int64(userID)}
	resp, err := rpc.basketClient.Create(rpc.ctx, order)
	if err != nil {
		return -1, err
	}
	return int(resp.OrderID), nil
}

func (rpc *rpcCall) CheckOutOrder(orderID int) error {
	order := &protopb.NewOrderCreated{OrderID: int64(orderID)}
	_, err := rpc.basketClient.Checkout(rpc.ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (rpc *rpcCall) AddToBasket(orderID int, items []model.Product) error {
	var carItems []*protopb.Item
	for _, item := range items {
		ci := &protopb.Item{
			Quantity: int64(item.Quantity),
			ProductID: item.ProductID,
		}
		carItems = append(carItems, ci)
	}
	cart := &protopb.AddItemsToOrder{OrderID: int64(orderID), Items: carItems}
	_, err := rpc.basketClient.AddToBasket(rpc.ctx, cart)
	if err != nil {
		return err
	}
	return nil
}

func (rpc *rpcCall) GetTotalAmount(orderID int) (float32, error) {
	order := &protopb.NewOrderCreated{OrderID: int64(orderID)}
	resp, err := rpc.basketClient.GetTotalAmount(rpc.ctx, order)
	if err != nil {
		return -1, err
	}
	return resp.Amount, nil
}

func (rpc *rpcCall) RemoveAll(orderID int) error {
	order := &protopb.NewOrderCreated{OrderID: int64(orderID)}
	_, err := rpc.basketClient.RemoveAll(rpc.ctx, order)
	return err
}
