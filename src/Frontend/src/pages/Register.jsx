import React from 'react';

const Register = () => {

  const handleSubmit = (event) => {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    fetch('/api/submit-register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: username,
        password: password
      }),
    })
  }


  return (
    <div class="container">
      <h2>Register</h2>
      <form onSubmit={handleSubmit}>
        <div class="form-group">
          <input type="text" id="username" placeholder="Username" />
          <input type="password" id="password" placeholder="Password" />
          <button type="submit">Register</button>
        </div>
      </form>
    </div>
  );
};

export default Register;