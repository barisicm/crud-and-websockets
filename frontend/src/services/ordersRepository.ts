import Repository from './Repository';
import { Order } from '../models/Order';
import { AxiosResponse } from 'axios';

export default {
    getAll(): Promise<AxiosResponse<Order[]>> {
        return Repository.get(`/list`);
    },

    deleteOrder(orderId: number) {
        return Repository.delete(`/delete/${orderId}`);
    },

    createOrder(order: Order) {
        return Repository.post(`/orders`, order, config);
    },
};

const config = {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',
    },
};
