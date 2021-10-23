package store

func (s *Store) CreateUser(user User) error {
	if err := s.Open(); err != nil {
		return err
	}
	defer s.Database.Close()
	_, err := s.Database.Exec(
		"INSERT INTO users (account, password, role) VALUES ($1, $2, $3)",
		user.Account,
		user.Password,
		user.Role,
	)
	if err != nil {
		return err
	}

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return err
	// }

	// if _, err := s.Database.Exec(
	// 	"INSERT INTO baskets (id_user) VALUES ($1)",
	// 	id,
	// ); err != nil {
	// 	return err
	// }

	return nil
}

func (s *Store) GetUser(id int) (User, error) {
	if err := s.Open(); err != nil {
		return User{}, err
	}
	defer s.Database.Close()
	rows, err := s.Database.Query("SELECT account, password, role FROM users WHERE id=$1", id)
	if err != nil {
		return User{}, err
	}
	rows.Next()
	var user User
	if err := rows.Scan(&user.Account, &user.Password, &user.Role); err != nil {
		return User{}, err
	}

	if err := rows.Close(); err != nil {
		return User{}, err
	}

	return user, nil
}
