<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="域名">
          <el-input v-model="searchInfo.domain" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="src归属">
          <el-input v-model="searchInfo.belongToSrc" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="是否为泛解析" prop="isWildDomain">
          <el-select v-model="searchInfo.isWildDomain" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true"
            >
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="域名发现时间">
          <el-input v-model="searchInfo.discoveryTime" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="IP地址">
          <el-input v-model="searchInfo.ipAddress" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="开放端口">
          <el-input v-model="searchInfo.openPorts" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="是否监控该域名" prop="isMonitoring">
          <el-select v-model="searchInfo.isMonitoring" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true"
            >
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="web响应码">
          <el-input v-model="searchInfo.webStatusCode" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="web页面标题">
          <el-input v-model="searchInfo.webTitles" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="是否已扫描" prop="hadScanned">
          <el-select v-model="searchInfo.hadScanned" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true"
            >
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="中间件信息">
          <el-input v-model="searchInfo.middleware" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="指纹信息">
          <el-input v-model="searchInfo.fingerprint" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="页面大小">
          <el-input v-model="searchInfo.webPageSize" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="一级域名">
          <el-input v-model="searchInfo.primaryDomain" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="是否在黑名单中" prop="isOnBlacklist">
          <el-select v-model="searchInfo.isOnBlacklist" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true"
            >
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="el-icon-delete" size="mini" style="margin-left: 10px;"
                       :disabled="!multipleSelection.length"
            >删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55"/>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="域名" prop="domain" width="120"/>
        <el-table-column align="left" label="src归属" prop="belongToSrc" width="120"/>
        <el-table-column align="left" label="是否为泛解析" prop="isWildDomain" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isWildDomain) }}</template>
        </el-table-column>
        <el-table-column align="left" label="域名发现时间" prop="discoveryTime" width="120"/>
        <el-table-column align="left" label="IP地址" prop="ipAddress" width="120"/>
        <el-table-column align="left" label="开放端口" prop="openPorts" width="120"/>
        <el-table-column align="left" label="是否监控该域名" prop="isMonitoring" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isMonitoring) }}</template>
        </el-table-column>
        <el-table-column align="left" label="web响应码" prop="webStatusCode" width="120"/>
        <el-table-column align="left" label="web页面标题" prop="webTitles" width="120"/>
        <el-table-column align="left" label="是否已扫描" prop="hadScanned" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.hadScanned) }}</template>
        </el-table-column>
        <el-table-column align="left" label="中间件信息" prop="middleware" width="120"/>
        <el-table-column align="left" label="指纹信息" prop="fingerprint" width="120"/>
        <el-table-column align="left" label="页面大小" prop="webPageSize" width="120"/>
        <el-table-column align="left" label="一级域名" prop="primaryDomain" width="120"/>
        <el-table-column align="left" label="是否在黑名单中" prop="isOnBlacklist" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isOnBlacklist) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button"
                       @click="updateCollector(scope.row)"
            >变更
            </el-button>
            <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="域名:">
          <el-input v-model="formData.domain" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="src归属:">
          <el-input v-model="formData.belongToSrc" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="是否为泛解析:">
          <el-switch v-model="formData.isWildDomain" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable
          ></el-switch>
        </el-form-item>
        <el-form-item label="域名发现时间:">
          <el-date-picker v-model="formData.discoveryTime" type="date" style="width:100%" placeholder="选择日期" clearable/>
        </el-form-item>
        <el-form-item label="IP地址:">
          <el-input v-model="formData.ipAddress" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="开放端口:">
          <el-input v-model="formData.openPorts" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="是否监控该域名:">
          <el-switch v-model="formData.isMonitoring" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable
          ></el-switch>
        </el-form-item>
        <el-form-item label="web响应码:">
          <el-input v-model.number="formData.webStatusCode" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="web页面标题:">
          <el-input v-model="formData.webTitles" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="是否已扫描:">
          <el-switch v-model="formData.hadScanned" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable
          ></el-switch>
        </el-form-item>
        <el-form-item label="中间件信息:">
          <el-input v-model="formData.middleware" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="指纹信息:">
          <el-input v-model="formData.fingerprint" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="页面大小:">
          <el-input v-model="formData.webPageSize" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="一级域名:">
          <el-input v-model="formData.primaryDomain" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="是否在黑名单中:">
          <el-switch v-model="formData.isOnBlacklist" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable
          ></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createCollector,
  deleteCollector,
  deleteCollectorByIds,
  findCollector,
  getCollectorList,
  updateCollector,
} from '@/api/collector' //  此处请自行替换地址
import infoList from '@/mixins/infoList'

export default {
  name: 'Collector',
  mixins: [infoList],
  data() {
    return {
      listApi: getCollectorList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        domain: '',
        belongToSrc: '',
        isWildDomain: false,
        discoveryTime: new Date(),
        ipAddress: '',
        openPorts: '',
        isMonitoring: false,
        webStatusCode: 0,
        webTitles: '',
        hadScanned: false,
        middleware: '',
        fingerprint: '',
        webPageSize: '',
        primaryDomain: '',
        isOnBlacklist: false,
      },
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.isWildDomain === '') {
        this.searchInfo.isWildDomain = null
      }
      if (this.searchInfo.isMonitoring === '') {
        this.searchInfo.isMonitoring = null
      }
      if (this.searchInfo.hadScanned === '') {
        this.searchInfo.hadScanned = null
      }
      if (this.searchInfo.isOnBlacklist === '') {
        this.searchInfo.isOnBlacklist = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        this.deleteCollector(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据',
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteCollectorByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateCollector(row) {
      const res = await findCollector({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recollector
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        domain: '',
        belongToSrc: '',
        isWildDomain: false,
        discoveryTime: new Date(),
        ipAddress: '',
        openPorts: '',
        isMonitoring: false,
        webStatusCode: 0,
        webTitles: '',
        hadScanned: false,
        middleware: '',
        fingerprint: '',
        webPageSize: '',
        primaryDomain: '',
        isOnBlacklist: false,
      }
    },
    async deleteCollector(row) {
      const res = await deleteCollector({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createCollector(this.formData)
          break
        case 'update':
          res = await updateCollector(this.formData)
          break
        default:
          res = await createCollector(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
  },
}
</script>

<style>
</style>
