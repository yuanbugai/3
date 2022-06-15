package main

import "fmt"

//
func main() {
	fmt.Println("hello world")
}

//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}
//
//var db *sql.DB
//
///*func initDB() (err error) {
//	// dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4"
//	// open函数只是验证格式是否正确，并不是创建数据库连接
//	db, err = sql.Open("mysql", dsn)
//
//	if err != nil {
//		return err
//	}
//	// 与数据库建立连接
//	err2 := db.Ping()
//	if err2 != nil {
//		return err2
//	}
//	return nil
//}
//func insertData() {
//	sqlStr := "insert into user_tb1 (username,password,status) values (?,?,?)"
//	r, err := db.Exec(sqlStr, "叫你妈叫", "79", 2)
//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//		return
//	}
//	i2, err2 := r.LastInsertId()
//	if err2 != nil {
//		fmt.Printf("err2: %v\n", err2)
//		return
//	}
//	fmt.Printf("i2: %v\n", i2)
//}
//
//type user struct {
//	id       int
//	username string
//	password string
//	status   int
//}
//
//func queryRowData() {
//	sqlStr := "select Id,username,password,status from user_tb1 where id=?"
//	var u user
//	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password, &u.status)
//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//		return
//	}
//	fmt.Printf("Id:%d username:%v password:%v status:%d", u.id, u.username, u.password, u.status)
//}*/
//type insertData interface {
//	insert(db *gorm.DB, code string, price uint)
//}
//
//type insert struct{}
//
//func (insert) insert(db *gorm.DB, code string, price uint) {
//	//TODO implement me
//	p := Product{
//
//		Code:  code,
//		Price: price,
//	}
//	db.Create(&p)
//
//}
//
///*func insert(db *gorm.DB,code string,price uint)  {
//	p := Product{
//
//		Code:  code  ,
//		Price: price,
//	}
//	db.Create(&p)
//}*/
//
//func queryfromprimary(db *gorm.DB, i int) {
//	var products Product
//	db.First(&products, i)
//	fmt.Println(products)
//}
//func main() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4&parseTime=true&loc=Local"
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	queryfromprimary(db, 1)
//	//queryfromprimary(db, 1)
//	// 迁移 schema
//	//db.AutoMigrate(&Product{})
//	//p := Product{
//	//
//	//	Code:  "10086",
//	//	Price: 100,
//	//}
//	//// 插入数据
//	//db.Create(&p)
//	//db.Create(&Product{Code: "D42", Price: 100})
//
//	//
//	//// Read
//	var product Product
//	db.First(&product, 1) // 根据整型主键查找
//	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
//
//	//// Update - 将 product 的 price 更新为 200
//	/*db.Model(&product).Update("Price", 123456)
//	db.First(&product, 1)*/
//	//// Update - 更新多个字段
//	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
//	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
//	//
//	//// Delete - 删除 product
//	db.Delete(&product, 1)
//}
//*/
