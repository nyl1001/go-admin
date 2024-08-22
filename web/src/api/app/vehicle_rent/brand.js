import request from '@/utils/request'

// 查询Brand列表
export function listBrand(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/brand',
    method: 'get',
    params: query
  })
}

// 查询Brand详细
export function getBrand(id) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/brand/' + id,
    method: 'get'
  })
}

// 新增Brand
export function addBrand(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/brand',
    method: 'post',
    data: data
  })
}

// 修改Brand
export function updateBrand(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/brand/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Brand
export function delBrand(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/brand',
    method: 'delete',
    data: data
  })
}

// 导出Brand列表
export function exportBrand(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/brand/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

