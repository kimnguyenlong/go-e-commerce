package entities

type Product struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Title       string   `json:"title" bson:"title,omitempty"`
	Price       float64  `json:"price" bson:"price,omitempty"`
	Categories  []string `json:"categories" bson:"categories,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	ProviderID  string   `json:"providerId" bson:"providerId,omitempty"`
}
