import React from 'react';
import { Nav } from 'react-bootstrap';
import { Link, useLocation } from 'react-router-dom';

const Sidebar = () => {
    const location = useLocation();
    const user = JSON.parse(localStorage.getItem('user'));

    return (
        <Nav className="flex-column bg-light h-100 p-3">
        <Nav.Link 
            as={Link} 
            to="/" 
            active={location.pathname === '/'}
            className="mb-2"
        >
            Home
        </Nav.Link>
        <Nav.Link 
            as={Link} 
            to="/dashboard" 
            active={location.pathname === '/dashboard'}
            className="mb-2"
        >
            Dashboard
        </Nav.Link>
        {user?.role === 'manager' && (
            <>
            <Nav.Link disabled className="mt-3 mb-1 fw-bold">Manager</Nav.Link>
            <Nav.Link as={Link} to="/dashboard?tab=room-bookings">
                Room Bookings
            </Nav.Link>

            <Nav.Link as={Link} to="/items/new">
                Add New Item
            </Nav.Link>
            </>

        )}
        {user?.role === 'seller' && (
            <>
            <Nav.Link disabled className="mt-3 mb-1 fw-bold">Seller</Nav.Link>
            <Nav.Link as={Link} to="/dashboard?tab=item-bookings">
                Item Bookings
            </Nav.Link>
            </>
        )}
        </Nav>
    );
};

export default Sidebar;