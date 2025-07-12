package models

import "time"

type OrderPickups struct {
	PickupId   int64     `gorm:"column:pickup_id;type:int;comment:接单记录唯一标识;primaryKey;" json:"pickup_id"`                      // 接单记录唯一标识
	OrderId    int64     `gorm:"column:order_id;type:int;comment:订单ID;not null;" json:"order_id"`                              // 订单ID
	DriverId   int64     `gorm:"column:driver_id;type:int;comment:司机ID;not null;" json:"driver_id"`                            // 司机ID
	PickupTime time.Time `gorm:"column:pickup_time;type:timestamp;comment:接单时间;default:CURRENT_TIMESTAMP;" json:"pickup_time"` // 接单时间
	Status     string    `gorm:"column:status;type:enum('已接单', '已拒单');comment:接单状态;not null;" json:"status"`                   // 接单状态
}
