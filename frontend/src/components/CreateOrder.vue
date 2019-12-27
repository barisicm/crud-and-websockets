<template>
    <div>
        <h3>Create Order</h3>
        <form>
            <div class="form-group">
                <label for="coinSymbol">Coin</label>
                <input type="coinSymbol" class="form-control" id="coinSymbol" disabled value="BTC/XLM">
            </div>
            <div class="form-group">
                <label for="buySell">BUY / SELL</label>
                <select class="form-control" id="buySell">
                    <option>Buy</option>
                    <option>Sell</option>
                </select>
            </div>
            <div class="form-group">
                <label for="price">Price</label>
                <input type="number" class="form-control" id="price">
            </div>
            <div class="form-group">
                <label for="amount">Amount</label>
                <input type="number" class="form-control" id="amount">
            </div>
            <button type="submit" class="btn btn-primary">Create Order</button>
        </form>
    </div>
</template>

<script lang="ts">
    import {
        Component,
        Prop,
        Vue,
    } from 'vue-property-decorator';

    @Component
    export default class CreateOrder extends Vue {
        @Prop() private msg!: string;
    }
</script>