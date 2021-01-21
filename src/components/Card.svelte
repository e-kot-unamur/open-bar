<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  let timer;
</script>

<div
  class="card"
  on:mousedown={() => (timer = Date.now())}
  on:mouseup={() => {
    Date.now() - timer > 500 ? dispatch("longClick") : dispatch("click");
  }}
>
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
    height: min-content;
    background-color: #f5f5f5;
    border-radius: 1rem;
    transition: 200ms;
  }
  .title {
    font-weight: 300;
    font-size: large;
    user-select: none;
    margin-bottom: 0.8rem;
  }
  .body {
    width: 100%;
    min-height: 1rem;
    font-weight: 400;
    text-align: end;
    overflow: hidden;
    word-break: break-all;
    user-select: none;
  }
  .card:active{
    background-color: var(--highlight);
  }
</style>
