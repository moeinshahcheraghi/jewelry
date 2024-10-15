// frontend/src/components/Navbar.js
import React from 'react';
import { Link, useHistory } from 'react-router-dom';

const Navbar = () => {
    const history = useHistory();

    const logout = () => {
        localStorage.removeItem('token');
        history.push('/login');
    };

    const isAuthenticated = !!localStorage.getItem('token');

    return (
        <nav>
            <Link to="/">Home</Link>
            {isAuthenticated ? (
                <>
                    <Link to="/admin">Admin</Link>
                    <button onClick={logout}>Logout</button>
                </>
            ) : (
                <>
                    <Link to="/login">Login</Link>
                    <Link to="/register">Register</Link>
                </>
            )}
        </nav>
    );
};

export default Navbar;

