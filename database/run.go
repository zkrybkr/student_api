package database

func RunDB() (err error) {

	if err = createDBPool(); err != nil {
		return
	}

	if err = RunMigration(); err != nil {
		return
	}

	return nil
}
