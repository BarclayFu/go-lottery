package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		userName  string = "root"
		password  string = "root"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "lottery"
		charset   string = "utf8mb4"
	)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
}
