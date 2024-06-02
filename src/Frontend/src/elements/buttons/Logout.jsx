import React from 'react';

const Logout = () => {

    const handleLogout = async () => {

        // Call the logout endpoint
        await fetch('/logout', {
            method: 'GET',
            credentials: 'include'
        });

        // Clear local storage and session storage
        localStorage.clear();
        sessionStorage.clear();

        // Go to root page
        window.location.href = '/';
    };

    return (
        <button onClick={handleLogout}>Logout</button>
    );
};

export default Logout;
