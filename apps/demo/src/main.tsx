import React from 'react';
import ReactDOM from 'react-dom/client';
import { ApteClient } from 'apte-client';
import App from './App';
import ApteProvider from './ApteProvider';

const client = new ApteClient({ url: 'http://localhost:3000/events' });

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <ApteProvider client={client}>
            <App />
        </ApteProvider>
    </React.StrictMode>
);
