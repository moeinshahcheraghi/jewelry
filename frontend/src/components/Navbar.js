import React, { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext'; // Corrected import

const Navbar = () => {
    const { user, logout } = useContext(AuthContext);

    return (
        <nav>
            <h1>My Jewelry Store</h1>
            {user ? (
                <div>
                    <span>Welcome, {user.name}</span>
                    <button onClick={logout}>Logout</button>
                </div>
            ) : (
                <button>Login</button>
            )}
        </nav>
    );
};

export default Navbar;

