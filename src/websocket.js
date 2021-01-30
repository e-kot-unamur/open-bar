import { writable } from "svelte/store";

export const readyState = writable(false);

function store() {

  const host = window.location.href.replace('http', 'ws');
  const { subscribe, set, update } = writable(new WebSocket(`${host}ws`));

  let socket;

  subscribe((value) => {
    socket = value;

    socket.onclose = () => {
      console.warn("Connection with server lost. Reconnecting...");
      readyState.set(false);
      set(new WebSocket(`${host}ws`));
    }

    socket.onerror = () => {
      readyState.set(false);
      socket.close();
    }

    socket.onopen = () => {
      console.log("Connection with server established.");
      readyState.set(true);
    }

  });

  const send = (payload) =>
    socket.readyState && socket.send(JSON.stringify(payload));

  return {
    subscribe,
    set,
    update,
    send,
  }
}


export default store()