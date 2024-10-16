// src/services/api.js
import axios from 'axios';

const API = axios.create({
  baseURL: 'http://backend:8080', // آدرس Backend شما
});

// افزودن توکن به هدر درخواست‌ها
API.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default API;

