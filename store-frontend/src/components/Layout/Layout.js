import React, { useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import Header from './Header';
import Sidebar from './Sidebar';
import { Container, Row, Col } from 'react-bootstrap';

const Layout = () => {
    const navigate = useNavigate();
    const user = JSON.parse(localStorage.getItem('user'));

    useEffect(() => {
        if (!user) {
        navigate('/register');
        }
    }, [user, navigate]);

    if (!user) {
        return null;
    }

    return (
        <div>
        <Header />
        <Container fluid>
            <Row>
            <Col md={2} className="p-0">
                <Sidebar />
            </Col>
            <Col md={10} className="p-4">
                <Outlet />
            </Col>
            </Row>
        </Container>
        </div>
    );
};

export default Layout;