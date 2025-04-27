import React from 'react';
import { Form, Button } from 'react-bootstrap';

const ItemForm = ({ item, onSubmit, onCancel }) => {
    const [formData, setFormData] = React.useState(item || {
        name: '',
        item_total: 0,
        provider: '',
        price: 0
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({
        ...prev,
        [name]: name === 'price' || name === 'item_total' ? Number(value) : value
        }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(formData);
    };

    return (
        <Form onSubmit={handleSubmit}>
        <Form.Group className="mb-3">
            <Form.Label>Name</Form.Label>
            <Form.Control
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                required
            />
        </Form.Group>
        <Form.Group className="mb-3">
            <Form.Label>Total Quantity</Form.Label>
            <Form.Control
                type="number"
                name="item_total"
                min="0"
                value={formData.item_total}
                onChange={handleChange}
                required
            />
        </Form.Group>
        <Form.Group className="mb-3">
            <Form.Label>Provider</Form.Label>
            <Form.Control
                type="text"
                name="provider"
                value={formData.provider}
                onChange={handleChange}
                required
            />
        </Form.Group>
        <Form.Group className="mb-3">
            <Form.Label>Price</Form.Label>
            <Form.Control
                type="number"
                name="price"
                min="1"
                value={formData.price}
                onChange={handleChange}
                required
            />
        </Form.Group>
        <div className="d-flex justify-content-end">
            <Button variant="secondary" onClick={onCancel} className="me-2">
            Cancel
            </Button>
            <Button variant="primary" type="submit">
                {item ? 'Update' : 'Create'}
            </Button>
        </div>
        </Form>
    );
};

export default ItemForm;