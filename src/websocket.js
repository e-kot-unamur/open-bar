import { writable } from "svelte/store";

function socket() {

  const host = window.location.href.replace('http', 'ws');
  const socket = new WebSocket(`${host}ws`)
  const { subscribe, set, update } = writable(socket);

  const send = (payload) =>
    socket.readyState && socket.send(JSON.stringify(payload));

  return {
    subscribe,
    set,
    update,
    send,
  }
}


export default socket()