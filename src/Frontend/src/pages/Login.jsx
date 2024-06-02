import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = ({ setMessage, setLoggedIn }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        const response = await fetch('/submit-login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            // set loggedIn to true
            setLoggedIn(true);
            // Redirect to the contacts page
            navigate('/contacts/');
            // set message to success
            setMessage('Login successful!');
        } else {
            // Display error message from response
            const data = await response.json();
            setMessage(data.error)
        }
    };

    return (
        <div className="container">
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <input
                        type="text"
                        id="username"
                        placeholder="Username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        type="password"
                        id="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <button type="submit">Login</button>
                </div>
            </form>
        </div>
    );
};

export default Login;
