package tests

import (
	"testing"

	"github.com/fenriz07/Desafio-W/models"
)

func TestApplyDiscount(t *testing.T) {

	producto := models.Product{1, "dassad", "zdczs omedat", "www.lider.cl/catalogo/images/electronicsIcon.svg", 986817}

	expected := 493408

	producto.ApplyDiscount()

	if expected != producto.Price {
		t.Errorf("Expected: %v, got: %v", expected, producto.Price)
	}
}

func TestIsProductPalindrome(t *testing.T) {
	producto := models.Product{1, "dassad", "zdczs omedat", "www.lider.cl/catalogo/images/electronicsIcon.svg", 986817}

	expected := true

	isPalindrome := producto.IsProductPalindrome()

	if expected != isPalindrome {
		t.Errorf("Expected: %v, got: %v", expected, isPalindrome)
	}
}
