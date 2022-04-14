package entities

type OrderItem struct {
	ProductId string `json:"productId" bson:"productId,omitempty"`
	Quantity  int    `json:"quantity" bson:"quantity,omitempty"`
}

type Order struct {
	ID         string       `json:"id" bson:"_id,omitempty"`
	CustomerId string       `json:"customerId" bson:"customerId,omitempty"`
	Items      []*OrderItem `json:"items" bson:"items,omitempty"`
	Created    int64        `json:"created" bson:"created,omitempty"`
	Updated    int64        `json:"updated" bson:"updated,omitempty"`
	Status     string       `json:"status" bson:"status,omitempty"`
}
