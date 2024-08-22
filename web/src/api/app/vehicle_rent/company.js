import request from '@/utils/request'

// 查询Companies列表
export function listCompanies(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company',
    method: 'get',
    params: query
  })
}

// 查询Companies详细
export function getAllCompanies() {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company/all',
    method: 'get'
  })
}

// 查询公司下拉树结构
export function companyTreeSelect() {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company/tree',
    method: 'get'
  })
}

// 查询Companies详细
export function getCompanies(id) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company/' + id,
    method: 'get'
  })
}

// 新增Companies
export function addCompanies(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company',
    method: 'post',
    data: data
  })
}

// 修改Companies
export function updateCompanies(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Companies
export function delCompanies(data) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company',
    method: 'delete',
    data: data
  })
}

// 导出Companies列表
export function exportCompanies(query) {
  return request({
    url: '/admin-api/v1/app/vehicle_rent/company/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

