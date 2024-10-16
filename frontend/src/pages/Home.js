// src/pages/Home.js
import React, { useEffect, useState } from 'react';
import API from '../services/api';

const Home = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    API.get('/products')
      .then((res) => setProducts(res.data.data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div style={styles.container}>
      <h1>Products</h1>
      <div style={styles.grid}>
        {products.map((product) => (
          <div key={product.ID} style={styles.card}>
            <h3>{product.Name}</h3>
            <p>{product.Description}</p>
            <p>Price: ${product.Price}</p>
            <p>Quantity: {product.Quantity}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

const styles = {
  container: {
    padding: '20px',
  },
  grid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))',
    gap: '20px',
  },
  card: {
    padding: '15px',
    border: '1px solid #ccc',
    borderRadius: '5px',
  },
};

export default Home;

