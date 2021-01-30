<script>
  import Switcher from "./components/Switcher.svelte";
  import Topbar from "./components/Topbar.svelte";
  import People from "./views/people/People.svelte";
  import Settings from "./views/settings/Settings.svelte";
  import History from "./views/history/History.svelte";
  import store from "./store.js";

  let loading;

  // FIXME : do not use the price as a ref...
  // This induce a dependance from front to back based on go default type...
  // Instead We could watch the websocket connection state
  $: loading = $store.price === undefined;
</script>

<main>
  <Topbar {loading} />
  <Switcher views={{ People, History, Settings }} />
</main>

<style>
  :global(body) {
    --primary: #272932;
    --highlight: #1de068;
    color: #fff;
    display: flex;
    justify-content: center;
    background-color: var(--primary);
  }
  main {
    height: 100%;
    width: 100%;
    max-width: 50rem;
    background-color: #fff;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }
</style>
