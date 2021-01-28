<script>
  import Switcher from "./components/Switcher.svelte";
  import Topbar from "./components/Topbar.svelte";
  import People from "./views/People.svelte";
  import Settings from "./views/Settings.svelte";
  import store from "./store.js";
import History from "./views/History.svelte";

  let loading;

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
  }
  main {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }
</style>
