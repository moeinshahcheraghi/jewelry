// frontend/src/pages/Home.js
import React, { useState, useEffect } from 'react';
import api from '../services/api';
import ProductCard from '../components/ProductCard';
import StoryForm from '../components/StoryForm';
import StoryList from '../components/StoryList';

const Home = () => {
    const [products, setProducts] = useState([]);
    const [query, setQuery] = useState('');

    const fetchProducts = async () => {
        try {
            const response = await api.get('/products');
            setProducts(response.data.data);
        } catch (error) {
            console.error(error);
        }
    };

    const searchProducts = async (e) => {
        e.preventDefault();
        try {
            const response = await api.get(`/search?q=${query}`);
            setProducts(response.data.data);
        } catch (error) {
            console.error(error);
        }
    };

    const handleStoryCreated = (newStory) => {
        // می‌توانید اینجا لیست داستان‌ها را به‌روزرسانی کنید
    };

    useEffect(() => {
        fetchProducts();
    }, []);

    return (
        <div>
            <form onSubmit={searchProducts}>
                <input
                    type="text"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    placeholder="Search products..."
                />
                <button type="submit">Search</button>
            </form>
            <div>
                {products.map(product => (
                    <ProductCard key={product.id} product={product} />
                ))}
            </div>
            <hr />
            <StoryForm onStoryCreated={handleStoryCreated} />
            <StoryList />
        </div>
    );
};

export default Home;

