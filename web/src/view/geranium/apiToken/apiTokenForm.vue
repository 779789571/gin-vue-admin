<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="API类型:">
          <el-select v-model="formData.apiType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in apiNameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="API状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in APIStatus" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="api内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="次数限制:">
          <el-input v-model.number="formData.limitTimes" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="剩余次数:">
          <el-input v-model.number="formData.remaingTimes" clearable placeholder="请输入" />
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
  createApiToken,
  updateApiToken,
  findApiToken
} from '@/api/apiToken' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'ApiToken',
  mixins: [infoList],
  data() {
    return {
      type: '',
      apiNameOptions: [],
      APIstatusOptions: [],
      formData: {
        apiType: undefined,
        status: undefined,
        content: '',
        limitTimes: 0,
        remaingTimes: 0,
        desc: '',
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findApiToken({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reapiToken
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('apiName')
    await this.getDict('APIStatus')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createApiToken(this.formData)
          break
        case 'update':
          res = await updateApiToken(this.formData)
          break
        default:
          res = await createApiToken(this.formData)
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
