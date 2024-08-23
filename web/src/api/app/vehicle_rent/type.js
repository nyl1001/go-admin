import request from '@/utils/request'

// 查询VehicleType列表
export function listVehicleType(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type',
    method: 'get',
    params: query
  })
}

// 查询所有
export function getAllVehicleTypes() {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type/all',
    method: 'get'
  })
}

// 查询VehicleType详细
export function getVehicleType(id) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type/' + id,
    method: 'get'
  })
}

// 新增VehicleType
export function addVehicleType(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type',
    method: 'post',
    data: data
  })
}

// 修改VehicleType
export function updateVehicleType(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除VehicleType
export function delVehicleType(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type',
    method: 'delete',
    data: data
  })
}

// 导出VehicleType列表
export function exportVehicleType(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/type/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

