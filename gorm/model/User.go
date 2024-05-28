package model

// 默认表名就是users
// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
	//Model
	ID   uint `gorm:"primary_key;auto_increment"` // 定义主键
	Age  int
	Name string `gorm:"size:256"`
	//Birthday     *time.Time
	//Email        string  `gorm:"type:varchar(100);unique_index"`
	//Role         string  `gorm:"size:255"`        // 设置字段大小为255
	//MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	//Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	//Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	//IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 使用方法可以指定名字
// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "users"
}

// 通过角色进行表区分
//func (u User) TableName() string {
//	if u.Role == "admin" {
//		return "admin_users"
//	} else {
//		return "users"
//	}
//}

/**
    // 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
*/
