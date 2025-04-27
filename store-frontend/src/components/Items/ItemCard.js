import React from 'react';
import { Card, Button } from 'react-bootstrap';

const ItemCard = ({ item, onView }) => {
    return (
        <Card style={{ width: '18rem' }} className="mb-3">
        <Card.Body>
            <Card.Title>{item.name}</Card.Title>
            <Card.Text>
            Price: ${item.price}<br />
            Available: {item.item_total - item.item_booked}<br />
            Provider: {item.provider}
            </Card.Text>
            <Button variant="primary" onClick={onView}>
            View Details
            </Button>
        </Card.Body>
        </Card>
    );
};

export default ItemCard;