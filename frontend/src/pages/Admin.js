// frontend/src/pages/Admin.js
import React, { useState } from 'react';
import api from '../services/api';

const Admin = () => {
    const [form, setForm] = useState({
        name: '',
        description: '',
        price: '',
        quantity: '',
    });

    const handleChange = (e) => {
        setForm({...form, [e.target.name]: e.target.value});
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await api.post('/admin/products', form);
            alert('Product added successfully');
            setForm({ name: '', description: '', price: '', quantity: '' });
        } catch (error) {
            alert(error.response.data.error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input name="name" placeholder="Product Name" value={form.name} onChange={handleChange} required />
            <textarea name="description" placeholder="Description" value={form.description} onChange={handleChange}></textarea>
            <input name="price" type="number" step="0.01" placeholder="Price" value={form.price} onChange={handleChange} required />
            <input name="quantity" type="number" placeholder="Quantity" value={form.quantity} onChange={handleChange} required />
            <button type="submit">Add Product</button>
        </form>
    );
};

export default Admin;

