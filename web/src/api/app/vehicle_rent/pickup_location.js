import request from '@/utils/request'

// 查询PickupLocation列表
export function listPickupLocation(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/pickup_location',
    method: 'get',
    params: query
  })
}

// 查询PickupLocation详细
export function getPickupLocation(id) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/pickup_location/' + id,
    method: 'get'
  })
}

// 新增PickupLocation
export function addPickupLocation(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/pickup_location',
    method: 'post',
    data: data
  })
}

// 修改PickupLocation
export function updatePickupLocation(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/pickup_location/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除PickupLocation
export function delPickupLocation(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/pickup_location',
    method: 'delete',
    data: data
  })
}

// 导出PickupLocation列表
export function exportPickupLocation(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/pickup_location/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

