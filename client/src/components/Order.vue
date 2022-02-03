<template>
    <td>{{order.ORDER_NAME}}</td>
    <td>{{order.COMPANY_NAME}}</td>
    <td>{{order.NAME}}</td>
    <td>{{datetimeFormat(order.ORDER_DATE)}}</td>
    <td>{{currencyFormat(order.DELIVERED_AMOUNT)}}</td>
    <td>{{currencyFormat(order.TOTAL)}}</td>
</template>

<script>
export default {
    name: 'Order',
    props: {
        order: Object,
    },
    methods: {
        currencyFormat(value) {
            if (value !=null){
                return '$' + value.toFixed(2)
            }
        },
        datetimeFormat(value) {
            // console.log(value)
            var options = {
            year:'numeric',
            month: 'short', day: 'numeric',
            hour: 'numeric', minute: 'numeric',
            timeZone: 'Australia/Melbourne',
            };
            if (value !=null){
                const date = new Date(Date.parse(value))
                return new Intl.DateTimeFormat('en-AU', options).format(date)
            }
        }
    }
}
</script>
<style>
</style>