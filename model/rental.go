package model

type Rental struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	OwnerName string `json:"ownerName"`
	Location string `json:"location"`
	Price string `json:"price"`
	DateCreated string `json:"dateCreated"`
}
