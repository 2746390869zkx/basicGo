package main

import (
	model1 "basicGo/gorm/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 创建
/*func main() {
	db, err := gorm.Open("mysql", "root:360428@(localhost)/go?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		println("mysql connect error")
		println(err.Error())
	} else {
		println("connect successfully")
	}

	user := model1.User{Name: "zhang", Age: 18}

	//record := db.NewRecord(user) // 主键为空返回`true`
	//println("record is", record)
	//if record {
	create := db.Create(&user) // 创建user
	if create.Error != nil {
		panic(create.Error)
	}

	//}
	//newRecord := db.NewRecord(user) // 创建`user`后返回`false`
	//println(newRecord)
}*/

// 查询
func main() {
	db, err := gorm.Open("mysql", "root:360428@(localhost)/go?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	user := model1.User{ID: 2}
	db.First(&user) // 会按照已有字段进行where条件拼接
	fmt.Println(user)

	// 随机获取一条记录
	db.Take(&user) // SELECT * FROM users LIMIT 1;

	// 根据主键查询最后一条记录
	db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	var users []model1.User
	// 查询所有的记录
	db.Find(&users)
	// SELECT * FROM users;
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;

}
