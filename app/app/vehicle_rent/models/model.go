package models

import (
	"time"
)

type Model struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	ModelName   string     `json:"modelName" gorm:"column:model_name;type:varchar(100);comment:车型名称"`
	Description string     `json:"description" gorm:"column:description;type:text;comment:车型描述（可选）"`
	BrandId     int64      `json:"brandId" gorm:"column:brand_id;type:int;comment:关联的品牌ID (参考vehicle_brands表的主键)"`
	TypeId      int64      `json:"typeId" gorm:"column:type_id;type:int;comment:关联的车辆种类ID (参考vehicle_types表的主键)"`
	CreateBy    int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy    int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (Model) TableName() string {
	return "vehicle_models"
}
