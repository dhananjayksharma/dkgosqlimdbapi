package mysql

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root#123PD"
const DB_NAME = "imdbdbdemo"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

type MySQLDbStore struct {
	DB *gorm.DB
}

func DBConn(dbsconn string) (*gorm.DB, error) {
	fmt.Println("dbsconn :", dbsconn)
	// var (
	// // databasename, host, userPass, port, userName string
	// )

	// userName = viper.GetString("database.db_one_readwrite.dbuser")
	// userPass = viper.GetString("database.db_one_readwrite.dbpassword")
	// host = viper.GetString("database.db_one_readwrite.hostname")
	// port = viper.GetString("database.db_one_readwrite.hostport")
	// databasename = viper.GetString("database.db_one_readwrite.dbname")

	// dsn := userName + ":" + userPass + "@tcp" + "(" + host + ":" + port + ")/" + databasename + "?" + "parseTime=true&loc=Local"
	// fmt.Println("dsn : ", dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// dsn := userName + ":" + userPass + "@tcp" + "(" + host + ":" + port + ")/" + databasename + "?" + "parseTime=true&loc=Local"
	// fmt.Println("dsn : ", dsn)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	rawSqlForSetDababase := "USE " + DB_NAME + ";"
	db.Exec(rawSqlForSetDababase)

	return db, nil
}

func GetDbConnect(c *gin.Context) {
	c.JSON(200, gin.H{"message": "In GetDbConnect test"})
}
