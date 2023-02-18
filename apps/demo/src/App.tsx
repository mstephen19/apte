import { useEffect } from 'react';

function App() {
    useEffect(() => {
        const source = new EventSource('http://localhost:3000/events?namespace=messages');

        return () => {
            source.close();
        };
    });

    return (
        <main>
            <h1>Hello Apte</h1>
        </main>
    );
}

export default App;
