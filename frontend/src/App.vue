<template>
  <div id="app">
    <div class="container">
      <h3>Sartura task app</h3>
      <div class="row">
        <multiselect v-model="value" :options="options"></multiselect>
        <div class="col-12 col-md-6">
          <LivePrice></LivePrice>
        </div>
        <div class="col-12 col-md-6">
          <LivePrice></LivePrice>
        </div>
      </div>
      <div class="row">
        <div class="col-12 col-md-8">
          <ListOrders></ListOrders>
        </div>
        <div class="col-12 col-md-4">
          <CreateOrder></CreateOrder>
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

@Component({
  components: {
    CreateOrder,
    ListOrders,
    LivePrice,
    Multiselect,
  },
})
export default class App extends Vue {
  data() {
    return {
      value: null,
      options: ['list', 'of', 'options'],
    };
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>