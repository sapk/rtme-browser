<template>
  <div id="app">
    <div v-if="appStatus.status === 'ready'">
      <select v-model="selectedGroup">
        <option>Tous</option>
        <option v-for="g in groups" :key="g">{{g}}</option>
      </select>
      <input type="date" id="day" v-model="selectedDate" required="required">
      <br>
      <br>
      <br>
      <Timesheet :timeline="timeline" :status="status"/>
    </div>
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
input[type="date"] {
  padding: 7px;
  padding-bottom: 7px;
  display: inline-block;
  margin-left: 4px;
  padding-bottom: 5px;
}
</style>


<script>
require('spectre.css');
import Timesheet from "./components/Timesheet.vue";
import Login from "./components/Login.vue";

import api from "@/lib/api";

export default {
  name: "app",
  components: {
    Timesheet,
    Login
  },
  data: function() {
    return {
      //selectedDate: "2018-09-10",
      selectedDate: new Date().toISOString().substring(0, 10),
      selectedGroup: "Tous",
      //groups: ["S1", "S2"],
      places: {}
    };
  },
  asyncComputed: {
    appStatus: function() {
      return api.GetAppStatus();
    },
    groups: function() {
      return api.GetGroups();
    },
    timeline: function() {
      return api.GetTimeline(this.selectedDate, this.selectedGroup);
    },
    status: function() {
      return api.GetTimelineStatus(this.selectedDate, this.selectedGroup);
    }
  }
};
</script>