export class Order {
    constructor(
        public id: number,
        public symbol: string,
        public type: string,
        public side: string,
        public price: number,
        public amount: number,
    ) { }

}
