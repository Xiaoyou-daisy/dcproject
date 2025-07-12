package models

type LxhOrderDetails struct {
	Id        int64  `gorm:"column:id;type:int;primaryKey;" json:"id"`
	OrderCode string `gorm:"column:order_code;type:varchar(50);comment:订单编号;" json:"order_code"`         // 订单编号
	TripKey   string `gorm:"column:trip_key;type:varchar(50);comment:mongdb中行程的key;" json:"trip_key"`    // mongdb中行程的key
	DriverKey string `gorm:"column:driver_key;type:varchar(50);comment:司机实际行程路线的key;" json:"driver_key"` // 司机实际行程路线的key
}
