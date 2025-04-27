import React from 'react';
import { Button, Navbar, Container } from 'react-bootstrap';
import { logout } from '../../api/auth';
import { useNavigate } from 'react-router-dom';

const Header = () => {
    const navigate = useNavigate();
    const user = JSON.parse(localStorage.getItem('user'));

    const handleLogout = () => {
        logout();
        navigate('/login');
    };

    return (
        <Navbar bg="dark" variant="dark" expand="lg">
        <Container>
            <Navbar.Brand href="/">Store Management</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse className="justify-content-end">
            {user && (
                <>
                <Navbar.Text className="me-3">
                    Signed in as: <strong>{user.role}</strong>
                </Navbar.Text>
                <Button variant="outline-light" onClick={handleLogout}>
                    Logout
                </Button>
                </>
            )}
            </Navbar.Collapse>
        </Container>
        </Navbar>
    );
};

export default Header;