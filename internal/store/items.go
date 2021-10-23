package store

func (s *Store) CreateItem(item Item) error {
	if err := s.Open(); err != nil {
		return err
	}
	defer s.Database.Close()
	_, err := s.Database.Exec(
		"INSERT INTO items (name, description) VALUES ($1, $2)",
		item.Name,
		item.Description,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetItem(id int) (Item, error) {
	if err := s.Open(); err != nil {
		return Item{}, err
	}
	defer s.Database.Close()
	rows, err := s.Database.Query("SELECT name, description FROM items WHERE id=$1", id)
	if err != nil {
		return Item{}, err
	}
	rows.Next()
	var item Item
	if err := rows.Scan(&item.Name, &item.Description); err != nil {
		return Item{}, err
	}

	if err := rows.Close(); err != nil {
		return Item{}, err
	}

	return item, nil
}

func (s *Store) GetItems() ([]Item, error) {
	if err := s.Open(); err != nil {
		return []Item{}, err
	}
	defer s.Database.Close()
	rows, err := s.Database.Query("SELECT name, description FROM items")
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Name, &item.Description); err != nil {
			return []Item{}, err
		}
		items = append(items, item)
	}

	if err := rows.Close(); err != nil {
		return []Item{}, err
	}

	return items, nil
}
