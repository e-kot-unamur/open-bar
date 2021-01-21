import { writable } from "svelte/store";

const store = writable(new WebSocket('ws://localhost:8000/ws'));

store.subscribe(((ws) => {
    // TODO handle connection state errors
    ws.onerror = () => console.error("Connection with server lost")
    ws.onclose = () => console.error("Connection with server lost")
}))

export default store;