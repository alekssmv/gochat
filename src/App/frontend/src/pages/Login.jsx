import React from 'react';

const Home = () => {
    return (
        <div class="container">
            <h2>Login</h2>
            <form>
                <div class="form-group">
                    <input type="text" id="username" placeholder="Username" />
                    <input type="password" id="password" placeholder="Password" />
                    <button type="submit">Login</button>
                </div>
            </form>
        </div>
    );
};

export default Home;