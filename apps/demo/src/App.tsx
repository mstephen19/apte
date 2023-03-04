import { useStream } from './hooks';

function App() {
    const error = useStream('x', {
        msg: (data) => {
            console.log(data);
        },
    });

    return (
        <main>
            <h1>Hello Apte</h1>
            <p>{error && error.message}</p>
        </main>
    );
}

export default App;
