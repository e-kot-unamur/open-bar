import { writable } from "svelte/store";
import ws from "./websocket.js";

// Link with backend
const store = writable({
  price: 0.33,
  users: [
    { id: 0, name: "Hugo", debt: 10 },
    { id: 1, name: "Piras", debt: 3 },
    { id: 2, name: "Baetsle", debt: 3 },
    { id: 3, name: "Pierre", debt: 3 },
    { id: 4, name: "Quentain", debt: 100 },
    { id: 5, name: "Basile", debt: 3 },
    { id: 6, name: "Evan", debt: 3 },
    { id: 7, name: "Thibaut", debt: 3 },
  ],
});


let socket;
ws.subscribe((value) => {
  socket = value;
  socket.onmessage = (e) => {
    const message = JSON.parse(e.data);
    console.debug(message);

    //TODO handle ws types here
    switch (message.type) {
      case "newUser":
        store.update((value) => {
          value.users = [...value.users, message.user];

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
            if (user.id === message.id) user.debt = message.debt
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

  addParticipant: (name) => send({ type: "newUser", name }),
  updatePrice: (price) => send({ type: "updatePrice", price }),
  updateDebt: (id, int) => send({ type: "updateDebt", id, debt: int }),
};