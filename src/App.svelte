<script>
  import Switcher from "./components/Switcher.svelte";
  import Topbar from "./components/Topbar.svelte";
  import People from "./views/People.svelte";
  import Settings from "./views/Settings.svelte";
</script>

<main>
  <Topbar />
  <Switcher views={{ People, Settings }} />
</main>

<style>
  :global(body) {
    --primary: #272932;
    --highlight: #5398be;
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
