<div class="container">
  <slot />
</div>

<style>
  .container {
    background-color: #fff;
    width: 100%;
    border-radius: 1rem 1rem 0 0;
    color: black;
    padding-top: 0.2rem;
    margin-bottom: 3rem; /* Keeps last element visible */
    position: relative;
  }
</style>
