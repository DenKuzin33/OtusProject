package bannersrotator

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

var db = initConnection()

func initConnection() *sql.DB {
	dsn := "host=localhost port=5432 dbname=banners_rotator user=user password=user"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("failed to load driver: %v", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("failed to connect to db: %v", err)
	}
	return db
}

func GetSlot(id int) (Entity, error) {
	query := "select * from slots where id=$1"
	return getEntity(query, id)
}

func GetBanner(id int) (Entity, error) {
	query := "select * from banners where id=$1"
	return getEntity(query, id)
}

func GetDemGroup(id int) (Entity, error) {
	query := "select * from dem_groups where id=$1"
	return getEntity(query, id)
}

func getEntity(query string, id int) (Entity, error) {
	result := Entity{}
	rows, err := db.Query(query, id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&result.id, &result.description); err != nil {
			return result, err
		}
	}
	return result, nil
}
