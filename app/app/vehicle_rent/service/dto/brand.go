package dto

import (
	"go-admin/core/dto"
	"time"
)

type BrandQueryReq struct {
	dto.Pagination `search:"-"`
	BrandName      string `form:"brandName"  search:"type:contains;column:brand_name;table:vehicle_brands" comment:"车辆品牌名称"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:vehicle_brands" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:vehicle_brands" comment:"创建时间"`
	BrandOrder
}

type BrandOrder struct {
	BrandNameOrder string     `form:"brandNameOrder"  search:"type:order;column:brand_name;table:vehicle_brands"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:vehicle_brands"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:vehicle_brands"`
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:vehicle_brands"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:vehicle_brands"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:vehicle_brands"`
}

func (m *BrandQueryReq) GetNeedSearch() interface{} {
	return *m
}

type BrandInsertReq struct {
	BrandName  string `json:"brandName" comment:"车辆品牌名称"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type BrandUpdateReq struct {
	BrandName  string `json:"brandName" comment:"车辆品牌名称"`
	Id         int64  `json:"-" uri:"id" comment:"主键，自动递增"` // 主键，自动递增
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// BrandGetReq 功能获取请求参数
type BrandGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// BrandDeleteReq 功能删除请求参数
type BrandDeleteReq struct {
	Ids []int64 `json:"ids"`
}
