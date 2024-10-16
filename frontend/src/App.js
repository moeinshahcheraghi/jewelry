// src/App.js
import React, { useContext } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import Admin from './pages/Admin';
import Stories from './pages/Stories';
import { AuthContext } from './contexts/AuthContext';

const App = () => {
  const { auth } = useContext(AuthContext);

  return (
    <div>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route
          path="/login"
          element={!auth.user ? <Login /> : <Navigate to="/" />}
        />
        <Route
          path="/register"
          element={!auth.user ? <Register /> : <Navigate to="/" />}
        />
        <Route
          path="/admin"
          element={
            auth.user && auth.isAdmin ? <Admin /> : <Navigate to="/" />
          }
        />
        <Route
          path="/stories"
          element={auth.user ? <Stories /> : <Navigate to="/login" />}
        />
        {/* مسیرهای دیگر */}
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </div>
  );
};

export default App;

