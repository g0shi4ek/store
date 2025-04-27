import axios from 'axios';

const API_URL = '/api';

export const register = (username, password, role) => {
    return axios.post(`${API_URL}/register`, { username, password, role });
};

export const login = (username, password) => {
    return axios.post(`${API_URL}/login`, { username, password })
        .then(response => {
            const { token, role } = response.data;
            localStorage.setItem('user', JSON.stringify({ token, role }));
            axios.defaults.headers.common['Authorization'] = token;
            return response.data;
        });
};

export const logout = () => {
    localStorage.removeItem('user');
    delete axios.defaults.headers.common['Authorization'];
    window.location.href = "/login";
};

export const getCurrentUser = () => {
    return JSON.parse(localStorage.getItem('user'));
};