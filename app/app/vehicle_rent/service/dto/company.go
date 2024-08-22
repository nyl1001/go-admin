package dto

import (
	"go-admin/core/dto"
	"time"
)

type CompaniesQueryReq struct {
	dto.Pagination `search:"-"`
	CompanyName    string `form:"CompanyName"  search:"type:contains;column:Companies_name;table:companies" comment:"公司名称"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:companies" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:companies" comment:"创建时间"`
	CompaniesOrder
}

type CompaniesOrder struct {
	AddressOrder         string     `form:"addressOrder"  search:"type:order;column:address;table:companies"`
	CompanyNameOrder     string     `form:"CompanyNameOrder"  search:"type:order;column:Companies_name;table:companies"`
	ContactNumberOrder   string     `form:"contactNumberOrder"  search:"type:order;column:contact_number;table:companies"`
	CreateByOrder        int64      `form:"createByOrder"  search:"type:order;column:create_by;table:companies"`
	CreatedAtOrder       *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:companies"`
	DescriptionOrder     string     `form:"descriptionOrder"  search:"type:order;column:description;table:companies"`
	EmailOrder           string     `form:"emailOrder"  search:"type:order;column:email;table:companies"`
	IdOrder              int64      `form:"idOrder"  search:"type:order;column:id;table:companies"`
	ParentCompanyIdOrder int64      `form:"parentCompanyIdOrder"  search:"type:order;column:parent_Companies_id;table:companies"`
	UpdateByOrder        int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:companies"`
	UpdatedAtOrder       *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:companies"`
}

func (m *CompaniesQueryReq) GetNeedSearch() interface{} {
	return *m
}

type CompaniesInsertReq struct {
	Address         string `json:"address" comment:"公司地址"`
	CompanyName     string `json:"CompanyName" comment:"公司名称"`
	ContactNumber   string `json:"contactNumber" comment:"联系电话"`
	Description     string `json:"description" comment:"公司描述（可选）"`
	Email           string `json:"email" comment:"公司邮箱"`
	ParentCompanyId int64  `json:"parentCompanyId" comment:"父公司ID (NULL表示顶级公司)"`
	CurrUserId      int64  `json:"-" comment:"当前登陆用户"`
}

type CompaniesUpdateReq struct {
	Address         string `json:"address" comment:"公司地址"`
	CompanyName     string `json:"CompanyName" comment:"公司名称"`
	ContactNumber   string `json:"contactNumber" comment:"联系电话"`
	Description     string `json:"description" comment:"公司描述（可选）"`
	Email           string `json:"email" comment:"公司邮箱"`
	Id              int64  `json:"-" uri:"id" comment:"主键，自动递增"` // 主键，自动递增
	ParentCompanyId int64  `json:"parentCompanyId" comment:"父公司ID (NULL表示顶级公司)"`
	CurrUserId      int64  `json:"-" comment:"当前登陆用户"`
}

// CompaniesGetReq 功能获取请求参数
type CompaniesGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// CompaniesDeleteReq 功能删除请求参数
type CompaniesDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type CompaniesLabel struct {
	Id       int64            `gorm:"-" json:"id"`
	Label    string           `gorm:"-" json:"label"`
	Children []CompaniesLabel `gorm:"-" json:"children"`
}
