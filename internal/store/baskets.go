package store

import (
	"errors"
)

func (s *Store) GetBasket(userId int) (UserBasket, error) {
	if err := s.Open(); err != nil {
		return UserBasket{}, err
	}
	defer s.Database.Close()

	rows, err := s.Database.Query("SELECT id_item FROM baskets WHERE id_user=$1", userId)
	if err != nil {
		return UserBasket{}, err
	}

	var itemListId []int
	for rows.Next() {
		var itemId int
		if err := rows.Scan(&itemId); err != nil {
			return UserBasket{}, err
		}
		itemListId = append(itemListId, itemId)
	}

	if err := rows.Close(); err != nil {
		return UserBasket{}, err
	}

	var userBasket UserBasket
	for _, itemId := range itemListId {
		itemsRow, err := s.Database.Query("SELECT name, description FROM items WHERE id=$1", itemId)
		if err != nil {
			return UserBasket{}, err
		}
		for itemsRow.Next() {
			var item Item
			if err := itemsRow.Scan(&item.Name, &item.Description); err != nil {
				return UserBasket{}, err
			}
			userBasket.Items = append(userBasket.Items, item)
		}
		if err := itemsRow.Close(); err != nil {
			return UserBasket{}, err
		}
	}
	return userBasket, nil
}

func (s *Store) GetBaskets() ([]Basket, error) {
	if err := s.Open(); err != nil {
		return []Basket{}, err
	}
	defer s.Database.Close()
	rows, err := s.Database.Query("SELECT id_user, id_item FROM baskets")
	if err != nil {
		return []Basket{}, err
	}

	var baskets []Basket
	for rows.Next() {
		var basket Basket
		if err := rows.Scan(&basket.UserId, &basket.ItemId); err != nil {
			return []Basket{}, err
		}
		baskets = append(baskets, basket)
	}

	if err := rows.Close(); err != nil {
		return []Basket{}, err
	}
	return baskets, nil
}

// Actions: ["add", "dell"]
func (s *Store) UpdateBasket(basket BasketUpdate) error {
	if err := s.Open(); err != nil {
		return err
	}
	defer s.Database.Close()

	switch basket.Action {
	case "add":
		_, err := s.Database.Exec(
			"INSERT INTO baskets (id_user, id_item) VALUES ($1, $2)",
			basket.UserId,
			basket.ItemId,
		)
		if err != nil {
			return err
		}
		return nil
	case "dell":
		_, err := s.Database.Exec(
			"DELETE FROM baskets WHERE id_user=$1 AND id_item=$2",
			basket.UserId,
			basket.ItemId,
		)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("wrong action")
	}
}
