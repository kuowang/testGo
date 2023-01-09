package main

/*
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	//insertData() //插入数据
	//queryOneRow() //查询一条数据
	//queryManyRow() //查询多条语句
	//update() //更改操作
	delete() //删除操作
}

type User struct {
	id       int
	name     string
	password string
	avatar   string
	phone    string
}

func delete() {
	s := "delete from users  where id = ?"
	r, err := db.Exec(s, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Println(i)
	}

}
func update() {
	s := "update users set name= ? where id =1"
	r, err := db.Exec(s, "王括")
	if err != nil {
		fmt.Println(err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Println(i)
	}

}

func queryManyRow() {
	s := "select id,name,password,avatar,phone from users  "
	r, err := db.Query(s)
	var u User
	defer r.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for r.Next() {
			r.Scan(&u.id, &u.name, &u.password, &u.avatar, &u.phone)
			fmt.Println(u)
		}
	}
}

func queryOneRow() {
	s := "select id,name,password,avatar,phone from users where id = ? "
	var u User
	err := db.QueryRow(s, 1).Scan(&u.id, &u.name, &u.password, &u.avatar, &u.phone)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", u)
	}

}

func insertData() {
	sqlStr := "insert into users(name,phone) values (?,?)"
	ret, err := db.Exec(sqlStr, "张三", "15210000000")
	if err != nil {
		fmt.Printf("insert failed, err :%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
func initDB() (err error) {
	// 定义一个初始化数据库的函数func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True" // 不会校验账号密码是否正确// 注意!!!这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	// 尝试与数据库建立连接(校验dsn是否正确)err = db .Ping()
	if err != nil {
		return err
	}
	return nil
}
*/
