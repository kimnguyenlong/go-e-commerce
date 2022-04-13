package entities

type Review struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Content    string `json:"content" bson:"content,omitempty"`
	CustomerID string `json:"customerId" bson:"customerId,omitempty"`
	ProductID  string `json:"productId" bson:"productId,omitempty"`
}
