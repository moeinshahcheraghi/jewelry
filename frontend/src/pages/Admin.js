// src/pages/Admin.js
import React, { useState, useContext } from 'react';
import API from '../services/api';
import { AuthContext } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

const Admin = () => {
  const { auth } = useContext(AuthContext);
  const navigate = useNavigate();

  const [form, setForm] = useState({
    name: '',
    description: '',
    price: '',
    quantity: '',
  });

  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  if (!auth.isAdmin) {
    navigate('/');
  }

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // تبدیل قیمت و تعداد به عدد
    const payload = {
      name: form.name,
      description: form.description,
      price: parseFloat(form.price),
      quantity: parseInt(form.quantity, 10),
    };

    API.post('/admin/products', payload)
      .then((res) => {
        setSuccess('Product added successfully!');
        setError('');
        setForm({ name: '', description: '', price: '', quantity: '' });
      })
      .catch((err) => {
        setError(err.response?.data?.error || 'Failed to add product');
        setSuccess('');
      });
  };

  return (
    <div style={styles.container}>
      <h1>Admin Panel - Add Product</h1>
      <form onSubmit={handleSubmit} style={styles.form}>
        <input
          type="text"
          name="name"
          placeholder="Product Name"
          value={form.name}
          onChange={handleChange}
          required
          style={styles.input}
        />
        <textarea
          name="description"
          placeholder="Product Description"
          value={form.description}
          onChange={handleChange}
          required
          style={styles.textarea}
        />
        <input
          type="number"
          step="0.01"
          name="price"
          placeholder="Price"
          value={form.price}
          onChange={handleChange}
          required
          style={styles.input}
        />
        <input
          type="number"
          name="quantity"
          placeholder="Quantity"
          value={form.quantity}
          onChange={handleChange}
          required
          style={styles.input}
        />
        {error && <p style={styles.error}>{error}</p>}
        {success && <p style={styles.success}>{success}</p>}
        <button type="submit" style={styles.button}>
          Add Product
        </button>
      </form>
    </div>
  );
};

const styles = {
  container: {
    padding: '20px',
    maxWidth: '600px',
    margin: '0 auto',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
  },
  input: {
    padding: '10px',
    marginBottom: '10px',
    fontSize: '16px',
  },
  textarea: {
    padding: '10px',
    marginBottom: '10px',
    fontSize: '16px',
    height: '100px',
  },
  button: {
    padding: '10px',
    backgroundColor: '#555555',
    color: '#fff',
    border: 'none',
    cursor: 'pointer',
    fontSize: '16px',
  },
  error: {
    color: 'red',
    marginBottom: '10px',
  },
  success: {
    color: 'green',
    marginBottom: '10px',
  },
};

export default Admin;

