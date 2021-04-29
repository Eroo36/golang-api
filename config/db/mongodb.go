package db

import (
	"log"
	"os"

	"github.com/goonode/mogo"
)

var mongoConnection *mogo.Connection = nil

//GetConnection is for get mongo connection
func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		connectionString := os.Getenv("DB_CONNECTION_STRING")
		dbName := os.Getenv("DB_NAME")
		config := &mogo.Config{
			ConnectionString: connectionString,
			Database:         dbName,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}
