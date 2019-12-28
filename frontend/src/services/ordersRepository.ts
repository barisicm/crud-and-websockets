import Repository from './Repository';
import { Order } from '../models/Order';
import { OrderList } from '@/models/OrderListInterface';
import { AxiosResponse } from 'axios';

export default {
    getAll(): Promise<AxiosResponse<Order[]>> {
        return Repository.get(`/list`);
    },

    deleteOrder(orderId: number) {
        return Repository.delete(`/delete/${orderId}`);
    },

    createOrder(order: Order) {
        return Repository.post(`/order`, order,config);
    },
};

const config = {
    headers: {
      'Content-Type': 'application/json'
    }
}