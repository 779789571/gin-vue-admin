<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="运行规则:">
          <el-input v-model="formData.useRules" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="域名:">
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发现时间在..之前:">
          <el-date-picker v-model="formData.timeBefore" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="发现时间在..之后:">
          <el-date-picker v-model="formData.timeAfter" type="date" placeholder="选择日期" clearable></el-date-picker>
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
  createrequiredInfo,
  updaterequiredInfo,
  findrequiredInfo
} from '@/api/requiredInfo' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'requiredInfo',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findrequiredInfo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rerequiredInfo
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
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
