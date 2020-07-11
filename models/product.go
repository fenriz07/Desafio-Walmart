package models

/*Product modelo del producto*/
type Product struct {
	ID          int    `bson:"product_id"   json:"id,omitempty"`
	Brand       string `bson:"brand" json:"brand,omitempty"`
	Description string `bson:"description" json:"description,omitempty"`
	Image       string `bson:"image" json:"image,omitempty"`
	Price       int    `bson:"price" json:"price,omitempty"`
}
