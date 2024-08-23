package dto

import (
	"go-admin/core/dto"
	"time"
)

type ModelQueryReq struct {
	dto.Pagination `search:"-"`
	ModelName      string `form:"modelName"  search:"type:contains;column:model_name;table:vehicle_models" comment:"车型名称"`
	BrandId        int64  `form:"brandId"  search:"type:exact;column:brand_id;table:vehicle_models" comment:"关联的品牌ID (参考vehicle_brands表的主键)"`
	TypeId         int64  `form:"typeId"  search:"type:exact;column:type_id;table:vehicle_models" comment:"关联的车辆种类ID (参考vehicle_types表的主键)"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:vehicle_models" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:vehicle_models" comment:"创建时间"`
	ModelOrder
}

type ModelOrder struct {
	IdOrder          int64      `form:"idOrder"  search:"type:order;column:id;table:vehicle_models"`
	ModelNameOrder   string     `form:"modelNameOrder"  search:"type:order;column:model_name;table:vehicle_models"`
	DescriptionOrder string     `form:"descriptionOrder"  search:"type:order;column:description;table:vehicle_models"`
	BrandIdOrder     int64      `form:"brandIdOrder"  search:"type:order;column:brand_id;table:vehicle_models"`
	TypeIdOrder      int64      `form:"typeIdOrder"  search:"type:order;column:type_id;table:vehicle_models"`
	CreateByOrder    int64      `form:"createByOrder"  search:"type:order;column:create_by;table:vehicle_models"`
	UpdateByOrder    int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:vehicle_models"`
	CreatedAtOrder   *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:vehicle_models"`
	UpdatedAtOrder   *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:vehicle_models"`
}

func (m *ModelQueryReq) GetNeedSearch() interface{} {
	return *m
}

type ModelInsertReq struct {
	ModelName   string `json:"modelName" comment:"车型名称"`
	Description string `json:"description" comment:"车型描述（可选）"`
	BrandId     int64  `json:"brandId" comment:"关联的品牌ID (参考vehicle_brands表的主键)"`
	TypeId      int64  `json:"typeId" comment:"关联的车辆种类ID (参考vehicle_types表的主键)"`
	CurrUserId  int64  `json:"-" comment:"当前登陆用户"`
}

type ModelUpdateReq struct {
	Id          int64  `json:"-" uri:"id" comment:"主键，自动递增"` // 主键，自动递增
	ModelName   string `json:"modelName" comment:"车型名称"`
	Description string `json:"description" comment:"车型描述（可选）"`
	BrandId     int64  `json:"brandId" comment:"关联的品牌ID (参考vehicle_brands表的主键)"`
	TypeId      int64  `json:"typeId" comment:"关联的车辆种类ID (参考vehicle_types表的主键)"`
	CurrUserId  int64  `json:"-" comment:"当前登陆用户"`
}

// ModelGetReq 功能获取请求参数
type ModelGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// ModelDeleteReq 功能删除请求参数
type ModelDeleteReq struct {
	Ids []int64 `json:"ids"`
}
