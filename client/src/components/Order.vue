<template>
    <h3>{{order.ORDER_NAME}}</h3>
</template>

<script>
export default {
    name: 'Order',
    props: {
        order: Object
    }
}
</script>