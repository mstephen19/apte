type ApteClientOptions = {
    url: string;
};

export class ApteClient {
    #base: string;
    constructor(options: ApteClientOptions) {
        this.#base = options.url;
    }

    namespace(name: string) {
        const url = new URL(this.#base);
        url.searchParams.set('namespace', name);
        return new Namespace(url);
    }
}

class Namespace {
    source: EventSource;
    url: URL;

    constructor(url: URL) {
        this.url = url;
        this.source = new EventSource(new URL(url));
    }

    receive(type: string, handler: (data: string) => void | Promise<void>) {
        this.source.addEventListener(type, (event) => handler(event.data));
    }

    async dispatch<Data extends string>(type: string, data: Data) {
        const url = new URL(this.url);
        url.searchParams.set('type', type);

        await fetch(url, {
            method: 'POST',
            body: data,
        });
    }

    cleanup() {
        this.source.close();
    }
}
