export class TradeStreamData {
    constructor(
        public eventType: string,
        public eventTime: number,
        public symbol: string,
        public tradeId: number,
        public price: number,
        public quantity: number,
        public buyerOrderId: number,
        public sellerOrderId: number,
        public tradeTime: Date,
        public marketMaker: boolean,
        public ignore: boolean,
    ) { }
}
