package apis

import (
	adminService "go-admin/app/admin/service"
	"go-admin/app/app/vehicle_rent/service"
	"go-admin/app/app/vehicle_rent/service/dto"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/dateutils"
	"time"

	"github.com/gin-gonic/gin"
)

type Model struct {
	api.Api
}

//
// GetPage
// @Description: 获取车辆型号列表

func (e Model) GetPage(c *gin.Context) {
	req := dto.ModelQueryReq{}
	s := service.Model{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(lang.SuccessCode, e.Lang))
}

//
// Get
// @Description: 获取车辆型号

func (e Model) Get(c *gin.Context) {
	req := dto.ModelGetReq{}
	s := service.Model{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

//
// Insert
// @Description: 创建车辆型号

func (e Model) Insert(c *gin.Context) {
	req := dto.ModelInsertReq{}
	s := service.Model{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

//
// Update
// @Description: 修改车辆型号

func (e Model) Update(c *gin.Context) {
	req := dto.ModelUpdateReq{}
	s := service.Model{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.Update(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

//
// Delete
// @Description:车辆型号

func (e Model) Delete(c *gin.Context) {
	s := service.Model{}
	req := dto.ModelDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Remove(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

//
// Export
// @Description: 导出车辆型号

func (e Model) Export(c *gin.Context) {
	req := dto.ModelQueryReq{}
	s := service.Model{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	sysConfService := adminService.NewSysConfigService(&s.Service)
	//最小导出数据量
	maxSize, respCode, err := sysConfService.GetWithKeyInt("sys_max_export_size")
	if err != nil {
		e.Error(respCode, err.Error())
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = maxSize
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.GetExcel(list)
	fileName := "model_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
