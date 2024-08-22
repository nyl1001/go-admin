package dto

import (
	"go-admin/core/dto"
	"time"
)

type FleetQueryReq struct {
	dto.Pagination `search:"-"`
	CompanyId      int64  `form:"companyId"  search:"type:exact;column:company_id;table:fleets" comment:"关联的公司ID (参考companies表的主键)"`
	FleetName      string `form:"fleetName"  search:"type:contains;column:fleet_name;table:fleets" comment:"车队名称"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:fleets" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:fleets" comment:"创建时间"`
	FleetOrder
}

type FleetOrder struct {
	CompanyIdOrder     int64      `form:"companyIdOrder"  search:"type:order;column:company_id;table:fleets"`
	ContactNumberOrder string     `form:"contactNumberOrder"  search:"type:order;column:contact_number;table:fleets"`
	CreateByOrder      int64      `form:"createByOrder"  search:"type:order;column:create_by;table:fleets"`
	CreatedAtOrder     *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:fleets"`
	DescriptionOrder   string     `form:"descriptionOrder"  search:"type:order;column:description;table:fleets"`
	EmailOrder         string     `form:"emailOrder"  search:"type:order;column:email;table:fleets"`
	FleetNameOrder     string     `form:"fleetNameOrder"  search:"type:order;column:fleet_name;table:fleets"`
	IdOrder            int64      `form:"idOrder"  search:"type:order;column:id;table:fleets"`
	ManagerNameOrder   string     `form:"managerNameOrder"  search:"type:order;column:manager_name;table:fleets"`
	UpdateByOrder      int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:fleets"`
	UpdatedAtOrder     *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fleets"`
}

func (m *FleetQueryReq) GetNeedSearch() interface{} {
	return *m
}

type FleetInsertReq struct {
	CompanyId     int64  `json:"companyId" comment:"关联的公司ID (参考companies表的主键)"`
	ContactNumber string `json:"contactNumber" comment:"联系电话"`
	Description   string `json:"description" comment:"车队描述（可选）"`
	Email         string `json:"email" comment:"负责人邮箱"`
	FleetName     string `json:"fleetName" comment:"车队名称"`
	ManagerName   string `json:"managerName" comment:"负责人"`
	CurrUserId    int64  `json:"-" comment:"当前登陆用户"`
}

type FleetUpdateReq struct {
	CompanyId     int64  `json:"companyId" comment:"关联的公司ID (参考companies表的主键)"`
	ContactNumber string `json:"contactNumber" comment:"联系电话"`
	Description   string `json:"description" comment:"车队描述（可选）"`
	Email         string `json:"email" comment:"负责人邮箱"`
	FleetName     string `json:"fleetName" comment:"车队名称"`
	Id            int64  `json:"-" uri:"id" comment:"主键，自动递增"` // 主键，自动递增
	ManagerName   string `json:"managerName" comment:"负责人"`
	CurrUserId    int64  `json:"-" comment:"当前登陆用户"`
}

// FleetGetReq 功能获取请求参数
type FleetGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// FleetDeleteReq 功能删除请求参数
type FleetDeleteReq struct {
	Ids []int64 `json:"ids"`
}

// FleetListByCompanyIDReq 功能获取请求参数
type FleetListByCompanyIDReq struct {
	CompanyId int64 ` json:"companyId"`
}
