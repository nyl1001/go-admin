import request from '@/utils/request'

// 查询Fleet列表
export function listFleet(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/fleet',
    method: 'get',
    params: query
  })
}

// 查询Fleet详细
export function getFleet(id) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/fleet/' + id,
    method: 'get'
  })
}

// 新增Fleet
export function addFleet(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/fleet',
    method: 'post',
    data: data
  })
}

// 修改Fleet
export function updateFleet(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/fleet/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Fleet
export function delFleet(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/fleet',
    method: 'delete',
    data: data
  })
}

// 导出Fleet列表
export function exportFleet(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/fleet/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

