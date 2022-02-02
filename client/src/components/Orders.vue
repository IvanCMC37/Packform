<template>
    <div>
        <div class="search-wrapper panel-heading col-sm-12">
            <input type="text" v-model="search" placeholder="Search" /> <br> <br>
        </div>  
        <h4>Total amount: ${{orders.grand}}</h4>
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
        name:"Orders",
        props: {
            orders: Object
        },
        components: {
            Order
        },
        data() {
            return {
                search: ""
            }
        },
        computed: {
            filteredProducts() {
                try{
                    return this.orders.orders.filter(p => {
                    /*
                    in this example we just check if the search string 
                    is a substring of the product name (case insensitive)
                    */
                    return p.PRODUCTS.toLowerCase().includes(this.search.toLowerCase())
                    })
                }
                catch(err){
                    if (err.name =="TypeError"){
                        console.log("Initializing data...")
                    }
                    else{
                        console.error(err)
                    }
                }
            }
        }
        
    }
</script>
<style scoped>
table, th, td {
  /* border:1px solid black; */
}
</style>