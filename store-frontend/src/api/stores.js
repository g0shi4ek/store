import axios from 'axios';

const API_URL = '/api';

export const getStore = (id) => {
    return axios.get(`${API_URL}/stores/${id}`);
};

export const getAllStores = () => {
    return axios.get(`${API_URL}/stores/all`);
};
