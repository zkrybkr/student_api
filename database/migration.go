package database

import (
	"context"
	"errors"
	c "studentApi/config"
)

func (db *DBConnData) tableExists(tableName string) (exists bool, err error) {
	if err = db.Pool.QueryRow(context.Background(), c.GetDBSQL().TableSelectSQL, tableName).Scan(&exists); err != nil {
		err = errors.New("tablo kontrol hatası, tablo adı: " + tableName)
		return
	}
	return
}

func (db *DBConnData) CreateTable(query, tableName string) error {
	exit, err := DBConn.tableExists(tableName)
	if err != nil {
		return err
	}

	if exit {
		return errors.New("Tablo zaten mevcut, tablo adı: " + tableName)
	}

	if _, err = db.Pool.Exec(context.Background(), query); err != nil {
		return errors.New("Tablo oluşturulurken hata oluştu, tablo adı: " + tableName)
	} else {
		exit, err = DBConn.tableExists(tableName)
		if err != nil {
			return err
		}

		if !exit {
			return errors.New("Oluşturulan tablo bulunamadı, tablo adı: " + tableName)
		}
	}

	return nil
}

func RunMigration() (err error) {
	dbSQL := c.GetDBSQL()
	if err = DBConn.CreateTable(dbSQL.TableCreateUsersSQL, dbSQL.TableUsersName); err != nil {
		return err
	}

	if err = DBConn.CreateTable(dbSQL.TableCreateRolesSQL, dbSQL.TableRolesName); err != nil {
		return err
	}

	return nil
}
