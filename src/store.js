import { writable } from "svelte/store";

// Link with backend
export default writable({
    price: 0.33,
    debts: [
        { name: "Hugo", bars: 10 },
        { name: "Piras", bars: 3 },
        { name: "Baetsle", bars: 3 },
        { name: "Pierre", bars: 3 },
        { name: "Quentain", bars: 100 },
        { name: "Basile", bars: 3 },
        { name: "Evan", bars: 3 },
        { name: "Thibaut", bars: 3 },
    ],
});
