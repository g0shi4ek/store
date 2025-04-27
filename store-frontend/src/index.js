import React from 'react';
import { createRoot } from 'react-dom/client';
import App from './App';
import 'bootstrap/dist/css/bootstrap.min.css';
import './styles.css';
// Получаем корневой элемент
const container = document.getElementById('root');

// Создаём корень приложения
const root = createRoot(container);

// Рендерим приложение
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);