package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type PickupLocation struct {
	Id            int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	CompanyId     int64           `json:"companyId" gorm:"column:company_id;type:int;comment:关联的公司ID (参考companies表的主键)"`
	LocationName  string          `json:"locationName" gorm:"column:location_name;type:varchar(100);comment:提车点名称"`
	Description   string          `json:"description" gorm:"column:description;type:text;comment:车型描述（可选）"`
	Address       string          `json:"address" gorm:"column:address;type:varchar(255);comment:提车点地址"`
	ContactNumber string          `json:"contactNumber" gorm:"column:contact_number;type:varchar(20);comment:提车点联系电话"`
	ManagerName   string          `json:"managerName" gorm:"column:manager_name;type:varchar(100);comment:提车点负责人"`
	Latitude      decimal.Decimal `json:"latitude" gorm:"column:latitude;type:decimal(10,7);comment:提车点纬度"`
	Longitude     decimal.Decimal `json:"longitude" gorm:"column:longitude;type:decimal(10,7);comment:提车点经度"`
	CreateBy      int64           `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy      int64           `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt     *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt     *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (PickupLocation) TableName() string {
	return "pickup_locations"
}
