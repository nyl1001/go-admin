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

type PickupLocation struct {
	api.Api
}

//
// GetPage
// @Description: 获取提车网点列表

//
func (e PickupLocation) GetPage(c *gin.Context) {
	req := dto.PickupLocationQueryReq{}
	s := service.PickupLocation{}
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
// @Description: 获取提车网点

//
func (e PickupLocation) Get(c *gin.Context) {
	req := dto.PickupLocationGetReq{}
	s := service.PickupLocation{}
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
// @Description: 创建提车网点

//
func (e PickupLocation) Insert(c *gin.Context) {
	req := dto.PickupLocationInsertReq{}
	s := service.PickupLocation{}
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
// @Description: 修改提车网点

//
func (e PickupLocation) Update(c *gin.Context) {
	req := dto.PickupLocationUpdateReq{}
	s := service.PickupLocation{}
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
// @Description:提车网点

//
func (e PickupLocation) Delete(c *gin.Context) {
	s := service.PickupLocation{}
	req := dto.PickupLocationDeleteReq{}
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
// @Description: 导出提车网点

//
func (e PickupLocation) Export(c *gin.Context) {
	req := dto.PickupLocationQueryReq{}
	s := service.PickupLocation{}
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
	fileName := "pickup_location_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
