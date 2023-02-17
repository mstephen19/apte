# Apte

This project is currently in its beginning stages of development.

**Apte** will be a framework for both the frontend and backend that will aim to not only make server-side events easy to implement, but also solve a lot of the problems that are currently present with server-side events (such as the uni-directional nature of SSE).

## Concepts to understand

1. **Namespaces** in the context of Apte will be named locations where events can both be streamed, as well as sent and handled.
2. A namespace can have an unlimited number of **Type**s. Types are just string values that are attached to an event to help discern it from other events sent in the namespace.

## Goals

- **Out-of-the-box Redis support** - A certain level of statefulness is required the more complex your use-case of SSE becomes. State can be stored in memory, but that gets messy quickly. Instead, Redis is a much more fantastic option for storing temporary data related to clients, or subscribing to events sent from the client to the server in a certain namespace.
- **Multi-language/framework support** - Initially, the main focus will be to develop, test, and release two backend Apte implementations (Golang + Node w/ TypeScript) along with a Vanilla TS frontend library and a React adaptation.
- **Fantastic type support** - As with [Nanolith](https://github.com/mstephen19/nanolith), the goal with Apte is to have seamless TypeScript support.
- **Intuitiveness** - Apte must be simple to use with as little new concepts to learn as possible. The learning curve should not be steep, and writing code with Apte should feel natural with as little boilerplate as possible.

## Plan of action

1. Golang library
2. Vanilla TS library
3. Node.js library
4. React library
