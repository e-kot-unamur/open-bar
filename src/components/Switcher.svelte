<script>
  export let views;

  let view = views[Object.keys(views)[0]];
</script>

{#if views}
  <div class="selector">
    {#each Object.entries(views) as [key, value]}
      <div class="view" on:click={() => (view = value)}>{key}</div>
    {/each}
  </div>
{/if}

<div class="slot">
  <svelte:component this={view} />
</div>

<style>
  .selector {
    width: 100%;
    display: flex;
    justify-content: space-around;
    margin-bottom: -1rem;
    z-index: 1;
    position: relative;
  }

  .view {
    width: 100%;
    height: 3rem;
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

  .slot {
    z-index: 20;
    position: relative;
  }
</style>
