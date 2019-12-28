<template>
  <div id="app">
    <div class="container">
      <div class="mt-4 mb-4"><h1>Price tracking</h1></div>
      <div class="row">
        <multiselect v-model="coinSymbol" :options="options"></multiselect>
        <div class="col-12">
          <LivePrice :coin="coinSymbol"></LivePrice>
        </div>
      </div>
      <div class="row">
        <div class="col-12 col-md-8">
          <ListOrders :newOrder="order"></ListOrders>
        </div>
        <div class="col-12 col-md-4">
          <CreateOrder :coinSymbol="coinSymbol" @newOrder="addToListOrders"></CreateOrder>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import CreateOrder from './components/CreateOrder.vue';
import ListOrders from './components/ListOrders.vue';
import LivePrice from './components/LivePrice.vue';
import Multiselect from 'vue-multiselect';
import { Order } from './models/Order';

@Component({
  components: {
    CreateOrder,
    ListOrders,
    LivePrice,
    Multiselect,
  },
})
export default class App extends Vue {
  public data() {
    return {
      coinSymbol: null,
      options: ['BTC/USDT', 'BTC/ETH', 'BTC/ALGO', 'ETH/USDT', 'ETH/BTC', 'ETH/XLM'],
      order: null,
    };
  }

  public addToListOrders(order: Order) {
    this.$data.order = order;
  }

}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>