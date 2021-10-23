package store

func (s *Store) Ping() (string, error) {
	var ping string
	if err := s.Open(); err != nil {
		return "", err
	}
	defer s.Database.Close()
	result, err := s.Database.Query("PING")
	if err != nil {
		return "", err
	}
	if err := result.Scan(&ping); err != nil {
		return "", err
	}
	return ping, nil
}
