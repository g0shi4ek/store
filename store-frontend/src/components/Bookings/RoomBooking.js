import React from 'react';
import { Card, Table } from 'react-bootstrap';

const RoomBooking = ({ bookings, onCancel, onBuy }) => {
    return (
        <Card>
        <Card.Body>
            <Card.Title>Room Bookings</Card.Title>
            <Table striped bordered hover>
            <thead>
                <tr>
                <th>ID</th>
                <th>Store ID</th>
                <th>Item ID</th>
                <th>Amount</th>
                <th>Status</th>
                <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {bookings.map(booking => (
                <tr key={booking.id}>
                    <td>{booking.id}</td>
                    <td>{booking.store_id}</td>
                    <td>{booking.item_id}</td>
                    <td>{booking.amount}</td>
                    <td>{booking.is_active ? 'Active' : 'Inactive'}</td>
                    <td>
                    <button 
                        className="btn btn-danger btn-sm me-2"
                        onClick={() => onCancel(booking.id)}
                    >
                        Cancel
                    </button>
                    <button 
                        className="btn btn-success btn-sm"
                        onClick={() => onBuy(booking.id)}
                    >
                        Buy
                    </button>
                    </td>
                </tr>
                ))}
            </tbody>
            </Table>
        </Card.Body>
        </Card>
    );
};

export default RoomBooking;