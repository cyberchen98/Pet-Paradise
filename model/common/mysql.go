package common

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var mysql *sqlx.DB

const (
	TBL_USER         = "user"
	TBL_USER_ADDRESS = "user_address"
	TBL_ORDER        = "order"
	TBL_PRODUCT      = "product"
)

func Conn() *sqlx.DB {
	return mysql
}

func connectDatabase(host, user, password, dbname string, port int) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true", user, password, host, port, dbname))
}

func ConfigureMysqlDatabase(host string, port int, user, password, dbname string) error {
	var err error
	mysql, err = connectDatabase(host, user, password, dbname, port)
	return err
}
