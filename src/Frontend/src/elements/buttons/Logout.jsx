import React from 'react';
import { useNavigate } from 'react-router-dom';

const Logout = () => {
    const navigate = useNavigate();

    const handleLogout = async () => {

        // Call the logout API
        await fetch('/logout', {
            method: 'POST',
            credentials: 'include'
        });

        // Clear local storage and session storage
        localStorage.clear();
        sessionStorage.clear();

        // Go to root
        window.location.href = '/';
    };

    return (
        <button onClick={handleLogout}>Logout</button>
    );
};

export default Logout;
