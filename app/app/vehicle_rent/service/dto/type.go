package dto

import (
	"go-admin/core/dto"
	"time"
)

type VehicleTypeQueryReq struct {
	dto.Pagination `search:"-"`
	TypeName       string `form:"typeName"  search:"type:contains;column:type_name;table:vehicle_types" comment:"车辆种类名称 (如“小型汽车”)"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:vehicle_types" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:vehicle_types" comment:"创建时间"`
	VehicleTypeOrder
}

type VehicleTypeOrder struct {
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:vehicle_types"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:vehicle_types"`
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:vehicle_types"`
	TypeNameOrder  string     `form:"typeNameOrder"  search:"type:order;column:type_name;table:vehicle_types"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:vehicle_types"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:vehicle_types"`
}

func (m *VehicleTypeQueryReq) GetNeedSearch() interface{} {
	return *m
}

type VehicleTypeInsertReq struct {
	TypeName   string `json:"typeName" comment:"车辆种类名称 (如“小型汽车”)"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type VehicleTypeUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"主键，自动递增"` // 主键，自动递增
	TypeName   string `json:"typeName" comment:"车辆种类名称 (如“小型汽车”)"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// VehicleTypeGetReq 功能获取请求参数
type VehicleTypeGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// VehicleTypeDeleteReq 功能删除请求参数
type VehicleTypeDeleteReq struct {
	Ids []int64 `json:"ids"`
}
