import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { register, login } from '../api/auth';
import { Button, Form, Card, Alert } from 'react-bootstrap';

const AuthPage = () => {
    const [isLogin, setIsLogin] = useState(true); // По умолчанию показываем логин
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('seller');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
        if (!isLogin) {
            // Регистрация
            await register(username, password, role);
            setError('');
            setIsLogin(true); // Переключаем на логин после успешной регистрации
        } else {
            // Логин
            await login(username, password);
            navigate('/dashboard');
        }
        } catch (err) {
        setError(err.response?.data?.message || 'Wrong password');
        }
    };

    return (
        <div className="d-flex justify-content-center align-items-center vh-100">
            <Card style={{ width: '400px' }}>
                <Card.Body>
                <Card.Title>{isLogin ? 'Login' : 'Register'}</Card.Title>
                {error && <Alert variant="danger" onClose={() => setError('')} dismissible>{error}</Alert>}
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
                    {!isLogin && (
                    <Form.Group className="mb-3">
                        <Form.Label>Role</Form.Label>
                        <Form.Select
                        value={role}
                        onChange={(e) => setRole(e.target.value)}
                        >
                        <option value="seller">Seller</option>
                        <option value="manager">Manager</option>
                        </Form.Select>
                    </Form.Group>
                    )}
                    <Button variant="primary" type="submit" className="w-100 mb-3">
                    {isLogin ? 'Login' : 'Register'}
                    </Button>
                    
                    <div className="text-center">
                    {isLogin ? (
                        <p>
                        Don't have an account?{' '}
                        <Button 
                            variant="link" 
                            onClick={() => {
                            setIsLogin(false);
                            setError('');
                            }}
                            className="p-0"
                        >
                            Register
                        </Button>
                        </p>
                    ) : (
                        <p>
                        Already have an account?{' '}
                        <Button 
                            variant="link" 
                            onClick={() => {
                            setIsLogin(true);
                            setError('');
                            }}
                            className="p-0"
                        >
                            Login
                        </Button>
                        </p>
                    )}
                    </div>
                </Form>
                </Card.Body>
            </Card>
        </div>
    );
};

export default AuthPage;