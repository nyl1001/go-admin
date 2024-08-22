package service

import (
	"fmt"
	adminService "go-admin/app/admin/service"
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

type VehicleType struct {
	service.Service
}

// NewVehicleTypeService
// @Description: 实例化VehicleType
// @return *VehicleType
func NewVehicleTypeService(s *service.Service) *VehicleType {
	var srv = new(VehicleType)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

//
// GetPage
// @Description: 获取VehicleType列表

// @return []models.VehicleType
// @return int64
// @return int
// @return error
func (e *VehicleType) GetPage(c *dto.VehicleTypeQueryReq, p *middleware.DataPermission) ([]models.VehicleType, int64, int, error) {
	var data models.VehicleType
	var list []models.VehicleType
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

// Get
// @Description: 获取VehicleType对象
// @receiver e
// @return *models.VehicleType
// @return int
// @return error
func (e *VehicleType) Get(id int64, p *middleware.DataPermission) (*models.VehicleType, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.VehicleType{}
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
// @Description: 通过自定义条件获取VehicleType一条记录
// @receiver e
// @return *models.VehicleType
// @return error
func (e *VehicleType) QueryOne(queryCondition *dto.VehicleTypeQueryReq, p *middleware.DataPermission) (*models.VehicleType, int, error) {
	data := &models.VehicleType{}
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
func (e *VehicleType) Count(queryCondition *dto.VehicleTypeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.VehicleType{}).
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
// @Description: 创建VehicleType对象

// @return int64 插入数据的主键
// @return int
// @return error
func (e *VehicleType) Insert(c *dto.VehicleTypeInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	now := time.Now()
	var data models.VehicleType
	data.CreateBy = c.CurrUserId
	data.CreatedAt = &now
	data.TypeName = c.TypeName
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
// @Description: 修改VehicleType对象

// @return bool 是否有数据更新
// @return error
func (e *VehicleType) Update(c *dto.VehicleTypeUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.TypeName != "" && data.TypeName != c.TypeName {
		updates["type_name"] = c.TypeName
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
// @Description: 删除VehicleType
// @receiver e
// @return int
// @return error
func (e *VehicleType) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.VehicleType
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出VehicleType excel数据
// @receiver e
// @return []byte
// @return int
// @return error
func (e *VehicleType) GetExcel(list []models.VehicleType) ([]byte, error) {
	sheetName := "VehicleType"
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("type_name", item.TypeName)

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
