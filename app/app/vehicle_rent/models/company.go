package models

import (
	"time"
)

type Companies struct {
	Address         string     `json:"address" gorm:"column:address;type:varchar(255);comment:公司地址"`
	CompanyName     string     `json:"companyName" gorm:"column:company_name;type:varchar(100);comment:公司名称"`
	ContactNumber   string     `json:"contactNumber" gorm:"column:contact_number;type:varchar(20);comment:联系电话"`
	CreateBy        int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	CreatedAt       *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	Description     string     `json:"description" gorm:"column:description;type:text;comment:公司描述（可选）"`
	Email           string     `json:"email" gorm:"column:email;type:varchar(100);comment:公司邮箱"`
	Id              int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	ParentCompanyId *int64     `json:"parentCompanyId" gorm:"column:parent_company_id;type:int;comment:父公司ID (NULL表示顶级公司)"`
	UpdateBy        int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt       *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (Companies) TableName() string {
	return "companies"
}
