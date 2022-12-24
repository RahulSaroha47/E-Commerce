package database

import "errors"

var (
	ErrCantFindProduct = errors.New("cant find the product")
	ErrCantDecodeProduct = errors.New("cant find the product")
	ErrUserIdIsNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("cannot add this product to cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from cart")
	ErrCantGetItem  = errors.New("unable to get item from cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}