import React from 'react';
import { Button } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

const NotFound = () => {
    const navigate = useNavigate();

    return (
        <div className="d-flex flex-column align-items-center justify-content-center vh-100">
        <h1>404 - Page Not Found</h1>
        <p>The page you are looking for does not exist.</p>
        <Button variant="primary" onClick={() => navigate('/')}>
            Go to Home
        </Button>
        </div>
    );
};

export default NotFound;