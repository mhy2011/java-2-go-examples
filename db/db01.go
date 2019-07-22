package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK = "tcp"
	SERVER = "localhost"
	PORT = 3306
	DATABASE = "go-examples"
)

func main() {

	mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, e := sql.Open("mysql", mysql)
	defer db.Close()
	checkError(e)
	//insert(db)
	//update(db)
	delete(db)
	//query(db)
}

func insert(db *sql.DB) (affectRow int64) {
	stmt, e := db.Prepare("insert into userinfo set username=?, department=?, created=?")
	checkError(e)
	//执行数据库操作
	result, e := stmt.Exec("李四", "研发部", "2019-07-18")
	checkError(e)
	affectRow, _ = result.RowsAffected()
	return
}

func update(db *sql.DB) (affectRow int64) {
	stmt, e := db.Prepare("update userinfo set username=? where uid=?")
	checkError(e)
	//执行数据库操作
	result, e := stmt.Exec("李四01", 2)
	checkError(e)
	affectRow, _ = result.RowsAffected()
	return
}

func delete(db *sql.DB) (affectRow int64) {
	stmt, e := db.Prepare("delete from userinfo where uid=?")
	checkError(e)
	result, e := stmt.Exec(1)
	checkError(e)
	affectRow, _ = result.RowsAffected()
	return
}

func query(db *sql.DB) {
	rows, e := db.Query("select uid, username from userinfo")
	checkError(e)
	for rows.Next()  {
		var uid int
		var username string
		rows.Scan(&uid, &username)
		fmt.Printf("uid=%d, username=%s\n", uid, username)
	}

}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
