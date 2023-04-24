package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUserIdIsNotValid   = errors.New("This user is not valid")
	ErrCantUpdateUser     = errors.New("")
	ErrCantRemoveItemCart = errors.New("")
	ErrCantGetItem        = errors.New("")
	ErrCantBuyCartItem    = errors.New("the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer()
