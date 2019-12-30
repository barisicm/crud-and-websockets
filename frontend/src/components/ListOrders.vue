<template>
    <div class="table-responsive">
        <h3>List Orders</h3>
        <table class="table">
            <thead>
                <tr>
                    <th scope="col">Coin</th>
                    <th scope="col">Type</th>
                    <th scope="col">Side</th>
                    <th scope="col">Price</th>
                    <th scope="col">Amount</th>
                    <th scope="col"></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(order, index) in allOrders" v-bind:key=order.id>
                    <td>{{ order.symbol }}</td>
                    <td>{{ order.type }}</td>
                    <td>{{ order.side }}</td>
                    <td>{{ order.price }}</td>
                    <td>{{ order.amount }}</td>
                    <td class="cursor-pointer"><a @click="removeOrder(order.id, index)">Delete</a></td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import ordersRepository from '../services/ordersRepository';
import { Order } from '../models/Order';

@Component
export default class ListOrders extends Vue {
    @Prop() public newOrder!: Order;

    @Watch('newOrder')
    private onPropertyChanged(newOrder: Order, oldOrder: Order) {
        this.$data.allOrders.push(newOrder);
    }

    public data() {
        return {
            allOrders: new Array<Order>(),
        };
    }

    public mounted() {
        const dataObjRef = this.$data.allOrders;
        ordersRepository.getAll()
        .then((response) => {
            if (response.data) {
                response.data.forEach((order: Order) => {
                    dataObjRef.push(order);
                });
            }
        }).catch((err) => alert(err));
    }

    private removeOrder(orderId: number, orderIndex: number) {
        ordersRepository.deleteOrder(orderId).then(
            (response) => this.$data.allOrders.splice(orderIndex, 1),
        ).catch((err) => alert(err));
    }
}
</script>

<style lang="css">
    .cursor-pointer{
        cursor: pointer
    }
</style>