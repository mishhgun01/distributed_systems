import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import './index.css'; // Если у вас есть стили для приложения

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);