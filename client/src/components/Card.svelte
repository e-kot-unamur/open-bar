<script>
  import { handleClickDuration as clickduration } from "./clickduration";

  export let important = false;
</script>

<div
  class="card {important && 'important'}"
  use:clickduration
  on:longclick
  on:shortclick
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
    font-weight: 300;
    text-align: end;
    overflow: hidden;
    word-break: break-all;
    user-select: none;
  }
  .card:active {
    background-color: var(--highlight);
  }
  .important:active {
    background-color: var(--important) !important;
  }
</style>
