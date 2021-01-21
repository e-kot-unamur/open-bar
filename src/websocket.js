import { writable } from "svelte/store";

const host = window.location.href.replace('http', 'ws');

export default writable(new WebSocket(`${host}ws`));
