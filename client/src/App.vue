<template>
  <div class="container">
    <Header />
    <Orders :orders="orders"/>
    <Footer />
  </div>
</template>

<script>
import Header from'./components/Header'
import Footer from'./components/Footer'
import Orders from'./components/Orders'


export default {
  name: 'App',
  components: {
    Header,
    Footer,
    Orders,
  },
  data() {
    return {
      orders: [],
    }
  },
  methods: { 
    async fetchOrders() {
      const res = await fetch('http://localhost:5000/orders', 
      { mode: 'no-cors'})
      .then(
        function(response) {
          if (response.status !== 200) {
            console.log('Looks like there was a problem. Status Code: ' +
              response.status);
            return;
          }

          // Examine the text in the response
          response.json().then(function(data) {
          console.log(data);
          });
        }
      )
  .catch(function(err) {
    console.log('Fetch Error :-S', err);
  });
      
    }
  },
  async created() {
    this.orders = await this.fetchOrders()
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>
