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

type Companies struct {
	service.Service
}

// NewCompaniesService
// @Description: 实例化Companies
// @return *Companies
func NewCompaniesService(s *service.Service) *Companies {
	var srv = new(Companies)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

//
// GetPage
// @Description: 获取Companies列表

// @return []models.Companies
// @return int64
// @return int
// @return error
func (e *Companies) GetPage(c *dto.CompaniesQueryReq, p *middleware.DataPermission) ([]models.Companies, int64, int, error) {
	var data models.Companies
	var list []models.Companies
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

// @return []models.Companies
// @return int64
// @return int
// @return error
func (e *Companies) GetAll(p *middleware.DataPermission) ([]models.Companies, int, error) {
	var data models.Companies
	var list []models.Companies
	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			middleware.Permission(data.TableName(), p),
		).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}

// getList 获取公司数据
func (e *Companies) getList(c *dto.CompaniesQueryReq) ([]models.Companies, int, error) {
	var list []models.Companies
	err := e.Orm.Model(&models.Companies{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}

// Get
// @Description: 获取Companies对象
// @receiver e
// @return *models.Companies
// @return int
// @return error
func (e *Companies) Get(id int64, p *middleware.DataPermission) (*models.Companies, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.Companies{}
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
// @Description: 通过自定义条件获取Companies一条记录
// @receiver e
// @return *models.Companies
// @return error
func (e *Companies) QueryOne(queryCondition *dto.CompaniesQueryReq, p *middleware.DataPermission) (*models.Companies, int, error) {
	data := &models.Companies{}
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
func (e *Companies) Count(queryCondition *dto.CompaniesQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.Companies{}).
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
// @Description: 创建Companies对象

// @return int64 插入数据的主键
// @return int
// @return error
func (e *Companies) Insert(c *dto.CompaniesInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	now := time.Now()
	var data models.Companies
	data.Address = c.Address
	data.CompanyName = c.CompanyName
	data.ContactNumber = c.ContactNumber
	data.CreateBy = c.CurrUserId
	data.CreatedAt = &now
	data.Description = c.Description
	data.Email = c.Email
	if c.ParentCompanyId > 0 {
		data.ParentCompanyId = &c.ParentCompanyId
	}
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
// @Description: 修改Companies对象

// @return bool 是否有数据更新
// @return error
func (e *Companies) Update(c *dto.CompaniesUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.Address != "" && data.Address != c.Address {
		updates["address"] = c.Address
	}
	if c.CompanyName != "" && data.CompanyName != c.CompanyName {
		updates["company_name"] = c.CompanyName
	}
	if c.ContactNumber != "" && data.ContactNumber != c.ContactNumber {
		updates["contact_number"] = c.ContactNumber
	}
	if c.Description != "" && data.Description != c.Description {
		updates["description"] = c.Description
	}
	if c.Email != "" && data.Email != c.Email {
		updates["email"] = c.Email
	}
	if c.ParentCompanyId != 0 && (data.ParentCompanyId == nil || *data.ParentCompanyId != c.ParentCompanyId) {
		updates["parent_company_id"] = c.ParentCompanyId
	}
	if c.ParentCompanyId == 0 {
		updates["parent_company_id"] = nil
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
// @Description: 删除Companies
// @receiver e
// @return int
// @return error
func (e *Companies) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.Companies
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出Company excel数据
// @receiver e
// @return []byte
// @return int
// @return error
func (e *Companies) GetExcel(list []models.Companies) ([]byte, error) {
	sheetName := "Company"
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{"公司地址", "公司名称", "联系电话", "创建者", "创建时间", "公司描述（可选）", "公司邮箱", "主键，自动递增", "父公司ID (NULL表示顶级公司)", "更新者", "更新时间"})
	// dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{item.Address, item.CompanyName, item.ContactNumber, item.CreateBy, item.CreatedAt, item.Description, item.Email, item.Id, item.ParentCompanyId, item.UpdateBy, item.UpdatedAt})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// SetCompanyTree 设置公司数据
func (e *Companies) SetCompanyTree(c *dto.CompaniesQueryReq) ([]dto.CompaniesLabel, int, error) {
	list, respCode, err := e.getList(c)
	if err != nil {
		return nil, respCode, err
	}

	m := make([]dto.CompaniesLabel, 0)
	for i := 0; i < len(list); i++ {
		// 只处理顶级公司（ParentCompanyId 为 nil 或 0）
		if list[i].ParentCompanyId != nil && *list[i].ParentCompanyId != 0 {
			continue
		}

		// 构建顶级公司节点
		labelObj := dto.CompaniesLabel{
			Id:       list[i].Id,
			Label:    list[i].CompanyName,
			Children: []dto.CompaniesLabel{},
		}

		// 递归构建子公司结构
		labelInfo := companyTreeCall(&list, labelObj)
		m = append(m, labelInfo)
	}

	return m, lang.SuccessCode, nil
}

// companyTreeCall 递归构造公司数据
func companyTreeCall(companyList *[]models.Companies, labelObj dto.CompaniesLabel) dto.CompaniesLabel {
	list := *companyList
	minList := make([]dto.CompaniesLabel, 0)

	for j := 0; j < len(list); j++ {
		// 跳过没有父公司或父公司ID不匹配的公司
		if list[j].ParentCompanyId == nil || labelObj.Id != *list[j].ParentCompanyId {
			continue
		}

		// 构建子公司节点
		mi := dto.CompaniesLabel{
			Id:       list[j].Id,
			Label:    list[j].CompanyName,
			Children: []dto.CompaniesLabel{},
		}

		// 递归构建子公司树
		ms := companyTreeCall(companyList, mi)
		minList = append(minList, ms)
	}

	// 将所有子公司附加到当前公司节点
	labelObj.Children = minList
	return labelObj
}
