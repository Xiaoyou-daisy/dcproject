package models

type LxhPassengers struct {
	Id       int64  `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Name     string `gorm:"column:name;type:varchar(20);comment:姓名;" json:"name"`           // 姓名
	NickName string `gorm:"column:nick_name;type:varchar(20);comment:昵称;" json:"nick_name"` // 昵称
	FileId   int64  `gorm:"column:file_id;type:int;comment:头像文件ID;" json:"file_id"`         // 头像文件ID
	Mobile   string `gorm:"column:mobile;type:char(11);comment:手机号;" json:"mobile"`         // 手机号
	Age      int64  `gorm:"column:age;type:int;comment:年龄;" json:"age"`                     // 年龄
	Sex      string `gorm:"column:sex;type:varchar(2);comment:性别;" json:"sex"`              // 性别
	Mileage  string `gorm:"column:mileage;type:varchar(10);comment:里程数;" json:"mileage"`    // 里程数
}
