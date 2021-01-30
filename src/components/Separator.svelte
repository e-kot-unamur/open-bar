<div class="container">
  <div class="title"><slot name="title" /></div>
  <div class="body"><slot name="body" /></div>
</div>

<style>
  .container {
    margin: 1rem;
  }
  .title {
    font-weight: 200;
    font-size: x-large;
  }
  .body {
    font-weight: 200;
    font-size: small;
  }
</style>
