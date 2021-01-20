<script>
  export let views;

  let view = views[Object.keys(views)[0]];
</script>

<div class="container">
  {#if views}
    <div class="selector">
      {#each Object.entries(views) as [key, value]}
        <div
          class="view {value === view && 'active'}"
          on:click={() => (view = value)}
        >{key}</div>
      {/each}
    </div>
  {/if}

  <div class="slot">
    <svelte:component this={view} />
  </div>
</div>

<style>
  .container {
    height: 90vh;
    width: 100%;
    z-index: 1;
  }
  .selector {
    width: 100%;
    height: 4rem;
    display: flex;
    justify-content: space-around;
    margin-bottom: -1rem;
    z-index: 2;
    position: relative;
    background-color: var(--primary);
  }
  .view {
    width: 100%;
    height: auto;

    padding-top: 0.8rem;
    border-radius: 0.2rem;
    text-align: center;
    cursor: pointer;
    transition-duration: 150ms;
    transition-delay: 10ms;
  }
  .view:active {
    background-color: var(--highlight);
  }
  .active {
    color: var(--highlight);
  }

  .slot {
    height: calc(100% - 4rem);
    z-index: 3;
    position: relative;
    overflow-y: auto;
  }
</style>
