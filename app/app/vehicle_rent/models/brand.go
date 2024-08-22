package models

import (
	"time"
)

type Brand struct {
	BrandName string     `json:"brandName" gorm:"column:brand_name;type:varchar(100);comment:车辆品牌名称"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (Brand) TableName() string {
	return "vehicle_brands"
}
