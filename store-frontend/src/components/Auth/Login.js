import React from 'react';
import { Form, Button } from 'react-bootstrap';

const Login = ({ onSubmit, error }) => {
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(username, password);
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
        {error && <div className="text-danger mb-3">{error}</div>}
        <Button variant="primary" type="submit">
            Login
        </Button>
        </Form>
    );
};

export default Login;