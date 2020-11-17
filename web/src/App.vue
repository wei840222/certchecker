<template>
  <el-container id="app">
    <el-header style="text-align: start;height: unset;">
      <el-button type="primary" @click="dialogFormVisible = true"
        >新增</el-button
      >
      <el-input
        v-model="filter"
        placeholder="請输入内容以搜尋"
        prefix-icon="el-icon-search"
        style="width: 200px;margin-left: 10px;"
      ></el-input>
      <el-dialog title="新增域名" :visible.sync="dialogFormVisible">
        <el-form :model="form">
          <el-form-item label="名稱" label-width="100px">
            <el-input
              v-model="form.name"
              autocomplete="off"
              placeholder="產品名 功能 EX:NN api"
            ></el-input>
          </el-form-item>
          <el-form-item label="域名" label-width="100px">
            <el-input
              v-model="form.host"
              autocomplete="off"
              placeholder="https://abc.com:443"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取消</el-button>
          <el-button type="primary" @click="handleCreate">確定</el-button>
        </div>
      </el-dialog>
    </el-header>

    <el-main>
      <el-table :data="filteredDomain" stripe style="width: 100%">
        <el-table-column
          prop=""
          type="index"
          :index="indexMethod"
          sortable
          width="50"
        />
        <el-table-column prop="host" sortable label="域名" width="250" />
        <el-table-column prop="name" sortable label="名稱" width="300" />
        <el-table-column
          label="開始"
          width="300"
          sortable
          :sort-method="
            (a, b) =>
              Date.parse(a.since) > Date.parse(b.since)
                ? 1
                : Date.parse(a.since) < Date.parse(b.since)
                ? -1
                : 0
          "
        >
          <template slot-scope="scope">
            <i v-if="scope.row.since" class="el-icon-time"></i>
            <span v-if="scope.row.since" style="margin-left: 5px">{{
              rfc3339ToLocaleString(scope.row.since)
            }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="結束"
          sortable
          width="300"
          :sort-method="
            (a, b) =>
              Date.parse(a.end) > Date.parse(b.end)
                ? 1
                : Date.parse(a.end) < Date.parse(b.end)
                ? -1
                : 0
          "
        >
          <template slot-scope="scope">
            <i v-if="scope.row.end" class="el-icon-time"></i>
            <span v-if="scope.row.end" style="margin-left: 5px">{{
              rfc3339ToLocaleString(scope.row.end)
            }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="error" label="錯誤" />
        <el-table-column label="操作" width="150">
          <template slot-scope="scope">
            <el-popconfirm
              title="确定删除嗎？"
              @confirm="handleDelete(scope.row.id)"
            >
              <el-button slot="reference" size="mini" type="danger"
                >刪除</el-button
              >
            </el-popconfirm>
          </template>
        </el-table-column>
        <el-table-column label="Ignore alert" width="150">
          <template slot-scope="scope">
            <el-switch
              @change="handleSwitch(scope.row)"
              v-model="scope.row.isDisableSentNotify"
              active-text="True"
              inactive-text="False"
            >
            </el-switch>
          </template>
        </el-table-column>
        <el-table-column prop="id" label="ID" width="50" />
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: "app",
  data() {
    return {
      dialogFormVisible: false,
      form: {
        name: "",
        host: ""
      },
      filter: "",
      domain: []
    };
  },
  computed: {
    filteredDomain() {
      return this.filter === ""
        ? this.domain
        : this.domain.filter(
            d =>
              d.name.includes(this.filter) ||
              d.host.includes(this.filter) ||
              (d.since && d.since.includes(this.filter)) ||
              (d.end && d.end.includes(this.filter))
          );
    }
  },
  mounted() {
    this.fetchData();
    setInterval(() => this.fetchData(), 3000);
  },
  methods: {
    indexMethod(index) {
      index++;
      return index;
    },
    async handleSwitch(e) {
      await this.$axios.patch(`/domain/${e.id}`, {
        isDisableSentNotify: e.isDisableSentNotify
      });
    },

    rfc3339ToLocaleString: timeStr =>
      new Date(Date.parse(timeStr)).toLocaleString("zh-TW", {
        timeZone: "Asia/Taipei"
      }),
    async fetchData() {
      const res = await this.$axios.get("/domain");
      this.domain = res.data;
    },
    async handleCreate() {
      this.dialogFormVisible = false;
      await this.$axios.post("/domain", {
        name: this.form.name,
        host: this.form.host
      });
      this.form = {
        name: "",
        host: ""
      };
      this.fetchData();
    },
    async handleDelete(id) {
      await this.$axios.delete(`/domain/${id}`);
      this.fetchData();
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}

body {
  margin: 0px;
  padding: 0px;
  background: #fff
    url(https://c.files.bbci.co.uk/6577/production/_110957952_42f5b28f-0145-42c8-b5b9-7333611a3a02.jpg)
    center center fixed no-repeat;
  background-size: cover;
}
</style>
