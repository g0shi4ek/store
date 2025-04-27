import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getAllItems } from '../api/items';
import ItemCard from '../components/Items/ItemCard';
import { Row, Col, Container } from 'react-bootstrap';

const Home = () => {
    const [items, setItems] = useState([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchItems = async () => {
        try {
            const response = await getAllItems();
            setItems(response.data);
        } catch (error) {
            console.error('Error fetching items:', error);
        } finally {
            setLoading(false);
        }
        };
        fetchItems();
    }, []);

    if (loading) return <div>Loading...</div>;

    return (
        <Container>
        <h2 className="my-4">Item Catalog</h2>
        <Row xs={1} md={2} lg={3} className="g-4">
            {items.map(item => (
            <Col key={item.id}>
                <ItemCard item={item} onView={() => navigate(`/items/${item.id}`)} />
            </Col>
            ))}
        </Row>
        </Container>
    );
};

export default Home;