package database

func RunDB() error {
	pool, err := CreateDBPool()
	if err != nil {
		return err
	}

	InitDBEngine(pool)

	return nil
}
