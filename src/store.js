import { writable } from "svelte/store";
import ws from "./websocket.js";

// Link with backend
const store = writable({
  price: undefined,
  users: undefined,
});

let socket;
ws.subscribe((value) => {
  socket = value;

  socket.onerror = socket.onclose = () => {
    console.error("Connection with server lost");
  }

  socket.onmessage = (e) => {
    const message = JSON.parse(e.data);
    console.log(message);

    switch (message.type) {
      case "allData":
        console.log(message.all);
        store.set(message.all);
        break;
      case "newUser":
        store.update((value) => {
          value.users = [...value.users??[], message.user];
          return value;
        })
        break;
      case "updatePrice":
        store.update((value) => {
          value.price = message.price;
          return value;
        })
        break;
      case "updateDebt":
        store.update((value) => {
          value.users.map(user => {
            if (user.id === message.id) user.debt = message.debt;
          })
          return value;
        })
        break;
    }
  }
})

function send(payload) {
  if (socket.readyState) socket.send(JSON.stringify(payload))
}

export default {
  subscribe: store.subscribe,
  set: store.set,
  addParticipant: (name) => send({ type: "newUser", name }),
  updatePrice: (price) => send({ type: "updatePrice", price }),
  updateDebt: (id, int) => send({ type: "updateDebt", id, debt: int }),
};