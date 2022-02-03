<template>
    <div>
        <div class="search-wrapper panel-heading col-sm-12">
            <span>search
                <input type="text" v-model="searchText" placeholder="Search" /> <br> <br>
            </span>
        </div>  
        <h4>Total amount: {{calTotal}}</h4>
        <table>
            <tr>
                <th>Order name</th>
                <th>Customer Company</th>
                <th>Customer name</th>
                <th>order date</th>
                <th>Delivered Amount</th>
                <th>Total Amount</th>
            </tr>
            <tr :key="order.ORDER_NAME" v-for="order in filteredProducts">
                <Order :order="order" />
            </tr>
        </table>
            
        
    </div>
</template>

<script>
    import Order from './Order'
    
    export default {
        name:"Body",
        props: {
            orders: Object,
        },
        components: {
            Order
        },
        data() {
            return {
                searchText: ""
            }
        },
        computed: {
            filteredProducts() {
                try{
                    const filteredOrders = this.filterOrdersByName((this.orders.orders))
                    return filteredOrders
                }
                catch(err){
                    if (err.name =="TypeError"){
                        console.log("Initializing data...")
                    }
                    else{
                        console.error(err)
                    }
                }
            },
            calTotal(){
                try{
                    let total = 0.0
                    
                    for(const x in this.filteredProducts){
                        total += this.filteredProducts[x].TOTAL
                    }
                    return total.toFixed(2)
                }
                catch(err){
                    if (err.name =="TypeError"){
                        console.log("here")
                    }
                    else{
                        console.error(err)
                    }
                }
            }
        },

        methods: {
            filterOrdersByName(orders){
                return orders.filter(p => { return p.ORDER_NAME.toLowerCase().indexOf(this.searchText.toLowerCase()) != -1})
            },

            filterOrdersByProduct(orders) {
                return orders.filter(p => { return p.PRODUCTS.toLowerCase().indexOf(this.searchText.toLowerCase()) != -1})
            },


        }
        
    }
</script>
<style scoped>
table, th, td {
  /* border:1px solid black; */
}
</style>