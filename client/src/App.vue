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

      const res = await fetch('http://localhost:5000/orders')
      .catch(function(error){
        console.error(error)
      });
      const data = await res.json()
      return data
    }
  },
  async created() {
    this.orders = await this.fetchOrders()
  }
}
</script>

<style>
#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
  /* -webkit-font-smoothing: antialiased; */
  /* -moz-osx-font-smoothing: grayscale; */
  /* text-align: center; */
  /* color: #2c3e50; */
  /* margin-top: 60px; */
}
</style>
