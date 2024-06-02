import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Register = ({ setMessage, setLoggedIn }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch('/submit-register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      // Set loggedIn to true
      setLoggedIn(true);
      // Redirect to the contacts page
      navigate('/contacts/');
      // Set message to success
      setMessage('Registration successful!');
    } else {
      // Display error message from response
      const data = await response.json();
      setMessage(data.error)
    }
  };

  return (
    <div className="container">
      <h2>Register</h2>
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
          <button type="submit">Register</button>
        </div>
      </form>
    </div>
  );
};

export default Register;
