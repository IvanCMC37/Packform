<template>
    <div>
        <div class="search-wrapper panel-heading col-sm-12">
            <span>Search
                <input class="searchTextTop" type="text" v-model="searchTextTop" placeholder="Search by Order name/ Product" /> <br> <br>
            </span>
        </div>
        <div>
            <span>Created date</span>
        </div>
        <div v-if="(filteredProducts != undefined)">
            <div v-if="filteredProducts.length >0">
                <h4>Total amount: {{calTotal}}</h4>
                <table>
                    <tr>
                        <th class="unhoverable">Order name</th>
                        <th class="unhoverable">Customer Company</th>
                        <th class="unhoverable">Customer name</th>
                        <th class="hoverable" @click="sortedOrders()">Order date</th>
                        <th class="unhoverable">Delivered Amount</th>
                        <th class="unhoverable">Total Amount</th>
                    </tr>
                    <tr :key="order.ORDER_NAME" v-for="order in filteredProducts">
                        <Order :order="order" />
                    </tr>
                </table>
                <div class="page-navigation">
                    <select @change="rangeSelector" name="cars" id="cars">
                        <option value="2">2/page</option>
                        <option selected="selected" value="5">5/page</option>
                        <option value="8">8/page</option>
                        <option :value="searchTextFilteredOrders">Max/page</option>
                    </select>
                    <span>Total: {{maxPage}} page(s)</span>
                    <button v-bind:class="{notActive: noPrevPage}" @click="prevPage">Previous</button> 
                    <span>{{currentPage}}</span>
                    <button v-bind:class="{notActive: noNextPage}"  @click="nextPage">Next</button>
                    <span>Go to page:
                        <input @keyup.enter=pageJumper class="searchTextBottom" type="number" v-model="searchTextBottom" min="1" :max=maxPage step="1" 
                        placeholder=""/>
                        <label v-bind:class="{isActive: hasError}">{{customLabel}}</label>
                    </span>
                </div>       
            </div>
            <div class="no-record" v-if="filteredProducts.length == 0">
                <span class="no-record">No matching records...</span>
            </div>
        </div>
        <div v-else class="loading"><span class="loading">Loading data...</span></div>
    </div>
</template>

<script>
    import Order from './Order'
    // import VCalendar from 'v-calendar';

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
                    this.searchTextFilteredOrders = product.length;
                    console.log(this.searchTextFilteredOrders)
                    console.log("this.currentPage: "+ this.currentPage)
                    //calculated if current page passed the maxpage or not
                    // this.pageJumper()
                    this.currentPageControl()

                    //sort by date
                    if(this.currentSortDir){
                        product = product.sort((a, b) => new Date(a.ORDER_DATE) - new Date(b.ORDER_DATE))
                    }
                    else{
                        product = product.sort((b, a) => new Date(a.ORDER_DATE) - new Date(b.ORDER_DATE))
                    }
                    console.log("prodcut ="+product.length)
                    //filter by pagesize
                    // console.log("this.currentPage: "+ this.currentPage)
                    // console.log("this.pageSize: "+ this.pageSize)
                    // console.log("end: "+ this.currentPage*this.pageSize)
                    product= product.filter((row, index) => {
                            let start = (this.currentPage-1)*this.pageSize;
                            let end = this.currentPage*this.pageSize;
                            if(index >= start && index < end) return true;
                        });

                    //control page button on/off
                    this.pageButtonControl();
                    console.log("prodcut ="+product.length)
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
            }
        }      
    }
</script>

<style scoped>
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
.page-navigation button.notActive{
    visibility: hidden;
}
.page-navigation label{
    visibility: hidden;
}
.page-navigation label.isActive{
    visibility: visible;
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
</style>