<script>
  import Container from "./components/Container.svelte";
  import Switcher from "./components/Switcher.svelte";
  import People from "./views/People.svelte";
  import Settings from "./views/Settings.svelte";
</script>

<main>
  <div class="topbar">
    <div class="title">OpenBar</div>
  </div>

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
    background-color: var(--primary);
    overflow: hidden;
  }

  .topbar {
    padding: 4rem;
    text-align: center;
    font-size: xx-large;
    font-weight: 300;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
