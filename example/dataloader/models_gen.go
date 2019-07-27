// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package dataloader

type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	Country string `json:"country"`

	Loader *AddressLoader
}

type Item struct {
	OrderID int
	Name string `json:"name"`

	Loader *ItemSliceLoader
}
