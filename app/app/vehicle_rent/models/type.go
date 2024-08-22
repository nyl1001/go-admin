package models

import (
	"time"
)

type VehicleType struct {
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	TypeName  string     `json:"typeName" gorm:"column:type_name;type:varchar(50);comment:车辆种类名称 (如“小型汽车”)"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (VehicleType) TableName() string {
	return "vehicle_types"
}
