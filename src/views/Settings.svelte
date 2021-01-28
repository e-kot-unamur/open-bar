<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";
  import Separator from "../components/Separator.svelte";
  import store from "../store";

  function updatePrice(float) {
    store.updatePrice(parseFloat(($store.price + float).toFixed(2)));
  }
</script>

<style>
  button {
    width: 5rem;
    margin: 10px;
  }
</style>

<Container>
  <Separator>
    <span slot="title">Settings</span>
  </Separator>

  <Card>
    <span slot="title"> Price </span>
    <span slot="body">
      <button on:click={() => updatePrice(-0.01)}>-</button>
      {($store.price).toFixed(2)}€
      <button on:click={() => updatePrice(0.01)}>+</button>
    </span>
  </Card>

  {#if !$store.users}
    <Separator>
      <span slot="title">There is nothing here yet...</span>
      <span slot="body">Start by adding someone first !</span>
    </Separator>
  {:else}
    <Separator>
      <span slot="title">Mark a debt as paid</span>
      <span slot="body">Long click on the concerned participant</span>
    </Separator>
    {#each $store.users as { id, name, debt }}
      <Card on:longclick={() => store.updateDebt(id, 0)}>
        <span slot="title">
          {name}
        </span>
        <span slot="body">
          {(debt * $store.price).toFixed(2)}€
        </span>
      </Card>
    {/each}
  {/if}
</Container>
