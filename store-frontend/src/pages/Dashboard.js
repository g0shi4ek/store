import React, { useEffect, useState } from 'react';
import { getAllItems } from '../api/items';
import { getAllStores } from '../api/stores';
import { getAllRoomBookings, getAllItemBookings } from '../api/bookings';
import { Card, Container, Tab, Tabs, Table, Alert, Spinner } from 'react-bootstrap';

const Dashboard = () => {
    const [activeTab, setActiveTab] = useState('items');
    const [items, setItems] = useState([]);
    const [stores, setStores] = useState([]);
    const [roomBookings, setRoomBookings] = useState([]);
    const [itemBookings, setItemBookings] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const user = JSON.parse(localStorage.getItem('user')) || {};

    useEffect(() => {
        const fetchData = async () => {
        try {
            setLoading(true);
            setError('');
            
            const [itemsResponse, storesResponse] = await Promise.all([
                getAllItems().catch(e => ({ data: [] })),
                getAllStores().catch(e => ({ data: [] }))
            ]);

            setItems(itemsResponse?.data || []);
            setStores(storesResponse?.data || []);

            if (user?.role === 'manager') {
                const bookingsResponse = await getAllRoomBookings().catch(e => ({ data: [] }));
                setRoomBookings(bookingsResponse?.data || []);
            }

            if (user?.role === 'seller') {
                const bookingsResponse = await getAllItemBookings().catch(e => ({ data: [] }));
                setItemBookings(bookingsResponse?.data || []);
            }
        } catch (err) {
            setError(err.response?.data?.message || 'Error fetching data');
            console.error('Fetch error:', err);
        } finally {
            setLoading(false);
        }
        };
        
        fetchData();
    }, [user?.role]);

    if (loading) {
        return (
        <Container className="d-flex justify-content-center my-5">
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        </Container>
        );
    }

    if (error) {
        return (
        <Container className="my-4">
            <Alert variant="danger">{error}</Alert>
        </Container>
        );
    }

    return (
        <Container className="my-4">
        <h2>Dashboard</h2>
        
        <Tabs activeKey={activeTab} onSelect={setActiveTab} className="mb-3">
            <Tab eventKey="items" title="Items">
            <h3>Items</h3>
            {items.length > 0 ? (
                <div className="row">
                {items.map(item => (
                    <div key={item.id} className="col-md-4 mb-3">
                    <Card>
                        <Card.Body>
                        <Card.Title>{item.name}</Card.Title>
                        <Card.Text>
                            Price: ${item.price}<br />
                            Available: {item.item_total - item.item_booked}<br />
                            Booked: {item.item_booked}
                        </Card.Text>
                        </Card.Body>
                    </Card>
                    </div>
                ))}
                </div>
            ) : (
                <Alert variant="info">No items found</Alert>
            )}
            </Tab>
            
            <Tab eventKey="stores" title="Stores">
            <h3>Stores</h3>
            {stores.length > 0 ? (
                <div className="row">
                {stores.map(store => (
                    <div key={store.id} className="col-md-6 mb-3">
                    <Card>
                        <Card.Body>
                        <Card.Title>Store #{store.number}</Card.Title>
                        <Card.Text>
                            Total Rooms: {store.room_total}<br />
                            Booked Rooms: {store.room_booked}<br />
                            Occupied Rooms: {store.room_occupied}
                        </Card.Text>
                        </Card.Body>
                    </Card>
                    </div>
                ))}
                </div>
            ) : (
                <Alert variant="info">No stores found</Alert>
            )}
            </Tab>
            
            {user?.role === 'manager' && (
            <Tab eventKey="room-bookings" title="Room Bookings">
                <h3>Room Bookings</h3>
                {roomBookings.length > 0 ? (
                <Table striped bordered hover>
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Amount</th>
                        <th>Store</th>
                        <th>Item</th>
                        <th>Status</th>
                    </tr>
                    </thead>
                    <tbody>
                    {roomBookings.map(booking => (
                        <tr key={booking.id}>
                        <td>{booking.id}</td>
                        <td>{booking.amount}</td>
                        <td>{booking.store_id}</td>
                        <td>{booking.item_id}</td>
                        <td>{booking.is_active ? 'Active' : 'Inactive'}</td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
                ) : (
                <Alert variant="info">No room bookings found</Alert>
                )}
            </Tab>
            )}
            
            {user?.role === 'seller' && (
            <Tab eventKey="item-bookings" title="Item Bookings">
                <h3>Item Bookings</h3>
                {itemBookings.length > 0 ? (
                <Table striped bordered hover>
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Amount</th>
                        <th>Item</th>
                        <th>Status</th>
                    </tr>
                    </thead>
                    <tbody>
                    {itemBookings.map(booking => (
                        <tr key={booking.id}>
                        <td>{booking.id}</td>
                        <td>{booking.amount}</td>
                        <td>{booking.item_id}</td>
                        <td>{booking.is_active ? 'Active' : 'Inactive'}</td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
                ) : (
                <Alert variant="info">No item bookings found</Alert>
                )}
            </Tab>
            )}
        </Tabs>
        </Container>
    );
};

export default Dashboard;