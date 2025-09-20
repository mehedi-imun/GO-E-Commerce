package db

import (
	"ecommace/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


func GetConnectionString(cnf *config.DBConfig) string {
	sslMode := "disable"
	if cnf.EnableSSLMode {
		sslMode = "enable"
	}

	// key=value format (standard for pg driver)
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
		sslMode,
	)
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon,err:=sqlx.Connect("postgres",dbSource)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	return dbCon,nil


}
