import { writable } from "svelte/store";
import ws from "./websocket.js";

// Link with backend
const store = writable({
  price: undefined,
  users: undefined,
  history: undefined,
});

ws.subscribe((socket) => {

  socket.onmessage = (e) => {
    const message = JSON.parse(e.data);

    switch (message.type) {
      case "allData":
        store.set(message.data);
        break;
      case "newUser":
        store.update((value) => {
          value.users = [...(value.users ?? []), message.user];
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
        handleHistory(message);
        updateDebt(message);
        break;
    }
  }
})

function handleHistory(event) {
  store.update(({ price, users, history }) => {
    let numberOfBars;
    users.map(user =>
      numberOfBars = (user.id === event.id) ? (event.debt - user.debt) : numberOfBars);
    if (typeof numberOfBars === "undefined") return { price, users, history };

    if (history?.length > 0) {
      const [lastHistory] = history.slice(-1);

      const then = new Date(lastHistory.date);
      const end = new Date(then.getTime() + 2 * 60000); // then + 2 minutes
      const now = new Date();
      const inTimeSpan = now <= end && now >= then; //TODO

      if (lastHistory.targetId === event.id && inTimeSpan) {
        lastHistory.numberOfBars += numberOfBars;
        (lastHistory.numberOfBars === 0) && history.pop();
        return { price, users, history }
      }
    }

    history = [
      ...(history ?? []),
      { date: Date.now(), targetId: event.id, numberOfBars }
    ];
    return { price, users, history }
  })
}

function updateDebt(event) {
  store.update((value) => {
    value.users?.map(user => {
      if (user.id === event.id) user.debt = event.debt;
    });
    return value;
  })
}

export default {
  subscribe: store.subscribe,
  set: store.set,
  addParticipant: (name) => ws.send({ type: "newUser", name }),
  updatePrice: (price) => ws.send({ type: "updatePrice", price }),
  updateDebt: (id, int) => ws.send({ type: "updateDebt", id, debt: int }),
  reset: (keepParticipants) => ws.send({ type: "reset", keepParticipants }),
};