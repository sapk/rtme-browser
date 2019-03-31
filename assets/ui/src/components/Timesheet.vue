<template>
  <div class="timesheet" v-if="timeline">
    <aside>
      <ul>
        <li v-for="(sessions, user) in timeline.Sessions" :key="user">{{user}} ({{sessions.length}})</li>
      </ul>
    </aside>
    <section>
      <div>
        <header>
          <ul>
            <li v-for="t in timeArray" :key="t">
              <time :datetime="t">{{t}}</time>
            </li>
          </ul>
        </header>
        <ul class="timeline">
          <li v-for="(sessions, user) in timeline.Sessions" :key="user">
            <a
              v-for="(session, id) in sessions"
              :key="id"
              :title="hour(session.Start)+' -> '+hour(session.End)+' ('+placeDetails(session.Place)+')'"
              href="#"
              class="time-entry"
              v-bind:style="{ width: (session.End-session.Start)*100/86400 + '%', left: (session.Start-timeline.Start)*100/86400 + '%' }"
            >
              <small>{{hour(session.Start)}} -> {{hour(session.End)}} ({{placeDetails(session.Place)}})</small>
            </a>
            <a
              v-for="(status, id) in statusByUser(user)"
              :key="user+'-'+id"
              :title="hour(status.Start)+' -> '+hour(status.End)+' ('+translateStatus(status.Status)+')'"
              href="#"
              class="time-entry status-entry"
              :class="'status-entry-'+status.Status"
              v-bind:style="{ width: (status.End-status.Start)*100/86400 + '%', left: (status.Start-timeline.Start)*100/86400 + '%' }"
            ></a>
          </li>
        </ul>
      </div>
    </section>
  </div>
</template>

<style>
.btn input[type="date"] {
  background: transparent;
  border: 0;
  color: #000;
  margin-right: -18px;
  margin-left: 2px;
}
input[type="date"]::-webkit-search-cancel-button,
input[type="date"]::-webkit-clear-button,
input[type="date"]::-moz-clear-button,
input[type="date"]::-moz-clear,
input[type="date"]::-ms-clear,
input[type="date"]::clear,
input[type="date"]::-webkit-inner-spin-button {
  appearance: none;
  -moz-appearance: none;
  -webkit-appearance: none;
  display: none;
}
.timesheet aside {
  width: 15%;
  padding: 0 !important;
  margin-top: 42px;
  border-right: 5px solid transparent;
  margin-left: 0.5%;
}
.timesheet ul,
.timesheet li {
  margin: 0;
  padding: 0;
  list-style-type: none;
}
.timesheet aside li:not(:last-of-type) {
  border-bottom: 1px solid #fff;
}
.timesheet aside li,
.timesheet header li,
.timesheet ul li {
  height: 42px;
}
.timesheet aside li {
  padding: 0 15px;
  background-color: #efefef;
  line-height: 42px;
}
.timesheet section {
  width: 84%;
  padding: 0 !important;
  overflow-x: auto;
}
.timesheet section div {
  white-space: nowrap;
  display: inline-block;
}
.timesheet aside,
.timesheet section {
  float: left;
}
.timesheet section header li {
  display: inline-block;
  font-size: 1rem;
  width: 0;
  line-height: 42px;
  display: inline-block;
  position: relative;
}
.timesheet section header li time {
  display: block;
  position: absolute;
  left: 0;
}
.timesheet time {
  font-size: 0.8em;
  font-weight: 500;
}
.timesheet ul.timeline {
  margin-top: -6px;
  border-left: none;
  position: relative;
  overflow: hidden;
}
.timesheet ul.timeline li:not(:last-of-type) {
  border-bottom: none;
}
.timesheet ul.timeline li:first-of-type {
  border-top: 1px solid #e5e5e5;
}
.timesheet ul.timeline li:nth-of-type(2n + 1) {
  background-color: #fdfdfd;
}
.timesheet ul.timeline li {
  position: relative;
  background-color: #f4f4f4;
}
.timesheet .time-entry {
  background-color: #39bae7;
  transition: 200ms background-color;
  height: 37px;
  display: block;
  position: absolute;
  z-index: 2;
  margin: 2px 0;
  padding: 0 1px;
  white-space: normal;
  overflow: hidden;
  color: #fff;
  border: 1px solid #3f9ed8;
  border-radius: 6px;
}
.timesheet ul.timeline li::after,
.timesheet ul.timeline li::before {
  content: "";
  position: absolute;
  z-index: 1;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
}
.timesheet ul.timeline li::before {
  background-image: linear-gradient(to right, #e5e5e5 1px, transparent 1px);
  background-size: 21px auto;
}
.timesheet ul.timeline li::after {
  background-image: linear-gradient(
    to right,
    #e5e5e5,
    #e5e5e5 1px,
    #f4f4f4 1px,
    #f4f4f4 2px,
    #e5e5e5 2px,
    #e5e5e5 3px,
    transparent 3px,
    transparent
  );
  background-size: 84px auto;
  background-position: -2px 0;
}
.timesheet a {
  text-decoration: none;
}
.timesheet .time-entry small {
  position: relative;
  top: 26%;
  -webkit-transform: translateY(-50%);
  transform: translateY(-50%);
  display: block;
  padding: 0 5px;
}
.timesheet section header li:not(:last-of-type) {
  width: 84px;
}
.timesheet section header li:not(:first-of-type) time {
  -webkit-transform: translateX(-50%);
  transform: translateX(-50%);
}
.timesheet section header li:last-of-type time {
  -webkit-transform: translateX(-100%);
  transform: translateX(-100%);
}
.timesheet .time-entry.status-entry {
  height: 9px;
  margin-top: 29px;
  background-color: #ccc;
}
.time-entry.status-entry.status-entry-4 {
  background-color: #7eb77e;
}
.time-entry.status-entry.status-entry-8 {
  background-color: #ccc;
}
.time-entry.status-entry.status-entry-9 {
  background-color: #ccc;
}
.time-entry.status-entry.status-entry-7 {
  background-color: #fffc2d;
  margin-top: 22px;
  height: 16px;
}
.time-entry.status-entry.status-entry-22 {
  background-color: green;
}
</style>


<script>
import api from "@/lib/api";

export default {
  props: ["group", "timeline", "status"],
  data: function() {
    return {
      places: {}
    };
  },
  computed: {
    timeArray: function() {
      var times = [];
      for (var i = 0; i <= 24; i++) {
        var formattedNumber = ("0" + i).slice(-2);
        times.push(formattedNumber + ":00");
      }
      return times;
    }
  },
  /*
  asyncComputed: {
    agents: function() {
      if (typeof group != "undefined" && group != "Tous" && group != "") {
        return api.GetGroupAgents(group);
      }
      return this.timeline.Users || [] //Reurn all user of timelin for all
    }
  },
  */
  methods: {
    translateStatus: function(status) {
      switch (
        status //TODO use an array
      ) {
        case "4":
          return "Ready";
        case "5":
          return "OffHook";
        case "6":
          return "CallDialing";
        case "7":
          return "CallRinging";
        case "8":
          return "NotReadyForNextCall";
        case "9":
          return "AfterCallWork";
        case "13":
          return "CallOnHold";
        case "16":
          return "ASM_Engaged";
        case "17":
          return "ASM_Outbound";
        case "18":
          return "CallUnknown";
        case "19":
          return "CallConsult";
        case "20":
          return "CallInternal";
        case "21":
          return "CallOutbound";
        case "22":
          return "CallInbound";
        case "23":
          return "LoggedOut";
        default:
          return status;
      }
    },
    hour: function(timestamp) {
      var date = new Date(timestamp * 1000);
      return date.toLocaleTimeString().substring(0, 5);
    },
    //todo TRANSLATE STATUS
    statusByUser: function(user) {
      var v = this;
      if (v.status === null) {
        return {}
      }
      var dbid = v.timeline.Users[user];
      return v.status.Status.filter(s => s.Agent == dbid);
    },
    placeDetails: function(dbid) {
      var v = this;
      if (typeof v.places[dbid] == "undefined") {
        v.$set(v.places, dbid, dbid); //Default value //TODO sync
        api.GetPlace(dbid).then(function(place) {
          //console.log("DEBUG",dbid, place.name)
          v.$set(v.places, dbid, place.name); //update
        });
      }
      return v.places[dbid];
    }
  }
};
</script>