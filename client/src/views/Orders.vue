<template>
  <div class="container">
    <div class="search-wrapper panel-heading col-sm-12">
        <span><i class="fas fa-search"></i>Search
            <input class="searchTextTop" type="text" v-model="searchTextTop" placeholder="Search by Order name/ Product" /> <br> <br>
        </span>
    </div>
    <div>
        <span>Created date</span>
        <Datepicker id="picker-manual" :modelValue="date" :startDate="dataStartDate" @update:modelValue="dateRangeSetter" @cleared="dateRangeSetter" :enableTimePicker="true" range/>
    </div>
    <div v-if="(filteredProducts != undefined)">
        <div v-if="filteredProducts.length >0">
            <h4>Total amount: {{calTotal}}</h4>
            <table>
                <tr>
                    <th class="unhoverable">Order name</th>
                    <th class="unhoverable">Customer Company</th>
                    <th class="unhoverable">Customer name</th>
                    <th class="hoverable" @click="sortedOrders()">Order date 
                        <span class="icons-wrapper">
                            <i v-bind:class="{isActive: !currentSortDir}" class="fas fa-sort-down"></i>
                            <i v-bind:class="{isActive: currentSortDir}" class="fas fa-sort-up"></i>
                        </span>
                    </th>
                    <th class="unhoverable">Delivered Amount</th>
                    <th class="unhoverable">Total Amount</th>
                </tr>
                <tr :key="order.ORDER_NAME" v-for="order in filteredProducts">
                    <Order :order="order" />
                </tr>
            </table>
            <div class="page-navigation">
                <select @change="rangeSelector" name="range" id="range">
                    <option value="2">2/page</option>
                    <option selected="selected" value="5">5/page</option>
                    <option value="10">10/page</option>
                    <option :value="searchTextFilteredOrders">Max/page</option>
                </select>
                <span>Total: {{maxPage}} page(s)</span>
                <button v-bind:class="{notActive: noPrevPage}" @click="prevPage"><i class="fas fa-backward"></i></button> 
                <span>{{currentPage}}</span>
                <button v-bind:class="{notActive: noNextPage}"  @click="nextPage"><i class="fas fa-forward"></i></button>
                <span>Go to page:
                    <input @keyup.enter=pageJumper class="searchTextBottom" type="number" v-model="searchTextBottom" min="1" :max=maxPage step="1" 
                    placeholder=""/>
                    <br> <br>
                    <label v-bind:class="{isActive: hasError}">{{customLabel}}</label>
                </span>
            </div>       
        </div>
        <div class="no-record" v-if="filteredProducts.length == 0">
            <span class="no-record">No matching records...</span>
        </div>
    </div>
    <div v-else class="loading"><span class="loading">Loading data...</span></div>
    <div class="goBack"><router-link to="/">Go back</router-link></div>
    
  </div>
</template>

<script>
    import Order from '../components/Order'
    import Datepicker from 'vue3-date-time-picker';
    import 'vue3-date-time-picker/dist/main.css'
    import { ref } from 'vue';
    const date = ref();

    export default {
        name:"Orders",
        // props: {
        //     orders: Object,
        // },
        components: {
            Order,
            Datepicker,
        },
        setup() {
            const date = ref();
            
            return {
                date,
            }
        },
        data() {
            return {
                orders: [],
                searchTextTop: "",
                searchTextBottom: 1,
                currentSortDir:true,
                pageSize:5,
                maxPage: 1,
                currentPage:1,
                searchTextFilteredOrders: 0,
                noNextPage: true,
                noPrevPage: false,
                hasError: false,
                customLabel: "",
                date: null,
                filterStartDate : null,
                filterendDate: null,
                dataStartDate: null,
                dataEndDate: null,
            }
        },
        computed: {
            filteredProducts() {
                try{
                    //Filter by text field
                    let product = this.orders.orders.filter(p => { 
                            return p.ORDER_NAME.toLowerCase().indexOf(this.searchTextTop.toLowerCase()) != -1
                            || p.PRODUCTS.toLowerCase().indexOf(this.searchTextTop.toLowerCase()) != -1
                        })

                    //sort by date
                    if(this.currentSortDir){
                        product = product.sort((a, b) => new Date(a.ORDER_DATE) - new Date(b.ORDER_DATE))
                        // if(product.length>0){
                        //     this.dataStartDate = product[0].ORDER_DATE
                        //     this.dataEndDate = product[product.length - 1].ORDER_DATE
                        // }
                    }
                    else{
                        product = product.sort((b, a) => new Date(a.ORDER_DATE) - new Date(b.ORDER_DATE))
                    }

                    //apply filter if datepicker is enabled
                    if(this.filterStartDate){
                        product = product.filter(p =>{
                            return (new Date(p.ORDER_DATE) >= this.filterStartDate && new Date(p.ORDER_DATE) <= this.filterEndDate )
                        })
                    }

                    //calculated if current page passed the maxpage or not
                    this.searchTextFilteredOrders = product.length;
                    this.currentPageControl()

                    //filter by pagesize
                    product = product.filter((row, index) => {
                        let start = (this.currentPage-1)*this.pageSize;
                        let end = this.currentPage*this.pageSize;
                        if(index >= start && index < end) return true;
                    });

                    //control page button on/off
                    this.pageButtonControl();

                    return product
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
                    }
                    else{
                        console.error(err)
                    }
                }
            },
        },

        methods: {
            sortedOrders() {
                this.currentSortDir= this.currentSortDir === true? false:true
            },
            nextPage:function() {
                if((this.currentPage*this.pageSize) < this.searchTextFilteredOrders) this.currentPage++;
            },
            prevPage:function() {
                if(this.currentPage > 1) this.currentPage--;
            },
            maxPageCal(){
                return Math.ceil(this.searchTextFilteredOrders/this.pageSize);
            },
            pageButtonControl(){
                this.maxPage = this.maxPageCal();

                if(this.currentPage == 1) this.noPrevPage=true
                else this.noPrevPage = false

                if(this.currentPage == this.maxPage) this.noNextPage=true
                else this.noNextPage = false
            },
            currentPageControl(){
                if (this.currentPage >this.searchTextFilteredOrders/this.pageSize){
                    let maxpageavailable = this.maxPageCal()
                    this.currentPage = maxpageavailable > 0 ? maxpageavailable : 1 
                }
            },
            pageJumper(){
                if(Number.isInteger(this.searchTextBottom) && this.searchTextBottom > 0){  
                    if(this.searchTextBottom > this.maxPageCal()){
                        this.currentPage = this.maxPageCal()
                        this.customLabel = "Number is out of range, headed to last page instead..."
                        this.hasError = true
                    } 
                    else{
                        this.currentPage = this.searchTextBottom
                        this.hasError = false
                    }
                }
                else {
                    this.customLabel = "Please input a correct number..."
                    this.hasError = true
                }

                this.searchTextBottom = 1
            },
            rangeSelector(e) {
                this.pageSize = e.target.value
            },
            dateRangeSetter(value) {
                if(value != null){
                    this.filterStartDate = value[0]
                    this.filterEndDate = value[1]
                    console.log("Assigned dates")                   
                }
                else{
                    this.filterStartDate = null
                    this.filterEndDate = null
                }
            },
            async fetchOrders() {
                console.log("Loading dataset")
                const res = await fetch('api/orders')
                .catch(function(error){
                console.error(error)
                });
                const data = await res.json()
                console.log("Finished loading dataset")
                return data
            }, 
        },
        async created() {
            this.orders = await this.fetchOrders()
        } 
    }
</script>

<style scoped>
.fa-search{
    padding-right: 12px;
}
th, tr, td{
    text-align: left;
    height: 70px;
}
tr:not(:first-child){
    border-bottom: 3px grey solid;
}
table{
    width: 100%;
    border-collapse: collapse;
}
input, span{
    height: 30px;
    font-size: 20px;
}
input.searchTextTop{
    margin-left: 20px;
    width: 80%;
}
.no-record , .loading{
    font-size: 35px;
    text-align: center;
    padding: 50px;
}
tr .unhoverable{
    cursor: default;
    user-select: none;
}
tr .hoverable{
    cursor: pointer;
    user-select: none;
}
tr .hoverable:hover{
    color: blue;
    font-weight: bold;
}
.page-navigation{
    text-align: center;
    padding-top: 16px;
}
.page-navigation button{
    background-color: transparent;
    border: none;
    cursor: pointer;
}
.page-navigation button.notActive{
    visibility: hidden;
}
.page-navigation label{
    visibility: hidden;
}
.page-navigation label.isActive{
    visibility: visible;
    color: red;
}
.page-navigation input{
    margin-left: 30px;
    font-size: 12px;
}
.page-navigation{
    user-select: none;
}
.page-navigation span{
    font-size: 12px;
    padding-left: 12px;
    padding-right: 12px;
}
#picker-manual{
    width: 500px;
}
.goBack{
    text-align: center;
}
.icons-wrapper {
  display: inline-flex;
  flex-direction: column;
  padding-left: 4px;
}
.icons-wrapper i.fa{
  line-height: 0;
}
.hoverable .isActive{
    visibility: collapse;
}
.fa-sort-up{
    margin-top: -16px;
}
</style>