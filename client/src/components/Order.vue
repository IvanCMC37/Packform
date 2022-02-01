<template>
    <td>{{order.ORDER_NAME}}</td>
    <td>{{order.COMPANY_NAME}}</td>
    <td>{{order.NAME}}</td>
    <td>{{order.ORDER_DATE}}</td>
    <td>{{order.DELIVERED_AMOUNT}}</td>
    <td>{{order.TOTAL}}</td>
</template>

<script>
export default {
    name: 'Order',
    props: {
        order: Object
    }
}
</script>
<style>
table, th, td {
  border:1px solid black;
}
</style>