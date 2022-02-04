<template>
    <footer class="footer">
        <div>
            <p>Ivan Chan, 2022</p>
            <router-link to="/orders">Orders</router-link>
        </div>
    </footer>
</template>

<script>
    export default {
        name:"Footer",
    }
</script>

<style scoped>
.footer{
    text-align: center;
    padding-top: 12px;
    text-align: center;
    font-style: italic;
    font-size: 14px;
    user-select: none;
}
p{
    cursor: default;
}
</style>