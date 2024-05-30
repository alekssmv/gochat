import React, { useState, useEffect } from 'react';
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link,
    useLocation
} from 'react-router-dom';
import Login from './pages/Login';
import Register from './pages/Register';
import Contacts from './pages/Contacts';
import Logout from './elements/buttons/Logout';
import UsernameDisplay from './elements/buttons/UsernameDisplay';

function App() {

    const location = useLocation();
    const [loggedIn, setLoggedIn] = useState(false);

    useEffect(() => {

        // Check if the cookie is set
        const cookie = document.cookie.split(';').find(c => c.trim().startsWith('gochat-session='));
        if (cookie) {
            setLoggedIn(true);
        }
    }, []);

    const links = [
        { path: '/login/', label: 'Login' },
        { path: '/register/', label: 'Register' }
    ];

    { /* dynamic links */ }
    const filteredLinks = links.filter(link => link.path !== location.pathname);

    return (
        <div>
            <div className="nav-links">
                <nav>
                    <ul className="nav-links">

                        {/* Add the logout button if the user is logged in */}
                        {loggedIn && (
                            <div>
                                <Logout />
                                <UsernameDisplay />
                            </div>
                        )}

                        {/* Add the login and register links if the user is not logged in */}
                        {!loggedIn && (
                            <ul>
                                {filteredLinks.map(link => (
                                    <li key={link.path}>
                                        <Link to={link.path}>{link.label}</Link>
                                    </li>
                                ))}
                            </ul>
                        )}

                    </ul>
                </nav>
            </div>
            <Routes>
                <Route path="/login/" element={<Login />} exact />
                <Route path="/register/" element={<Register />} />
                <Route path="/contacts/" element={<Contacts />} />
                <Route path="/logout" element={<Logout />} />
            </Routes>
        </div>
    );
}

function AppWrapper() {
    return (
        <Router>
            <App />
        </Router>
    );
}

export default AppWrapper;
