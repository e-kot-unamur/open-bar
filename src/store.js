import { writable } from "svelte/store";
import ws from "./websocket.js";

// Link with backend
const store = writable({
    price: 0.33,
    debts: [
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
ws.subscribe((store) => {
    socket = store;
    socket.onmessage = (e) => {
        const message = JSON.parse(e.data);
        console.debug(e);
        console.debug(message);

        //TODO handle ws events here
    }
})

function send(payload) {
    if (socket.readyState) socket.send(JSON.stringify(payload))
}

export default {
    subscribe: store.subscribe,

    addParticipant: (name) => send({ event: "newUser", name }),
    updatePrice: (price) => send({ event: "updatePrice", price }),
    updateDebt: (id, int) => send({ event: "updateDebt", id, debt: int }),
};