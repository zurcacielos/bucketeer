package db

type Database struct {

}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	return db, nil
}