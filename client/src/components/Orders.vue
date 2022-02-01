<template>
    <div>
        <div :key="order.ORDER_NAME" v-for="order in orders.orders">
            <Order :order="order" />
        </div>
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

</style>