package models

import (
	"time"
)

type Fleet struct {
	CompanyId     int64      `json:"companyId" gorm:"column:company_id;type:int;comment:关联的公司ID (参考companies表的主键)"`
	ContactNumber string     `json:"contactNumber" gorm:"column:contact_number;type:varchar(20);comment:联系电话"`
	CreateBy      int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	Description   string     `json:"description" gorm:"column:description;type:text;comment:车队描述（可选）"`
	Email         string     `json:"email" gorm:"column:email;type:varchar(100);comment:负责人邮箱"`
	FleetName     string     `json:"fleetName" gorm:"column:fleet_name;type:varchar(100);comment:车队名称"`
	Id            int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	ManagerName   string     `json:"managerName" gorm:"column:manager_name;type:varchar(100);comment:负责人"`
	UpdateBy      int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt     *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (Fleet) TableName() string {
	return "fleets"
}
