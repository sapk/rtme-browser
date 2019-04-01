import axios from "axios";

class API {
  //TODO use cache
  async GetGroups() {
    let response = await axios.get("api/v1/cfg/groups");
    return response.data;
  }
  async GetGroupAgents(dbid) {
    let response = await axios.get("api/v1/cfg/groups/" + dbid + "/users");
    return response.data;
  }
  async GetPlace(dbid) {
    let response = await axios.get("api/v1/cfg/places/" + dbid);
    return response.data;
  }
  async GetTimelineStatus(d, group) {
    //console.log(site_id,d)
    let url = "api/v1/rtme/" + d + "/status"
    if (!(typeof group === "undefined" || group === 'Tous')) {
      url += "?group=" + group
    }
    let response = await axios.get(url)
    return response.data
  }

  async GetTimeline(d, group) {
    //console.log(site_id,d)
    let url = "api/v1/rtme/" + d + "/login"
    if (!(typeof group === "undefined" || group === 'Tous')) {
      url += "?group=" + group
    }
    let response = await axios.get(url)
    return response.data
  }
}
let api = new API();

export default api;
export let statusCode = {
  4: "Ready",
  5: "OffHook",
  6: "CallDialing",
  7: "CallRinging",
  8: "NotReadyForNextCall",
  9: "AfterCallWork",
  13: "CallOnHold",
  16: "ASM_Engaged",
  17: "ASM_Outbound",
  18: "CallUnknown",
  19: "CallConsult",
  20: "CallInternal",
  21: "CallOutbound",
  22: "CallInbound",
  23: "LoggedOut"
};