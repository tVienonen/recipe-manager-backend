package database

import (
	"fmt"

	"github.com/go-ozzo/ozzo-dbx"
	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

var conn *dbx.DB

func GetDB(ctx iris.Context) *dbx.DB {
	if conn == nil {
		user := ctx.Values().GetString("DB_USER")
		pass := ctx.Values().GetString("DB_PASS")
		db := ctx.Values().GetString("DB_NAME")
		host := ctx.Values().GetString("DB_HOST")
		port := ctx.Values().GetString("DB_PORT")
		var err error
		conn, err = dbx.MustOpen("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, db))
		if err != nil {
			panic(err)
		}
	}
	return conn
}
