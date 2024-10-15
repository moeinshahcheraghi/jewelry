// frontend/src/pages/Register.js
import React, { useState } from 'react';
import api from '../services/api';

const Register = () => {
    const [form, setForm] = useState({
        username: '',
        email: '',
        password: '',
    });

    const handleChange = (e) => {
        setForm({...form, [e.target.name]: e.target.value});
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await api.post('/register', form);
            alert('Registration successful');
        } catch (error) {
            alert(error.response.data.error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input name="username" placeholder="Username" value={form.username} onChange={handleChange} required />
            <input name="email" type="email" placeholder="Email" value={form.email} onChange={handleChange} required />
            <input name="password" type="password" placeholder="Password" value={form.password} onChange={handleChange} required />
            <button type="submit">Register</button>
        </form>
    );
};

export default Register;

