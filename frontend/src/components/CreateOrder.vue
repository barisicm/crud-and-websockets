<template>
    <div>
        <h3>Create Order</h3>
        <div>
            <div class="form-group">
                <label for="coinSymbol">Coin</label>
                <input type="coinSymbol" class="form-control" id="coinSymbol" disabled v-bind:value="coinSymbol">
            </div>
            <div class="form-group">
                <label for="buySell">BUY / SELL</label>
                <select class="form-control" id="buySell" v-model="side">
                    <option value="buy">Buy</option>
                    <option value="sell">Sell</option>
                </select>
            </div>
            <div class="form-group">
                <label for="price">Price</label>
                <input type="number" class="form-control" id="price" v-model="price">
            </div>
            <div class="form-group">
                <label for="amount">Amount</label>
                <input type="number" class="form-control" id="amount" v-model="amount">
            </div>
            <button class="btn btn-primary" @click="createOrder()">Create Order</button>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
    import { Order } from '../models/Order';
    import ordersRepository from '../services/ordersRepository';
    // @ts-ignore
    import qs from 'qs';

    @Component
    export default class CreateOrder extends Vue {
        @Prop() private coinSymbol!: string;
        
        private price: number = 0;
        private amount: number = 0;
        private side: string = "";
        private order!: Order;


        private createOrder() {
            this.order = new Order(this.coinSymbol,"market",  this.side, this.price,this.amount);

            if(this.validateData(this.order)){
                ordersRepository.createOrder(qs.stringify(this.order)).then((response) => {
                    this.order.id = response.data.order.id
                }).catch((err) => {
                    console.log(err)
                }).finally(() => {
                    this.$emit("newOrder", this.order);
                })
            } else {
                alert('Please submit correct data.')
            }

            
            
        }

        private validateData(order: Order){
            if(order.symbol.length > 0 && order.side.length > 0 && order.price > 0 && order.amount > 0){
                return true;
            } else {
                return false;
            }
        }

        
    
    }
</script>