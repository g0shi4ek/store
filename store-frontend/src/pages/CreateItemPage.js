import React from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Button, Alert } from 'react-bootstrap';
import ItemForm from '../components/Items/ItemForm';
import { createItem } from '../api/items';

const CreateItemPage = () => {
    const navigate = useNavigate();
    const [error, setError] = React.useState('');

    const handleSubmit = async (itemData) => {
        try {
            await createItem(itemData);
            navigate('/'); // Перенаправляем после успешного создания
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to create item');
        }
    };

    return (
        <Container className="my-4">
            {error && <Alert variant="danger">{error}</Alert>}
            <div className="d-flex justify-content-between align-items-center mb-4">
                <h2>Create New Item</h2>
                <Button variant="secondary" onClick={() => navigate('/')}>
                    Back to Items
                </Button>
            </div>
            <ItemForm 
                onSubmit={handleSubmit}
                onCancel={() => navigate('/')}
            />
        </Container>
    );
};

export default CreateItemPage;