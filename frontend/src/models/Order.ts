export class Order {
    constructor(
        public symbol: string,
        public type: string,
        public side: string,
        public price: number,
        public amount: number,
        public id?: number,
    ) { }
}
