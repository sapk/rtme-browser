<template>
  <div class="container">
    <div class="columns">
      <div class="column col-2 hide-sm"></div>
      <div class="column col-4 col-sm-11 m-2">
        <select v-model="selectedGroup" class="form-select select-lg">
          <option>Tous</option>
          <option v-for="g in groups" :key="g">{{g}}</option>
        </select>
      </div>
      <div class="column col-4 col-sm-11 m-2">
        <input type="date" id="day" v-model="selectedDate" required="required" class="form-input">
      </div>
      <div class="column col-2 hide-sm"></div>
    </div>
    <Timesheet :timeline="timeline" :status="status"/>
  </div>
</template>

<style>
input[type="date"]{
  height: 40px;
}
</style>


<script>
import Timesheet from "./Timesheet.vue";

import api from "@/lib/api";

export default {
  components: {
    Timesheet
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