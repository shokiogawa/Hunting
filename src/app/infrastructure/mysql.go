package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

//他の場面で繰り返し使用したいため、Mysqlという構造体を作成。
type Mysql struct {
	database *sqlx.DB
}

func NewMysql() (*Mysql, error) {
	sql := new(Mysql)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	fmt.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	sql.database = db
	return sql, err
}

func (mysql *Mysql) Connect() (db *sqlx.DB, err error) {
	err = mysql.database.Ping()
	if err != nil {
		fmt.Println("Mysqlエラー")
		return nil, err
	}
	return mysql.database, nil
}
