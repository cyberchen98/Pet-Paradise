package common

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

type Table struct {
	GetDB     func() *sqlx.DB
	TableName string
}

func (d *Table) Get(dest interface{}, query string, args ...interface{}) error {
	return d.GetDB().Get(dest, query, args...)
}

func (d *Table) Select(dest interface{}, query string, args ...interface{}) error {
	return d.GetDB().Select(dest, query, args...)
}

func (d *Table) UpdateById(keys []string, id int, args ...interface{}) (sql.Result, error) {
	query := "UPDATE `" + d.TableName + "` SET "
	query += makeUpdaters(keys)
	query += " WHERE id=?"
	args = append(args, time.Now().Format(TIME_FORMAT), id)
	return d.GetDB().Exec(query, args...)
}

func (d *Table) Insert(fields map[string]interface{}) (sql.Result, error) {
	query := "INSERT INTO `" + d.TableName + "`("
	var values []interface{}
	values, query = makeValues(query, fields)
	ret, err := d.GetDB().Exec(query, values...)
	return ret, err
}

func (d *Table) DeleteById(id int) error {
	_, err := d.UpdateById([]string{"is_deleted"}, id, "1")
	if err != nil {
		return err
	}
	return nil
}

func wrapWhere(w string) string {
	if strings.Contains(w, ".") {
		var l []string
		for _, v := range strings.Split(w, ".") {
			l = append(l, "`"+v+"`")
		}
		w = strings.Join(l, ".")
	} else {
		w = "`" + w + "`"
	}
	return w
}

func makeValues(query string, fields map[string]interface{}) ([]interface{}, string) {
	first := true
	quotes := ""
	var values []interface{}
	for k, v := range fields {
		if !first {
			query += ","
			quotes += ","
		} else {
			first = false
		}
		query += wrapWhere(k)
		quotes += "?"
		values = append(values, v)
	}
	query += ") VALUES (" + quotes + ") "
	return values, query
}

func makeSelectors(keys []string) string {
	str := ""
	for i, v := range keys {
		if i > 0 {
			str += ","
		}
		str += v
	}
	return str
}

func makeUpdaters(keys []string) string {
	str := ""
	for i, v := range keys {
		if i > 0 {
			str += ","
		}
		str += v + "=?"
	}
	str += ",update_time=?"
	return str
}

func whereCauseIn(params []string, fieldName string) (string, []interface{}) {
	var whereCause string
	var values []interface{}
	if params != nil {
		whereCause += ` AND ` + fieldName + ` IN (`
		for i, v := range params {
			if i > 0 {
				whereCause += `,`
			}
			whereCause += `?`
			values = append(values, v)
		}
		whereCause += `)`
	}
	return whereCause, values
}
