// frontend/src/pages/Login.js
import React, { useState } from 'react';
import api from '../services/api';
import { useHistory } from 'react-router-dom';

const Login = () => {
    const [form, setForm] = useState({
        username: '',
        password: '',
    });
    const history = useHistory();

    const handleChange = (e) => {
        setForm({...form, [e.target.name]: e.target.value});
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await api.post('/login', form);
            localStorage.setItem('token', response.data.token);
            alert('Login successful');
            history.push('/');
        } catch (error) {
            alert(error.response.data.error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input name="username" placeholder="Username" value={form.username} onChange={handleChange} required />
            <input name="password" type="password" placeholder="Password" value={form.password} onChange={handleChange} required />
            <button type="submit">Login</button>
        </form>
    );
};

export default Login;

