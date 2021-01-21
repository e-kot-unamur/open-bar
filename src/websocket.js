import { writable } from "svelte/store";

export default writable(new WebSocket('ws://localhost:8000/ws'));
