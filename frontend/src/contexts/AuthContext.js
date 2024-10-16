import React, { createContext, useContext, useState } from 'react';
import { jwtDecode } from 'jwt-decode'; // Corrected import

// Create AuthContext
export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    const login = (token) => {
        const decoded = jwtDecode(token); // Decode JWT token
        setUser(decoded);
    };

    const logout = () => {
        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

// Custom hook for using AuthContext
export const useAuth = () => {
    return useContext(AuthContext);
};

