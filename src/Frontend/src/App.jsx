import React from 'react';
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link    
} from 'react-router-dom';
import Login from './pages/Login';
import Register from './pages/Register';
import Contacts from './pages/Contacts';

function App() {
    return (
        <Router>
            <div>
                <nav>
                    <ul>
                        <li><Link to="/login">Login</Link></li>
                        <li><Link to="/register">Register</Link></li>
                    </ul>
                </nav>
                <Routes>
                    <Route path="/login" element={<Login />} exact />
                    <Route path="/register" element={<Register />} />
                    <Route path="/contacts" element={<Contacts />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
