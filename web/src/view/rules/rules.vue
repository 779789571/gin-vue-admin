<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="规则名称">
          <el-input v-model="searchInfo.ruleName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="rule类型">
          <el-input v-model="searchInfo.ruleType" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="rule状态">
          <el-input v-model="searchInfo.status" placeholder="搜索条件" />
        </el-form-item>
            <el-form-item label="是否为动态规则" prop="isDynamic">
            <el-select v-model="searchInfo.isDynamic" clearable placeholder="请选择">
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
        <el-form-item label="rule内容">
          <el-input v-model="searchInfo.content" placeholder="搜索条件" />
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
        <el-table-column align="left" label="规则名称" prop="ruleName" width="120" />
        <el-table-column align="left" label="rule类型" prop="ruleType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.ruleType,"toolsName") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="rule状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,"APIStatus") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="动态规则" prop="isDynamic" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isDynamic) }}</template>
        </el-table-column>
        <el-table-column align="left" label="rule内容" prop="content" width="400" />
        <el-table-column align="left" label="备注" prop="desc" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateRules(scope.row)">变更</el-button>
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
        <el-form-item label="规则名称:">
          <el-input v-model="formData.ruleName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="rule类型:">
          <el-select v-model="formData.ruleType" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in toolsNameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="rule状态:">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in APIStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="动态规则:">
          <el-switch v-model="formData.isDynamic" active-color="#13ce66" inactive-color="#ff4949" active-text="动态" inactive-text="静态" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="rule内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.desc" clearable placeholder="请输入" />
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
  createRules,
  deleteRules,
  deleteRulesByIds,
  updateRules,
  findRules,
  getRulesList
} from '@/api/rules' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Rules',
  mixins: [infoList],
  data() {
    return {
      listApi: getRulesList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      toolsNameOptions: [],
      APIStatusOptions: [],
      formData: {
        ruleName: '',
        ruleType: undefined,
        status: undefined,
        isDynamic: false,
        content: '',
        desc: '',
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('toolsName')
    await this.getDict('APIStatus')
  },
  methods: {
  onReset() {
    this.searchInfo = {}
  },
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.isDynamic === ""){
        this.searchInfo.isDynamic=null
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
        this.deleteRules(row)
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
      const res = await deleteRulesByIds({ ids })
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
    async updateRules(row) {
      const res = await findRules({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rerules
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        ruleName: '',
        ruleType: undefined,
        status: undefined,
        isDynamic: false,
        content: '',
        desc: '',
      }
    },
    async deleteRules(row) {
      const res = await deleteRules({ ID: row.ID })
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
          res = await createRules(this.formData)
          break
        case 'update':
          res = await updateRules(this.formData)
          break
        default:
          res = await createRules(this.formData)
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
