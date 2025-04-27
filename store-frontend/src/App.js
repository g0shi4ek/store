import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import AuthPage from './pages/AuthPage';
import Home from './pages/Home';
import Dashboard from './pages/Dashboard';
import CreateItemPage from './pages/CreateItemPage';
import ItemDetailPage from './pages/ItemDetailPage';
import NotFound from './pages/NotFound';
import Layout from './components/Layout/Layout';

function App() {
    return (
        <Router>
        <Routes>
            <Route path="/login" element={<AuthPage type="login" />} />
            <Route path="/register" element={<AuthPage type="register" />} />
            <Route path="/" element={<Layout />}>
                <Route index element={<Home />} />
                <Route path="dashboard" element={<Dashboard />} />
                <Route path="items/:id" element={<ItemDetailPage />} />
                <Route path="/items/new" element={<CreateItemPage />} />
            </Route>
            <Route path="*" element={<NotFound />} />
        </Routes>
        </Router>
    );
}

export default App;