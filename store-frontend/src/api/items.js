import axios from 'axios';

const API_URL = '/api';

export const getAllItems = () => {
    return axios.get(`${API_URL}/items/all`);
};

export const getItem = (id) => {
    return axios.get(`${API_URL}/items/${id}`);
};

export const createItem = (itemData) => {
    return axios.post(`${API_URL}/manager/items/`, itemData);
};

export const updateItem = (id, itemData) => {
    return axios.post(`${API_URL}/manager/items/${id}`, itemData);
};