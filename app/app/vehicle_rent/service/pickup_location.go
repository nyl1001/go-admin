package service

import (
	"fmt"
	"go-admin/app/app/vehicle_rent/models"
	"go-admin/app/app/vehicle_rent/service/dto"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"time"

	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type PickupLocation struct {
	service.Service
}

//
// NewPickupLocationService
// @Description: 实例化PickupLocation
// @return *PickupLocation
//
func NewPickupLocationService(s *service.Service) *PickupLocation {
	var srv = new(PickupLocation)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

//
// GetPage
// @Description: 获取PickupLocation列表

// @return []models.PickupLocation
// @return int64
// @return int
// @return error
//
func (e *PickupLocation) GetPage(c *dto.PickupLocationQueryReq, p *middleware.DataPermission) ([]models.PickupLocation, int64, int, error) {
	var data models.PickupLocation
	var list []models.PickupLocation
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

//
// Get
// @Description: 获取PickupLocation对象
// @receiver e
// @return *models.PickupLocation
// @return int
// @return error
//
func (e *PickupLocation) Get(id int64, p *middleware.DataPermission) (*models.PickupLocation, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.PickupLocation{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

//
// QueryOne
// @Description: 通过自定义条件获取PickupLocation一条记录
// @receiver e
// @return *models.PickupLocation
// @return error
//
func (e *PickupLocation) QueryOne(queryCondition *dto.PickupLocationQueryReq, p *middleware.DataPermission) (*models.PickupLocation, int, error) {
	data := &models.PickupLocation{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

//
// Count
//  @Description: 获取条数
//  @receiver e
//  @param c
//  @return int64
//  @return int
//  @return error
//
func (e *PickupLocation) Count(queryCondition *dto.PickupLocationQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.PickupLocation{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

//
// Insert
// @Description: 创建PickupLocation对象

// @return int64 插入数据的主键
// @return int
// @return error
//
func (e *PickupLocation) Insert(c *dto.PickupLocationInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	now := time.Now()
	var data models.PickupLocation
	data.CompanyId = c.CompanyId
	data.LocationName = c.LocationName
	data.Description = c.Description
	data.Address = c.Address
	data.ContactNumber = c.ContactNumber
	data.ManagerName = c.ManagerName
	data.Latitude = c.Latitude
	data.Longitude = c.Longitude
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

//
// Update
// @Description: 修改PickupLocation对象

// @return bool 是否有数据更新
// @return error
//
func (e *PickupLocation) Update(c *dto.PickupLocationUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.CompanyId >= 0 && data.CompanyId != c.CompanyId {
		updates["company_id"] = c.CompanyId
	}
	if c.LocationName != "" && data.LocationName != c.LocationName {
		updates["location_name"] = c.LocationName
	}
	if c.Description != "" && data.Description != c.Description {
		updates["description"] = c.Description
	}
	if c.Address != "" && data.Address != c.Address {
		updates["address"] = c.Address
	}
	if c.ContactNumber != "" && data.ContactNumber != c.ContactNumber {
		updates["contact_number"] = c.ContactNumber
	}
	if c.ManagerName != "" && data.ManagerName != c.ManagerName {
		updates["manager_name"] = c.ManagerName
	}
	if c.Latitude.Cmp(decimal.Zero) >= 0 && data.Latitude.Cmp(c.Latitude) != 0 {
		updates["latitude"] = c.Latitude
	}
	if c.Longitude.Cmp(decimal.Zero) >= 0 && data.Longitude.Cmp(c.Longitude) != 0 {
		updates["longitude"] = c.Longitude
	}
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

//
// Remove
// @Description: 删除PickupLocation
// @receiver e
// @return int
// @return error
//
func (e *PickupLocation) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.PickupLocation
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

//
// GetExcel
// @Description: GetExcel 导出PickupLocation excel数据
// @receiver e
// @return []byte
// @return int
// @return error
//
func (e *PickupLocation) GetExcel(list []models.PickupLocation) ([]byte, error) {
	sheetName := "PickupLocation"
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{"主键，自动递增", "关联的公司ID (参考companies表的主键)", "提车点名称", "车型描述（可选）", "提车点地址", "提车点联系电话", "提车点负责人", "提车点纬度", "提车点经度", "创建者", "更新者", "创建时间", "更新时间"})
	// dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{item.Id, item.CompanyId, item.LocationName, item.Description, item.Address, item.ContactNumber, item.ManagerName, item.Latitude, item.Longitude, item.CreateBy, item.UpdateBy, item.CreatedAt, item.UpdatedAt,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
