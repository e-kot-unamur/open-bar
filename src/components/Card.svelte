<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();
</script>

<div class="card" on:click={() => dispatch("click")}>
  <div class="title">
    <slot name="title" />
  </div>
  <div class="body">
    <slot name="body" />
  </div>
</div>

<style>
  .card {
    padding: 0.5rem;
    margin: 0.3rem 0.3rem 0 0.3rem;
    height: 3rem;
    background-color: #f5f5f5;
    border-radius: 1rem;
  }
</style>
