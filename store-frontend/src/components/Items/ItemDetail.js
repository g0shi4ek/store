import React, { useState } from 'react';
import { Card, Form, Button, Alert, Table, Badge } from 'react-bootstrap';
import ItemForm from './ItemForm'

const ItemDetail = ({ 
    item, 
    onBookItem,
    onBuyItem,
    onCancelItemBooking,
    itemBookings,
    bookingAmount,
    onBookingAmountChange,
    onEditItem,
    onCreateItem,
    userRole 
}) => {
    const [error, setError] = useState('');
    const [isEditing, setIsEditing] = useState(false);
    const [isCreating, setIsCreating] = useState(false);

    const handleBookSubmit = (e) => {
        e.preventDefault();
        if (bookingAmount > item.item_total - item.item_booked) {
            setError('Not enough items available');
            return;
        }
        onBookItem();
        setError('');
    };

    const getStatusBadge = (isActive) => {
        return isActive ? (
            <Badge bg="success">Active</Badge>
        ) : (
            <Badge bg="secondary">Completed</Badge>
        );
    };

    if (isCreating) {
        return (
            <Card>
                <Card.Body>
                    <Card.Title>Create New Item</Card.Title>
                    <ItemForm 
                        item={null}
                        onSubmit={(data) => {
                            onCreateItem(data);
                            setIsCreating(false);
                        }}
                        onCancel={() => setIsCreating(false)}
                    />
                </Card.Body>
            </Card>
        );
    }

    if (isEditing) {
        return (
            <Card>
                <Card.Body>
                    <Card.Title>Edit Item</Card.Title>
                    <ItemForm 
                        item={item}
                        onSubmit={(data) => {
                            onEditItem(data.id, data);
                            setIsEditing(false);
                        }}
                        onCancel={() => setIsEditing(false)}
                    />
                </Card.Body>
            </Card>
        );
    }

    return (
        <Card>
        <Card.Body>
            <Card.Title>{item.name}</Card.Title>
            <Card.Text>
            <strong>Price:</strong> ${item.price}<br />
            <strong>Total:</strong> {item.item_total}<br />
            <strong>Booked:</strong> {item.item_booked}<br />
            <strong>Available:</strong> {item.item_total - item.item_booked}<br />
            <strong>Provider:</strong> {item.provider}
            </Card.Text>

            {userRole === 'manager' && (
                <div className="d-flex justify-content-end mb-3">
                    <Button 
                        variant="warning" 
                        onClick={() => setIsEditing(true)}
                        className="me-2"
                    >
                        Edit Item
                    </Button>
                    <Button 
                        variant="success" 
                        onClick={() => setIsCreating(true)}
                    >
                        Add New Item
                    </Button>
                </div>
            )}

            {userRole === 'seller' && (
            <>
                {error && <Alert variant="danger">{error}</Alert>}
                
                <Form onSubmit={handleBookSubmit} className="mb-4">
                <Form.Group controlId="formAmount" className="mb-3">
                    <Form.Label>Amount to book</Form.Label>
                    <Form.Control
                    type="number"
                    min="1"
                    max={item.item_total - item.item_booked}
                    value={bookingAmount}
                    onChange={(e) => onBookingAmountChange(Number(e.target.value))}
                    />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Create Item Booking
                </Button>
                </Form>

                <h4 className="mb-3">Item Bookings</h4>
                {itemBookings.length > 0 ? (
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
                    {itemBookings.map(booking => (
                        <tr key={booking.id}>
                        <td>{booking.id}</td>
                        <td>{booking.amount}</td>
                        <td>{getStatusBadge(booking.is_active)}</td>
                        <td className="d-flex gap-2">
                            <Button 
                                variant="success" 
                                size="sm"
                                onClick={() => onBuyItem(booking.id)}
                                disabled={booking.is_active}
                            >
                            Confirm Purchase
                            </Button>
                            <Button 
                                variant="danger" 
                                size="sm"
                                onClick={() => onCancelItemBooking(booking.id)}
                            >
                            Cancel Booking
                            </Button>
                        </td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
                ) : (
                <Alert variant="info">No item bookings found</Alert>
                )}
            </>
            )}
        </Card.Body>
        </Card>
    );
};

export default ItemDetail;