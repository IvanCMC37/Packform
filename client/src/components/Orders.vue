<template>
    <div>
        <table>
            <tr>
                <th>Order name</th>
                <th>Customer Company</th>
                <th>Customer name</th>
                <th>order date</th>
                <th>Delivered Amount</th>
                <th>Total Amount</th>
            </tr>
            <tr :key="order.ORDER_NAME" v-for="order in orders.orders">
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
        }
    }
</script>
<style scoped>
table, th, td {
  border:1px solid black;
}
</style>