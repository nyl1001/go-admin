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

type Model struct {
	service.Service
}

// NewModelService
// @Description: 实例化Model
// @return *Model
func NewModelService(s *service.Service) *Model {
	var srv = new(Model)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

//
// GetPage
// @Description: 获取Model列表

// @return []models.Model
// @return int64
// @return int
// @return error
func (e *Model) GetPage(c *dto.ModelQueryReq, p *middleware.DataPermission) ([]models.Model, int64, int, error) {
	var data models.Model
	var list []models.Model
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
// @Description: 获取Model对象
// @receiver e
// @return *models.Model
// @return int
// @return error
func (e *Model) Get(id int64, p *middleware.DataPermission) (*models.Model, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.Model{}
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
// @Description: 通过自定义条件获取Model一条记录
// @receiver e
// @return *models.Model
// @return error
func (e *Model) QueryOne(queryCondition *dto.ModelQueryReq, p *middleware.DataPermission) (*models.Model, int, error) {
	data := &models.Model{}
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
func (e *Model) Count(queryCondition *dto.ModelQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.Model{}).
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
// @Description: 创建Model对象

// @return int64 插入数据的主键
// @return int
// @return error
func (e *Model) Insert(c *dto.ModelInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	now := time.Now()
	var data models.Model
	data.ModelName = c.ModelName
	data.Description = c.Description
	data.BrandId = c.BrandId
	data.TypeId = c.TypeId
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
// @Description: 修改Model对象

// @return bool 是否有数据更新
// @return error
func (e *Model) Update(c *dto.ModelUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.ModelName != "" && data.ModelName != c.ModelName {
		updates["model_name"] = c.ModelName
	}
	if c.Description != "" && data.Description != c.Description {
		updates["description"] = c.Description
	}
	if c.BrandId >= 0 && data.BrandId != c.BrandId {
		updates["brand_id"] = c.BrandId
	}
	if c.TypeId >= 0 && data.TypeId != c.TypeId {
		updates["type_id"] = c.TypeId
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
// @Description: 删除Model
// @receiver e
// @return int
// @return error
func (e *Model) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.Model
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出Model excel数据
// @receiver e
// @return []byte
// @return int
// @return error
func (e *Model) GetExcel(list []models.Model) ([]byte, error) {
	sheetName := "Model"
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{"主键，自动递增", "车型名称", "车型描述（可选）", "关联的品牌ID (参考vehicle_brands表的主键)", "关联的车辆种类ID (参考vehicle_types表的主键)", "创建者", "更新者", "创建时间", "更新时间"})
	// dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{item.Id, item.ModelName, item.Description, item.BrandId, item.TypeId, item.CreateBy, item.UpdateBy, item.CreatedAt, item.UpdatedAt})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
