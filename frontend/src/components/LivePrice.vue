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
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
// @ts-ignore
import VueNativeSock from 'vue-native-websocket';
import { TradeStreamData } from '../models/TradeStreamData';

@Component
export default class LivePrice extends Vue {
    @Prop() coin!: string;
    private wss!: WebSocket;

    data(){
        return {
            tradeStreamsList: new Array<TradeStreamData>(),
        }
    }

    @Watch('coin')
    onPropertyChanged(newCoin: string, oldCoin: string) {
        let formatedCoin = newCoin.replace("/","").toLowerCase(); 
        if (this.wss !== undefined) {
            this.wss.close();
        }
        this.wss = this.openWebSocketsConnection(formatedCoin);
    }

    private openWebSocketsConnection(coinLabel: string): WebSocket {
        const websocketUrl: string = 'wss://stream.binance.com:9443/ws/' + coinLabel + '@aggTrade';
        const wss = new WebSocket(websocketUrl);
        wss.addEventListener('open', (event) => {
            wss.send(JSON.stringify({
                'method': 'SUBSCRIBE',
                'params': [
                    coinLabel+'@aggTrade'
                ],
                'id': 1,
                }));
        });

        const thisRef = this;
        const thisListRef = this.$data.tradeStreamsList;
        wss.addEventListener('message', (event) => {
            const jsonObj = JSON.parse(event.data);
            if (jsonObj['e'] = "aggTrade") {
                thisRef.pushToArray(jsonObj, thisListRef);
            }
        });

        return wss;
    }

    private pushToArray(value: any, tradeStreamsList: TradeStreamData[]) {
        const obj: TradeStreamData = this.prepeareObj(value);
        if (tradeStreamsList.length < 7) {
            tradeStreamsList.push(obj);
        } else {
            tradeStreamsList.shift();
            tradeStreamsList.push(obj);
        }
    }

    /* tslint:disable:no-string-literal */
    private prepeareObj(value: any): TradeStreamData {
        const TradeStreamDataObj: TradeStreamData = new TradeStreamData(
            value['e'], value['E'], value['s'],
            value['a'], value['p'], value['q'],
            value['f'], value['l'], new Date(parseInt(value['T'])),
            value['m'], value['M'],
        );
        return TradeStreamDataObj;
    }
}
</script>
