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
	UserId int `json:"user_id"`
	ItemId int `json:"item_id"`
}

type UserBasket struct {
	Items []Item
}
