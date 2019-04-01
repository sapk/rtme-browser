<template>
  <div>
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
</template>

<style>
input[type="date"] {
  padding: 7px;
  padding-bottom: 7px;
  display: inline-block;
  margin-left: 4px;
  padding-bottom: 5px;
}
</style>


<script>
import Timesheet from "./Timesheet.vue";

import api from "@/lib/api";

export default {
  components: {
    Timesheet,
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