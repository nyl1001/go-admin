package dto

import (
	"go-admin/core/dto"
	"time"

	"github.com/shopspring/decimal"
)

type PickupLocationQueryReq struct {
	dto.Pagination `search:"-"`
	CompanyId      int64  `form:"companyId"  search:"type:exact;column:company_id;table:pickup_locations" comment:"关联的公司ID (参考companies表的主键)"`
	LocationName   string `form:"locationName"  search:"type:exact;column:location_name;table:pickup_locations" comment:"提车点名称"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:pickup_locations" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:pickup_locations" comment:"创建时间"`
	PickupLocationOrder
}

type PickupLocationOrder struct {
	IdOrder            int64           `form:"idOrder"  search:"type:order;column:id;table:pickup_locations"`
	CompanyIdOrder     int64           `form:"companyIdOrder"  search:"type:order;column:company_id;table:pickup_locations"`
	LocationNameOrder  string          `form:"locationNameOrder"  search:"type:order;column:location_name;table:pickup_locations"`
	DescriptionOrder   string          `form:"descriptionOrder"  search:"type:order;column:description;table:pickup_locations"`
	AddressOrder       string          `form:"addressOrder"  search:"type:order;column:address;table:pickup_locations"`
	ContactNumberOrder string          `form:"contactNumberOrder"  search:"type:order;column:contact_number;table:pickup_locations"`
	ManagerNameOrder   string          `form:"managerNameOrder"  search:"type:order;column:manager_name;table:pickup_locations"`
	LatitudeOrder      decimal.Decimal `form:"latitudeOrder"  search:"type:order;column:latitude;table:pickup_locations"`
	LongitudeOrder     decimal.Decimal `form:"longitudeOrder"  search:"type:order;column:longitude;table:pickup_locations"`
	CreateByOrder      int64           `form:"createByOrder"  search:"type:order;column:create_by;table:pickup_locations"`
	UpdateByOrder      int64           `form:"updateByOrder"  search:"type:order;column:update_by;table:pickup_locations"`
	CreatedAtOrder     *time.Time      `form:"createdAtOrder"  search:"type:order;column:created_at;table:pickup_locations"`
	UpdatedAtOrder     *time.Time      `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:pickup_locations"`
}

func (m *PickupLocationQueryReq) GetNeedSearch() interface{} {
	return *m
}

type PickupLocationInsertReq struct {
	CompanyId     int64           `json:"companyId" comment:"关联的公司ID (参考companies表的主键)"`
	LocationName  string          `json:"locationName" comment:"提车点名称"`
	Description   string          `json:"description" comment:"车型描述（可选）"`
	Address       string          `json:"address" comment:"提车点地址"`
	ContactNumber string          `json:"contactNumber" comment:"提车点联系电话"`
	ManagerName   string          `json:"managerName" comment:"提车点负责人"`
	Latitude      decimal.Decimal `json:"latitude" comment:"提车点纬度"`
	Longitude     decimal.Decimal `json:"longitude" comment:"提车点经度"`
	CurrUserId    int64           `json:"-" comment:"当前登陆用户"`
}

type PickupLocationUpdateReq struct {
	Id            int64           `json:"-" uri:"id" comment:"主键，自动递增"` // 主键，自动递增
	CompanyId     int64           `json:"companyId" comment:"关联的公司ID (参考companies表的主键)"`
	LocationName  string          `json:"locationName" comment:"提车点名称"`
	Description   string          `json:"description" comment:"车型描述（可选）"`
	Address       string          `json:"address" comment:"提车点地址"`
	ContactNumber string          `json:"contactNumber" comment:"提车点联系电话"`
	ManagerName   string          `json:"managerName" comment:"提车点负责人"`
	Latitude      decimal.Decimal `json:"latitude" comment:"提车点纬度"`
	Longitude     decimal.Decimal `json:"longitude" comment:"提车点经度"`
	CurrUserId    int64           `json:"-" comment:"当前登陆用户"`
}

// PickupLocationGetReq 功能获取请求参数
type PickupLocationGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// PickupLocationDeleteReq 功能删除请求参数
type PickupLocationDeleteReq struct {
	Ids []int64 `json:"ids"`
}
