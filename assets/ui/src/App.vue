<template>
  <div id="app" v-if="appStatus">
    <viewer v-if="appStatus.status === 'ready'" />
    <login v-else />
  </div>
</template>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /*margin-top: 20px;*/
}
</style>


<script>
require('spectre.css');
import Login from "./components/Login.vue";
import Viewer from "./components/Viewer.vue";

import api from "@/lib/api";

export default {
  name: "app",
  components: {
    Login,
    Viewer
  },
  data: function() {
    return {};
  },
  asyncComputed: {
    appStatus: function() {
      return api.GetAppStatus();
    },
  }
};
</script>