<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="域名:">
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="src归属:">
          <el-input v-model="formData.belongToSrc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否为泛解析:">
          <el-switch v-model="formData.isWildDomain" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="域名发现时间:">
          <el-date-picker v-model="formData.discoveryTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="IP地址:">
          <el-input v-model="formData.ipAddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开放端口:">
          <el-input v-model="formData.openPorts" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否监控该域名:">
          <el-switch v-model="formData.isMonitoring" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="web响应码:">
          <el-input v-model.number="formData.webStatusCode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="web页面标题:">
          <el-input v-model="formData.webTitles" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已扫描:">
          <el-switch v-model="formData.hadScanned" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="中间件信息:">
          <el-input v-model="formData.middleware" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="指纹信息:">
          <el-input v-model="formData.fingerprint" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="页面大小:">
          <el-input v-model="formData.webPageSize" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="一级域名:">
          <el-input v-model="formData.primaryDomain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否在黑名单中:">
          <el-switch v-model="formData.isOnBlacklist" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createCollector,
  updateCollector,
  findCollector
} from '@/api/collector' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Collector',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findCollector({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.recollector
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
