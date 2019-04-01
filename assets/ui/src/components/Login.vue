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
            <select class="form-select select-lg" id="input-type">
                  <option value="mssql" selected>MSSQL</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <div class="column col-2 col-lg-3 col-sm-12">
            <label class="form-label text-left" for="input-host">Host</label>
          </div>
          <div class="column col-10 col-lg-9 col-sm-12">
            <input class="form-input" type="text" id="input-host" placeholder="Host">
          </div>
        </div>
        <div class="form-group">
          <div class="column col-2 col-lg-3 col-sm-12">
            <label class="form-label text-left" for="input-user">User</label>
          </div>
          <div class="column col-10 col-lg-9 col-sm-12">
            <input class="form-input" type="text" id="input-user" placeholder="User">
          </div>
        </div>
        <div class="form-group">
          <div class="column col-2 col-lg-3 col-sm-12">
            <label class="form-label text-left" for="input-pass">Password</label>
          </div>
          <div class="column col-10 col-lg-9 col-sm-12">
            <input class="form-input" type="text" id="input-pass" placeholder="Password">
          </div>
        </div>
        <div class="form-group">
          <div class="column col-2 col-lg-3 col-sm-12"> 
            <label class="form-label text-left" for="input-rtme">Base RTME</label>
          </div>
          <div class="column col-10 col-lg-9 col-sm-12">
            <input class="form-input" type="text" id="input-host" placeholder="RTME">
          </div>
        </div>
        <div class="form-group">
          <div class="column col-2 col-lg-3 col-sm-12">
            <label class="form-label text-left" for="input-cfg">Base CFG</label>
          </div>
          <div class="column col-10 col-lg-9 col-sm-12">
            <input class="form-input" type="text" id="input-cfg" placeholder="CFG">
          </div>
        </div>
      </form>
    </div>
</template>

<style>

</style>


<script>

export default {
  data: function() {
    return {};
  },
  asyncComputed: {
  }
};
</script>