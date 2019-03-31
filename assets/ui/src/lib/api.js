import axios from "axios";

class API {
    //TODO use cache
    async GetGroups() {
        let response = await axios.get("api/v1/cfg/groups");
        return response.data;
    }
    async GetGroupAgents(dbid) {
        let response = await axios.get("api/v1/cfg/groups/" + dbid+ "/users");
        return response.data;
    }
    async GetPlace(dbid) {
        let response = await axios.get("api/v1/cfg/places/" + dbid);
        return response.data;
    }
    async GetTimelineStatus(d,group) {
        //console.log(site_id,d)
        let url =  "api/v1/rtme/" + d + "/status"
        if (!(typeof group === "undefined" || group === 'Tous')) {
          url += "?group="+group
        }
        let response = await axios.get(url)
        return response.data
      }
    
      async GetTimeline(d,group) {
        //console.log(site_id,d)
        let url =  "api/v1/rtme/" + d + "/login"
        if (!(typeof group === "undefined" || group === 'Tous')) {
          url += "?group="+group
        }
        let response = await axios.get(url)
        return response.data
    }
}
let api = new API();

export default api;