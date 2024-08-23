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

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Fleet struct {
	service.Service
}

// NewFleetService
// @Description: 实例化Fleet
// @return *Fleet
func NewFleetService(s *service.Service) *Fleet {
	var srv = new(Fleet)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

//
// GetPage
// @Description: 获取Fleet列表

// @return []models.Fleet
// @return int64
// @return int
// @return error
func (e *Fleet) GetPage(c *dto.FleetQueryReq, p *middleware.DataPermission) ([]models.Fleet, int64, int, error) {
	var data models.Fleet
	var list []models.Fleet
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

func (e *Fleet) ListByCompanyID(companyID int64, p *middleware.DataPermission) ([]models.Fleet, int, error) {
	if companyID <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var list []models.Fleet
	var data models.Fleet
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Where("company_id=?", companyID).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return list, lang.SuccessCode, nil
}

// Get
// @Description: 获取Fleet对象
// @receiver e
// @return *models.Fleet
// @return int
// @return error
func (e *Fleet) Get(id int64, p *middleware.DataPermission) (*models.Fleet, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.Fleet{}
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

// QueryOne
// @Description: 通过自定义条件获取Fleet一条记录
// @receiver e
// @return *models.Fleet
// @return error
func (e *Fleet) QueryOne(queryCondition *dto.FleetQueryReq, p *middleware.DataPermission) (*models.Fleet, int, error) {
	data := &models.Fleet{}
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

// Count
//
//	@Description: 获取条数
//	@receiver e
//	@param c
//	@return int64
//	@return int
//	@return error
func (e *Fleet) Count(queryCondition *dto.FleetQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.Fleet{}).
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
// @Description: 创建Fleet对象

// @return int64 插入数据的主键
// @return int
// @return error
func (e *Fleet) Insert(c *dto.FleetInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	now := time.Now()
	var data models.Fleet
	data.CompanyId = c.CompanyId
	data.ContactNumber = c.ContactNumber
	data.CreateBy = c.CurrUserId
	data.CreatedAt = &now
	data.Description = c.Description
	data.Email = c.Email
	data.FleetName = c.FleetName
	data.ManagerName = c.ManagerName
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

//
// Update
// @Description: 修改Fleet对象

// @return bool 是否有数据更新
// @return error
func (e *Fleet) Update(c *dto.FleetUpdateReq, p *middleware.DataPermission) (bool, int, error) {
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
	if c.ContactNumber != "" && data.ContactNumber != c.ContactNumber {
		updates["contact_number"] = c.ContactNumber
	}
	if c.Description != "" && data.Description != c.Description {
		updates["description"] = c.Description
	}
	if c.FleetName != "" && data.FleetName != c.FleetName {
		updates["fleet_name"] = c.FleetName
	}
	if c.ManagerName != "" && data.ManagerName != c.ManagerName {
		updates["manager_name"] = c.ManagerName
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

// Remove
// @Description: 删除Fleet
// @receiver e
// @return int
// @return error
func (e *Fleet) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.Fleet
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出Fleet excel数据
// @receiver e
// @return []byte
// @return int
// @return error
func (e *Fleet) GetExcel(list []models.Fleet) ([]byte, error) {
	sheetName := "Fleet"
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{"关联的公司ID (参考companies表的主键)", "联系电话", "创建者", "创建时间", "车队描述（可选）", "负责人邮箱", "车队名称", "主键，自动递增", "负责人", "更新者", "更新时间"})
	// dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{item.CompanyId, item.ContactNumber, item.CreateBy, item.CreatedAt, item.Description, item.Email, item.FleetName, item.Id, item.ManagerName, item.UpdateBy, item.UpdatedAt})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
