package inventory

import "go.mongodb.org/mongo-driver/bson/primitive"

//Item ...
type Item struct {
	ID           primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	Code         string   `json:"code,omitempty"`
	UnitCost     float32  `json:"unit_cost,omitempty"`
	Description  string   `json:"description,omitempty"`
	SellingPrice string   `json:"selling_price,omitempty"`
	Quantity     string   `json:"quantity,omitempty"`
	Comments     []string `json:"comments,omitempty"`
}

//DTO ... data transfer object for item
type DTO struct {
	Code         string `json:"code,omitempty"`
	Description  string `json:"description,omitempty"`
	SellingPrice string `json:"selling_price,omitempty"`
}
