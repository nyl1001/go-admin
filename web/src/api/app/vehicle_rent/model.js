import request from '@/utils/request'

// 查询Model列表
export function listModel(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/model',
    method: 'get',
    params: query
  })
}

// 查询Model详细
export function getModel(id) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/model/' + id,
    method: 'get'
  })
}

// 新增Model
export function addModel(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/model',
    method: 'post',
    data: data
  })
}

// 修改Model
export function updateModel(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/model/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Model
export function delModel(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/model',
    method: 'delete',
    data: data
  })
}

// 导出Model列表
export function exportModel(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/model/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

