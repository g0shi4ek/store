import React, { useState } from 'react';
import { Card, Form, Button, Alert, Table, Badge } from 'react-bootstrap';

const StoreInfo = ({ 
    store, 
    onBookRoom,
    onBuyRoom,
    onCancelRoomBooking,
    roomBookings,
    bookingAmount,
    onBookingAmountChange,
    userRole 
}) => {
    const [error, setError] = useState('');

    const handleBookSubmit = (e) => {
        e.preventDefault();
        if (bookingAmount > store.room_total - store.room_booked - store.room_occupied) {
        setError('Not enough rooms available');
        return;
        }
        onBookRoom();
        setError('');
    };

    const getStatusBadge = (isActive) => {
        return isActive ? (
        <Badge bg="success">Active</Badge>
        ) : (
        <Badge bg="secondary">Occupied</Badge>
        );
    };

    return (
        <Card>
        <Card.Body>
            <Card.Title>Store #{store.number}</Card.Title>
            <Card.Text>
            <strong>Total rooms:</strong> {store.room_total}<br />
            <strong>Booked rooms:</strong> {store.room_booked}<br />
            <strong>Occupied rooms:</strong> {store.room_occupied}<br />
            <strong>Available rooms:</strong> {store.room_total - store.room_booked - store.room_occupied}
            </Card.Text>

            {userRole === 'manager' && (
            <>
                {error && <Alert variant="danger">{error}</Alert>}
                
                <Form onSubmit={handleBookSubmit} className="mb-4">
                <Form.Group controlId="formRoomAmount" className="mb-3">
                    <Form.Label>Rooms to book</Form.Label>
                    <Form.Control
                    type="number"
                    min="1"
                    max={store.room_total - store.room_booked - store.room_occupied}
                    value={bookingAmount}
                    onChange={(e) => onBookingAmountChange(Number(e.target.value))}
                    />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Create Room Booking
                </Button>
                </Form>

                <h4 className="mb-3">Room Bookings</h4>
                {roomBookings.length > 0 ? (
                <Table striped bordered hover responsive>
                    <thead className="table-dark">
                    <tr>
                        <th>ID</th>
                        <th>Amount</th>
                        <th>Status</th>
                        <th>Actions</th>
                    </tr>
                    </thead>
                    <tbody>
                    {roomBookings.map(booking => (
                        <tr key={booking.id}>
                        <td>{booking.id}</td>
                        <td>{booking.amount}</td>
                        <td>{getStatusBadge(booking.is_active)}</td>
                        <td className="d-flex gap-2">
                            <Button 
                            variant="success" 
                            size="sm"
                            onClick={() => onBuyRoom(booking.id)}
                            disabled={booking.is_active}
                            >
                            Activate Rooms
                            </Button>
                            <Button 
                            variant="danger" 
                            size="sm"
                            onClick={() => onCancelRoomBooking(booking.id)}

                            >
                            Cancel Booking
                            </Button>
                        </td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
                ) : (
                <Alert variant="info">No room bookings found</Alert>
                )}
            </>
            )}
        </Card.Body>
        </Card>
    );
};

export default StoreInfo;