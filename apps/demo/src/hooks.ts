import { useEffect, useState } from 'react';
import { useClientContext } from './ApteProvider';

export const useStream = (name: string, handlers: Record<string, (data: string) => void | Promise<void>> = {}) => {
    const client = useClientContext();
    const [error, setError] = useState<Error | null>(null);
    if (!client) return error;

    useEffect(() => {
        const namespace = client.namespace(name);
        namespace.source.addEventListener('error', () => {
            namespace.cleanup();
            setError(new Error(`An error occurred with stream "${name}"`));
        });

        // Apply all handler
        for (const type in handlers) {
            namespace.receive(type, handlers[type]);
        }
        // setDispatch(namespace.dispatch);
        return () => {
            namespace.cleanup();
        };
    }, [client]);

    return error;
};

export const useDispatch = (name: string) => {
    const client = useClientContext();
    if (!client) return;
    
};
