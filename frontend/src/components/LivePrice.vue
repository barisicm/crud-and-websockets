<template>
<div>
        <h3>Price</h3>
        <table class="table">
            <thead>
                <tr>
                    <th scope="col">Symbol</th>
                    <th scope="col">PRICE</th>
                    <th scope="col">Quantity</th>
                    <th scope="col">Trade Time</th>
                </tr>
            </thead>
            <tbody>

                 <tr v-for="item in tradeStreamsList" v-bind:key=item.id>
                    <td>{{ item.symbol }}</td>
                    <td>{{ item.price }}</td>
                    <td>{{ item.quantity }}</td>
                    <td>{{ item.tradeTime.getHours() + ":" +
                           item.tradeTime.getMinutes() + ":" +
                           item.tradeTime.getSeconds()  }}</td>
                </tr>
            </tbody>
        </table>
</div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
// @ts-ignore
import VueNativeSock from 'vue-native-websocket';
import { TradeStreamData } from '../models/TradeStreamData';

@Component
export default class LivePrice extends Vue {
    @Prop() Coin!: string;

    data(){
        return {
            tradeStreamsList: new Array<TradeStreamData>(),
        }
    }

    mounted() {

        //let websocketUrl :string = 'wss://stream.binance.com:9443/ws/' + this.Coin + '@aggTrade';
        let wss = new WebSocket('wss://stream.binance.com:9443/ws/btcusdt@aggTrade');
        // Connection opened
        let api_key = process.env.VUE_APP_CRYPTO_API_KEY
        wss.addEventListener('open', function (event) {
            wss.send(JSON.stringify({
                "method": "SUBSCRIBE",
                "params": [
                    "btcusdt@aggTrade"
                ],
                "id": 1
                }))
        });

        let thisListRef = this.$data.tradeStreamsList
        wss.addEventListener('message', function (event) {
            let jsonObj = JSON.parse(event.data)
            if(jsonObj['e']="aggTrade"){
                console.log(jsonObj)
                console.log(thisListRef)
                pushToArray(jsonObj, thisListRef);
            }
        });

        function pushToArray(value: any, tradeStreamsList: Array<TradeStreamData>){
            let obj: TradeStreamData = prepeareObj(value);
            if(tradeStreamsList.length < 7){
                tradeStreamsList.push(obj)
            } else {
                tradeStreamsList.shift()
                tradeStreamsList.push(obj)
            }
        }

        function prepeareObj(value: any): TradeStreamData{
            let TradeStreamDataObj: TradeStreamData = new TradeStreamData(
                value['e'],value['E'],value['s'],
                value['a'],value['p'],value['q'],
                value['f'],value['l'],new Date(parseInt(value['T'])),
                value['m'],value['M'],
            );
            return TradeStreamDataObj
        }

    }
}
</script>
