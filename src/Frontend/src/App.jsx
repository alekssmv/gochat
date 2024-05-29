import React from 'react';
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

function App() {
    const location = useLocation();

    const links = [
        { path: '/login/', label: 'Login' },
        { path: '/register/', label: 'Register' }
    ];

    const filteredLinks = links.filter(link => link.path !== location.pathname);

    return (
        <div>
            <div className="nav-links">
                <nav>
                    <ul className="nav-links">
                        {filteredLinks.map(link => (
                            <a key={link.path}>
                                <Link to={link.path}>{link.label}</Link>
                            </a>
                        ))}
                    </ul>
                </nav>
            </div>
            <Routes>
                <Route path="/login/" element={<Login />} exact />
                <Route path="/register/" element={<Register />} />
                <Route path="/contacts/" element={<Contacts />} />
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
