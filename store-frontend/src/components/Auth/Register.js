import React from 'react';
import { Form, Button } from 'react-bootstrap';

const Register = ({ onSubmit, error }) => {
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [role, setRole] = React.useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(username, password, role);
    };

    return (
        <Form onSubmit={handleSubmit}>
        <Form.Group className="mb-3">
            <Form.Label>Username</Form.Label>
            <Form.Control
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            />
        </Form.Group>
        <Form.Group className="mb-3">
            <Form.Label>Password</Form.Label>
            <Form.Control
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            />
        </Form.Group>
        <Form.Group className="mb-3">
            <Form.Label>Role</Form.Label>
            <Form.Select value={role} onChange={(e) => setRole(e.target.value)}>
            <option value="seller">Seller</option>
            <option value="manager">Manager</option>
            </Form.Select>
        </Form.Group>
        {error && <div className="text-danger mb-3">{error}</div>}
        <Button variant="primary" type="submit">
            Register
        </Button>
        </Form>
    );
};

export default Register;