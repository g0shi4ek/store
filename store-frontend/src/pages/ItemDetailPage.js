import React, { useEffect, useState, useCallback } from 'react';
import { useParams } from 'react-router-dom';
import { getItem, updateItem, createItem } from '../api/items';
import { getStore } from '../api/stores';
import { 
    bookItems, 
    buyItemBooking, 
    cancelItemBooking,
    getAllItemBookings 
} from '../api/bookings';
import { 
    bookRooms, 
    buyRoomBooking,
    cancelRoomBooking,
    getAllRoomBookings 
} from '../api/bookings';
import ItemDetail from '../components/Items/ItemDetail';
import StoreInfo from '../components/Stores/StoreInfo';
import { Container, Row, Col, Alert, Spinner } from 'react-bootstrap';

const ItemDetailPage = () => {
    const { id } = useParams();
    const [item, setItem] = useState(null);
    const [store, setStore] = useState(null);
    const [itemBookings, setItemBookings] = useState([]);
    const [roomBookings, setRoomBookings] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const [itemBookingAmount, setItemBookingAmount] = useState(1);
    const [roomBookingAmount, setRoomBookingAmount] = useState(1);
    const user = JSON.parse(localStorage.getItem('user')) || {};

    const fetchAllData = useCallback(async () => {
        try {
        setLoading(true);
        setError('');
        
        const [itemResponse, storeResponse] = await Promise.all([
            getItem(id),
            getStore(1) // Assuming store ID is 1
        ]);
        
        setItem(itemResponse.data);
        setStore(storeResponse.data);

        if (user.role === 'seller') {
            try {
                const bookingsResponse = await getAllItemBookings();
                console.log("Все бронирования:", bookingsResponse.data);
                setItemBookings(bookingsResponse.data.filter(b => b.item_id === parseInt(id, 10)));
            } catch (err) {
                console.error('Failed to fetch item bookings:', err);
                setItemBookings([]);
            }
        }

        if (user.role === 'manager') {
            try {
                const bookingsResponse = await getAllRoomBookings();
                console.log("Все бронирования:", bookingsResponse.data);
                setRoomBookings(bookingsResponse.data.filter(b => 
                    b.item_id === parseInt(id, 10) && 
                    b.store_id === 1
                ));
            } catch (err) {
                console.error('Failed to fetch room bookings:', err);
                setRoomBookings([]);
            }
        }
        } catch (err) {
            setError(err.response?.data?.message || 'Error loading data');
            console.error('Fetch error:', err);
        } finally {
            setLoading(false);
        }
    }, [id, user.role]);

    useEffect(() => {
        fetchAllData();
    }, [fetchAllData]);

    const handleBookItem = async () => {
        try {
            await bookItems(id, itemBookingAmount);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to create item booking');
        }
    };

    const handleBuyItem = async (bookingId) => {
        try {
            await buyItemBooking(bookingId);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to confirm purchase');
        }
    };

    const handleCancelItemBooking = async (bookingId) => {
        try {
            await cancelItemBooking(bookingId);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to cancel booking');
        }
    };

    const handleBookRoom = async () => {
        try {
            await bookRooms(1, id, roomBookingAmount);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to create room booking');
        }
    };

    const handleBuyRoom = async (bookingId) => {
        try {
            await buyRoomBooking(bookingId);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to activate rooms');
        }
    };

    const handleCancelRoomBooking = async (bookingId) => {
        try {
            await cancelRoomBooking(bookingId);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to cancel booking');
        }
    };

    const handleEditItem = async (id, itemData) => {
        try {
            await updateItem(id, itemData);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to update item');
        }
    };

    const handleCreateItem = async (itemData) => {
        try {
            await createItem(itemData);
            await fetchAllData();
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to create item');
        }
    };


    if (loading) {
        return (
        <Container className="d-flex justify-content-center my-5">
            <Spinner animation="border" role="status">
            <span className="visually-hidden">Loading...</span>
            </Spinner>
        </Container>
        );
    }

    if (!item || !store) {
        return (
        <Container className="my-5">
            <Alert variant="danger">Failed to load item or store data</Alert>
        </Container>
        );
    }

    return (
        <Container className="my-4">
        {error && (
            <Alert variant="danger" onClose={() => setError('')} dismissible>
            {error}
            </Alert>
        )}
        
        <Row>
            <Col md={6}>
            <ItemDetail 
                item={item} 
                onBookItem={handleBookItem}
                onBuyItem={handleBuyItem}
                onCancelItemBooking={handleCancelItemBooking}
                itemBookings={itemBookings}
                bookingAmount={itemBookingAmount}
                onBookingAmountChange={setItemBookingAmount}
                onEditItem={handleEditItem}
                onCreateItem={handleCreateItem}
                userRole={user?.role}
            />
            </Col>
            <Col md={6}>
            <StoreInfo 
                store={store} 
                onBookRoom={handleBookRoom}
                onBuyRoom={handleBuyRoom}
                onCancelRoomBooking={handleCancelRoomBooking}
                roomBookings={roomBookings}
                bookingAmount={roomBookingAmount}
                onBookingAmountChange={setRoomBookingAmount}
                userRole={user?.role}
            />
            </Col>
        </Row>
        </Container>
    );
};

export default ItemDetailPage;