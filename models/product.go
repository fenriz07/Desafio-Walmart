package models

/*Product modelo del producto*/
type Product struct {
	ID          int    `bson:"product_id"   json:"id,omitempty"`
	Brand       string `bson:"brand" json:"brand,omitempty"`
	Description string `bson:"description" json:"description,omitempty"`
	Image       string `bson:"image" json:"image,omitempty"`
	Price       int    `bson:"price" json:"price,omitempty"`
}

/*ApplyDiscount aplica un descuento al producto*/
func (p *Product) ApplyDiscount() {

	priceTmp := float64(p.Price)

	discount := 0.50

	dicountPrice := priceTmp * discount

	result := (priceTmp - dicountPrice)

	resultInt := int(result)

	p.Price = resultInt

}

/*IsProductPalindrome Comprueba si es el producto es palindromo*/
func (p Product) IsProductPalindrome() bool {
	if p.isPalindrome(p.Brand) == true || p.isPalindrome(p.Description) == true {
		return true
	}

	return false
}

func (p Product) isPalindrome(str string) bool {
	if str == p.reverse(str) {
		return true
	}
	return false
}

func (p Product) reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
