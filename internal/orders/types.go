package orders

type OrderItem struct {
	ProductId int64 `json: productId`
	Quantity  int32 `json: quantity`
}

type createOrderParams struct {
	CustomerId int64       `json: customerId`
	OrderItems []OrderItem `json: orderItems`
}
