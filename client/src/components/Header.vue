<template>
  <div class="pad-15-hor pad-15-ver header">
    <div>
      <!-- <img src="@/assets/logo.png" width="25px"> -->
       Responsive Search
    </div>
    <!-- <div class="pad-15-hor pad-15-ver search-parent">
        <div class="search-box">
            <input type="text" 
             placeholder="Search by product or ordername">
        </div>
    </div> -->
  </div>
    <!-- <header>
        <h1>Search</h1>
        <input type="text" id="searchBox" name="searchBox">
    </header> -->
</template>

<script>
    export default {
        name:"Header",
        methods:{
            // datasearch(){
            //     console.log("12412412124124")
            // console.log(this.search.text);
            //     var inside = this;
            //     this.wonders_data = this.wonders_data_actual.filter(function(wonder) {
            //     if (
            //         wonder.place
            //         .toLowerCase()
            //         .indexOf(inside.search.text.toLowerCase()) != "-1"
            //     ) {
            //         return wonder;
            //     }
            //     })
            // }
        },
        // data() {

        // },
        // return {

        // }
    }
</script>

<style scoped>
.header {
  display: flex;
  flex-flow: row wrap;
  /* justify-content: space-between; */
}
</style>