package entities

type CartItem struct {
	ProductId string `json:"productId" bson:"productId,omitempty"`
	Quantity  int    `json:"quantity" bson:"quantity,omitempty"`
}

type Cart struct {
	ID         string      `json:"id" bson:"_id,omitempty"`
	CustomerId string      `json:"customerId" bson:"customerId,omitempty"`
	Items      []*CartItem `json:"items" bson:"items,omitempty"`
}
