package store

// User model for db
type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Item model for db
type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Basket model for db
type Basket struct {
	UserId int `json:"id_user"`
	ItemId int `json:"id_item"`
}

type UserBasket struct {
	Items []Item
}

type BasketUpdate struct {
	Basket
	Action string `json:"action"`
}
