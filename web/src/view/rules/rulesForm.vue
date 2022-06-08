<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="100px">
        <el-form-item label="规则名称:">
          <el-input v-model="formData.ruleName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="rule类型:">
          <el-select v-model="formData.ruleType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in toolsNameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="rule状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createRules,
  updateRules,
  findRules
} from '@/api/rules' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Rules',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findRules({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rerules
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('toolsName')
    await this.getDict('APIStatus')
  },
  methods: {
    async save() {
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
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
