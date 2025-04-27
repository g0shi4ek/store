import axios from 'axios';

const API_URL = '/api';

export const bookRooms = (storeId, itemId, amount) => {
    return axios.post(`${API_URL}/manager/bookings/room`, 
        { store_id: storeId, item_id: parseInt(itemId), amount: amount }
    );
};

export const cancelRoomBooking = (id) => {
    return axios.delete(`${API_URL}/manager/bookings/room/${id}`);
};

export const getAllRoomBookings = () => {
    return axios.get(`${API_URL}/manager/bookings/all_rooms`);
};

export const buyRoomBooking = (id) => {
    return axios.post(`${API_URL}/manager/bookings/buy_room/${id}`, {});
};

export const bookItems = (itemId, amount) => {
    return axios.post(`${API_URL}/seller/bookings/item`, 
        { item_id: parseInt(itemId), amount: amount }
    );
};

export const cancelItemBooking = (id) => {
    return axios.delete(`${API_URL}/seller/bookings/item/${id}`);
};

export const getAllItemBookings = () => {
    return axios.get(`${API_URL}/seller/bookings/all_items`);
};

export const buyItemBooking = (id) => {
    return axios.post(`${API_URL}/seller/bookings/buy_item/${id}`, {});
};