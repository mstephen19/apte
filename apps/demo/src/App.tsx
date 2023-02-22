import { useEffect } from 'react';
import { ApteClient } from 'apte-client';

const client = new ApteClient({ url: 'http://localhost:3000/events' });

function App() {
    useEffect(() => {
        const namespace = client.namespace('messages');

        namespace.receive('message', (data) => {
            console.log(data);
            namespace.dispatch('message', JSON.stringify({ foo: 'bar' }));
        });

        return () => {
            namespace.cleanup();
        };
    });

    return (
        <main>
            <h1>Hello Apte</h1>
        </main>
    );
}

export default App;
