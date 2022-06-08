<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="运行规则">
          <el-input v-model="searchInfo.useRules" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="域名">
          <el-input v-model="searchInfo.domain" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="发现时间在..之前">
          <el-input v-model="searchInfo.timeBefore" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="发现时间在..之后">
          <el-input v-model="searchInfo.timeAfter" placeholder="搜索条件" />
        </el-form-item>
            <el-form-item label="自动生成时间段(当天)" prop="timeAutoCreate">
            <el-select v-model="searchInfo.timeAutoCreate" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
        <el-form-item label="网站标题">
          <el-input v-model="searchInfo.title" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="https证书">
          <el-input v-model="searchInfo.cert" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="Ip或ip段">
          <el-input v-model="searchInfo.ip" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="协议">
          <el-input v-model="searchInfo.protocol" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="icp备案号">
          <el-input v-model="searchInfo.icp" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="服务器状态码">
          <el-input v-model="searchInfo.webStatusCode" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="HTTP请求头">
          <el-input v-model="searchInfo.header" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="国家">
          <el-input v-model="searchInfo.country" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="省份">
          <el-input v-model="searchInfo.province" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="城市">
          <el-input v-model="searchInfo.city" placeholder="搜索条件" />
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
                <el-button icon="el-icon-delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
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
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="运行规则" prop="useRules" width="120" />
        <el-table-column align="left" label="域名" prop="domain" width="120" />
        <el-table-column align="left" label="发现时间在..之前"  width="120" >
          <template #default="scope">{{ formatDateWithoutTime(scope.row.timeBefore) }}</template>
        </el-table-column>
        <el-table-column align="left" label="发现时间在..之后"  width="120" >
          <template #default="scope">{{ formatDateWithoutTime(scope.row.timeAfter) }}</template>
        </el-table-column>
        <el-table-column align="left" label="自动生成时间段(当天)" prop="timeAutoCreate" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.timeAutoCreate) }}</template>
        </el-table-column>
        <el-table-column align="left" label="网站标题" prop="title" width="120" />
        <el-table-column align="left" label="https证书" prop="cert" width="120" />
        <el-table-column align="left" label="Ip或ip段" prop="ip" width="120" />
        <el-table-column align="left" label="协议" prop="protocol" width="120" />
        <el-table-column align="left" label="icp备案号" prop="icp" width="120" />
        <el-table-column align="left" label="服务器状态码" prop="webStatusCode" width="120" />
        <el-table-column align="left" label="HTTP请求头" prop="header" width="120" />
        <el-table-column align="left" label="国家" prop="country" width="120" />
        <el-table-column align="left" label="省份" prop="province" width="120" />
        <el-table-column align="left" label="城市" prop="city" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updaterequiredInfo(scope.row)">变更</el-button>
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
        <el-form-item label="运行规则:">
          <el-input v-model="formData.useRules" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="域名:">
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发现时间在..之前:">
          <el-date-picker v-model="formData.timeBefore" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="发现时间在..之后:">
          <el-date-picker v-model="formData.timeAfter" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="自动生成时间段(当天):">
          <el-switch v-model="formData.timeAutoCreate" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="网站标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="https证书:">
          <el-input v-model="formData.cert" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Ip或ip段:">
          <el-input v-model="formData.ip" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="协议:">
          <el-input v-model="formData.protocol" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="icp备案号:">
          <el-input v-model="formData.icp" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务器状态码:">
          <el-input v-model="formData.webStatusCode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="HTTP请求头:">
          <el-input v-model="formData.header" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="国家:">
          <el-input v-model="formData.country" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="省份:">
          <el-input v-model="formData.province" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市:">
          <el-input v-model="formData.city" clearable placeholder="请输入" />
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
  createrequiredInfo,
  deleterequiredInfo,
  deleterequiredInfoByIds,
  updaterequiredInfo,
  findrequiredInfo,
  getrequiredInfoList
} from '@/api/requiredInfo' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'requiredInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getrequiredInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        useRules: '',
        domain: '',
        timeBefore: new Date(),
        timeAfter: new Date(),
        timeAutoCreate: false,
        title: '',
        cert: '',
        ip: '',
        protocol: '',
        icp: '',
        webStatusCode: '',
        header: '',
        country: '',
        province: '',
        city: '',
      }
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
      if (this.searchInfo.timeAutoCreate === ""){
        this.searchInfo.timeAutoCreate=null
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
        type: 'warning'
      }).then(() => {
        this.deleterequiredInfo(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleterequiredInfoByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updaterequiredInfo(row) {
      const res = await findrequiredInfo({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rerequiredInfo
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        useRules: '',
        domain: '',
        timeBefore: new Date(),
        timeAfter: new Date(),
        timeAutoCreate: false,
        title: '',
        cert: '',
        ip: '',
        protocol: '',
        icp: '',
        webStatusCode: '',
        header: '',
        country: '',
        province: '',
        city: '',
      }
    },
    async deleterequiredInfo(row) {
      const res = await deleterequiredInfo({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
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
          res = await createrequiredInfo(this.formData)
          break
        case 'update':
          res = await updaterequiredInfo(this.formData)
          break
        default:
          res = await createrequiredInfo(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
