<script>
  import { onMount } from "svelte";

  let component;

  function handleScroll() {
    component.style.height = "10vh";
    component.style.lineHeight = "10vh";
    component.style.zIndex = "0";
    component.style.fontSize = "xx-large";
  }

  onMount(() => {
    setTimeout(handleScroll, 1000);
  });
</script>

<div class="topbar" bind:this={component}>
  <div class="title">OpenBar</div>
</div>

<style>
  .topbar {
    background-color: var(--primary);

    z-index: 100;
    width: 100%;
    height: 100vh;

    transition: 500ms;
    line-height: 100vh;
    text-align: center;
    font-size: 10vw;
    font-weight: 200;
  }
</style>
