import { useContext, createContext } from 'react';
import type { ApteClient } from 'apte-client';
import type { ReactNode } from 'react';

const clientContext = createContext<ApteClient | null>(null);
export const useClientContext = () => useContext(clientContext);

const ApteProvider = ({ children, client }: { children: ReactNode; client: ApteClient }) => {
    return <clientContext.Provider value={client}>{children && children}</clientContext.Provider>;
};

export default ApteProvider;
