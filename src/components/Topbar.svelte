<script>
  import { tick } from "svelte";
  export let loading = false;

  let component;

  async function minimize() {
    await tick();
    component.style.height = "10vh";
    component.style.lineHeight = "10vh";
    component.style.zIndex = "0";
    component.style.fontSize = "xx-large";
  }

  // TODO : use this when we lose ws connection
  async function maximize() {
    await tick();
    component.style.height = "100vh";
    component.style.lineHeight = "100vh";
    component.style.zIndex = "100";
    component.style.fontSize = "10vw";
  }

  // This is used as a loading screen
  $: !loading && setTimeout(minimize, 1000);
</script>

<div class="topbar" bind:this={component}>
  <div class="title">OpenBar</div>
  {#if loading}
    <div class="loader">Loading...</div>
  {/if}
</div>

<style>
  .topbar {
    background-color: var(--primary);

    z-index: 100;
    width: 100%;
    height: 100vh;
    line-height: 100vh;

    text-align: center;

    display: flex;
    flex-direction: column;
    justify-content: center;

    transition: 500ms;
  }
  .title {
    font-size: 10vw;
    font-weight: 200;
  }
  .loader {
    position: absolute;
    bottom: 1rem;
    width: 100%;
    line-height: initial;
    font-size: 5vw;
    font-weight: 100;
  }
</style>
