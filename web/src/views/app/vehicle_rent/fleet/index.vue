<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="公司" prop="companyId">
            <el-select v-model="queryParams.companyId" placeholder="公司" clearable size="small">
              <el-option
                v-for="item in companies"
                :key="item.companyName"
                :label="item.companyName"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="车队名称" prop="fleetName">
            <el-input v-model="queryParams.fleetName" placeholder="请输入车队名称" clearable size="small"
                      @keyup.enter.native="handleQuery"/>
          </el-form-item>
          <el-form-item label="创建时间">
            <el-date-picker
              v-model="dateRange"
              size="small"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              align="right"
              value-format="yyyy-MM-dd HH:mm:ss"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['app:vehicle_rent:fleet:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >
              新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['app:vehicle_rent:fleet:export']"
              type="success"
              icon="el-icon-download"
              size="mini"
              @click="handleExport"
            >
              Excel导出
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" stripe border :data="tableList">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column label="公司" align="center" prop="companyId" :formatter="companyIdFormat" width="100">
            <template slot-scope="scope">
              {{ companyIdFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="车队名称" align="center" prop="fleetName" :show-overflow-tooltip="true" />
          <el-table-column label="负责人" align="center" prop="managerName" :show-overflow-tooltip="true" />
          <el-table-column label="负责人邮箱" align="center" prop="email" :show-overflow-tooltip="true" />
          <el-table-column label="车队描述（可选）" align="center" prop="description" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="更新时间" align="center" prop="updatedAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.updatedAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column width="160" fixed="right" label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['app:vehicle_rent:fleet:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >
                修改
              </el-button>
              <el-button
                v-permisaction="['app:vehicle_rent:fleet:del']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total>0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :close-on-click-modal="false" :title="title" :visible.sync="open" width="500" append-to-body>
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="公司" prop="companyId">
              <treeselect
                v-model="form.companyId"
                :options="companyOptions"
                placeholder="请选择归属父级公司"
              />
            </el-form-item>
            <el-form-item label="车队名称" prop="fleetName">
              <el-input v-model="form.fleetName" placeholder="车队名称"/>
            </el-form-item>
            <el-form-item label="负责人" prop="managerName">
              <el-input v-model="form.managerName" placeholder="负责人"/>
            </el-form-item>
            <el-form-item label="联系电话" prop="contactNumber">
              <el-input v-model="form.contactNumber" placeholder="联系电话"/>
            </el-form-item>
            <el-form-item label="负责人邮箱" prop="email">
              <el-input v-model="form.email" placeholder="负责人邮箱"  />
            </el-form-item>
            <el-form-item label="车队描述（可选）" prop="description">
              <el-input v-model="form.description" placeholder="车队描述（可选）"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { addFleet, delFleet, exportFleet, getFleet, listFleet, updateFleet } from '@/api/app/vehicle_rent/fleet'
import { companyTreeSelect, getAllCompanies } from '@/api/app/vehicle_rent/company'

import { resolveBlob } from '@/utils/download'
import { selectItemsLabel } from '@/utils/costum'
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'

export default {
  name: 'Fleet',
  components: {Treeselect},
  data() {
    return {
      // 遮罩层
      loading: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 日期范围
      dateRange: [],
      // 类型数据字典
      tableList: [],
      // 公司树选项
      companyOptions: undefined,
      companies: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        companyId: undefined,
        fleetName: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        companyId: [{required: true, message: '关联的公司ID (参考companies表的主键)不能为空', trigger: 'blur'}],
        fleetName: [{required: true, message: '车队名称不能为空', trigger: 'blur'}]
      }
    }
  },
  created() {
    this.allCompanies()
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listFleet(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tableList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    allCompanies() {
      getAllCompanies().then(response => {
        this.companies = response.data
        if (this.companies.length > 0) {
          // 设置默认选中的品牌ID为列表中的第一个
          // this.queryParams.companyId = this.companies[0].id
        }
      })
    },
    /** 查询部门下拉树结构 */
    getCompanyTreeSelect() {
      companyTreeSelect().then(response => {
        this.companyOptions = response.data
      })
    },
    companyIdFormat(row) {
      return selectItemsLabel(this.companies, 'companyName', 'id', row.companyId)
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        companyId: undefined,
        description: undefined,
        fleetName: undefined,
        id: undefined,
      }
      this.resetForm('form')
    },
    // 搜索按钮操作
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    // 重置按钮操作
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    // 新增按钮操作
    handleAdd() {
      this.reset()
      this.getCompanyTreeSelect()
      this.open = true
      this.title = '添加车队表'
      this.isEdit = false
    },
    // 修改按钮操作
    handleUpdate(row) {
      this.reset()
      this.getCompanyTreeSelect()
      getFleet(row.id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改车队表'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFleet(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addFleet(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    // 删除按钮操作
    handleDelete(row) {
      const ids = [row.id]
      this.$confirm('是否确认删除编号为"' + ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function () {
        return delFleet({'ids': ids})
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function () {
      })
    },
    /** 下载excel */
    handleExport() {
      this.$confirm('是否确认导出所选数据？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        exportFleet(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, '车队表')
        })
      }).catch(() => {
      })
    }
  }
}
</script>
