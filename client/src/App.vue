<template>
  <div class="container">
    <Header />
  <!-- <p>Computed reversed message: "{{ reversedMessage }}"</p> -->
    <Body :orders="orders"/>
    <Footer />
  </div>
</template>

<script>
import Header from'./components/Header'
import Footer from'./components/Footer'
import Body from'./components/Body'


export default {
  name: 'App',
  components: {
    Header,
    Footer,
    Body,
  },
  data() {
    return {
      orders: [],
    }
  },
  methods: { 
    async fetchOrders() {
      console.log("Loading dataset")
      const res = await fetch('http://localhost:5000/orders')
        .catch(function(error){
          console.error(error)
      });
      const data = await res.json()
      console.log("Finished loading dataset")
      return data
    }
  },
  async created() {
    this.orders = await this.fetchOrders()
  },
}
</script>

<style>
#app {
font-family: sans-serif;
}
</style>
