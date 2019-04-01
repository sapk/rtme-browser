<template>
  <div class="container">
    <br>
    <h2 class="column col-12">Database configuration</h2>
    <form class="form-horizontal">
      <div class="form-group">
        <div class="column col-2 col-lg-3 col-sm-12">
          <label class="form-label text-left" for="input-type">Type</label>
        </div>
        <div class="column col-10 col-lg-9 col-sm-12">
          <select class="form-select select-lg" id="input-type" v-model="type">
            <option value="mssql" selected>Microsoft SQL Server</option>
          </select>
        </div>
      </div>
      <div class="form-group">
        <div class="column col-2 col-lg-3 col-sm-12">
          <label class="form-label text-left" for="input-host">Host</label>
        </div>
        <div class="column col-10 col-lg-9 col-sm-12">
          <input class="form-input" type="text" id="input-host" placeholder="Host" v-model="host">
        </div>
      </div>
      <div class="form-group">
        <div class="column col-2 col-lg-3 col-sm-12">
          <label class="form-label text-left" for="input-user">User</label>
        </div>
        <div class="column col-10 col-lg-9 col-sm-12">
          <input class="form-input" type="text" id="input-user" placeholder="User" v-model="user">
        </div>
      </div>
      <div class="form-group">
        <div class="column col-2 col-lg-3 col-sm-12">
          <label class="form-label text-left" for="input-pass">Password</label>
        </div>
        <div class="column col-10 col-lg-9 col-sm-12">
          <input class="form-input" type="password" id="input-pass" placeholder="Password" v-model="pass">
        </div>
      </div>
      <div class="form-group">
        <div class="column col-2 col-lg-3 col-sm-12">
          <label class="form-label text-left" for="input-rtme">Base RTME</label>
        </div>
        <div class="column col-10 col-lg-9 col-sm-12">
          <input class="form-input" type="text" id="input-host" placeholder="RTME" v-model="rtme">
        </div>
      </div>
      <div class="form-group">
        <div class="column col-2 col-lg-3 col-sm-12">
          <label class="form-label text-left" for="input-cfg">Base CFG</label>
        </div>
        <div class="column col-10 col-lg-9 col-sm-12">
          <input class="form-input" type="text" id="input-cfg" placeholder="CFG" v-model="cfg">
        </div>
      </div>
      <button class="float-right btn btn-lg m-2" @click.prevent="connect">Connexion</button>
    </form>
  </div>
</template>

<style>
</style>


<script>
import api from "@/lib/api";

export default {
  data: function() {
    return {
        type: localStorage.rtmeType || "mssql", //default
        host: localStorage.rtmeHost,
        user: localStorage.rtmeUser,
        pass: localStorage.rtmePass,
        rtme: localStorage.rtmeRTME,
        cfg: localStorage.rtmeCFG,
    };
  },
  computed: {
    rtmeURL: function () {
      //TODO support other db types
      return "server="+this.host+";user id="+this.user+";password="+this.pass+";database="+this.rtme+""
    },
    cfgURL: function () {
      //TODO support other db types
      return "server="+this.host+";user id="+this.user+";password="+this.pass+";database="+this.cfg+""
    }
  },
  methods: {
      connect: async function () {
          localStorage.setItem('rtmeType', this.type);
          localStorage.setItem('rtmeHost', this.host);
          localStorage.setItem('rtmeUser', this.user);
          localStorage.setItem('rtmeRTME', this.rtme);
          localStorage.setItem('rtmeCFG', this.cfg);
          await api.SetAppConfig(this.type, this.rtmeURL, this.cfgURL)
          window.location.reload(); //dead simple (a little too much)
      }
  }
};
</script>