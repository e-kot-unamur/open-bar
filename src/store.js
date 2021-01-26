import { writable } from "svelte/store";
import ws from "./websocket.js";

// Link with backend
const store = writable({
  price: undefined,
  users: undefined,
});

ws.subscribe((socket) => {

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

export default {
  subscribe: store.subscribe,
  set: store.set,
  addParticipant: (name) => ws.send({ type: "newUser", name }),
  updatePrice: (price) => ws.send({ type: "updatePrice", price }),
  updateDebt: (id, int) => ws.send({ type: "updateDebt", id, debt: int }),
};