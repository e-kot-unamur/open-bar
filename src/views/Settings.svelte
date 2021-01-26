<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";
  import Separator from "../components/Separator.svelte";
  import store from "../store";

  $: store.updatePrice($store.price); 
</script>

<Container>
  <Separator>
    <span slot="title">Settings</span>
  </Separator>

  <Card>
    <span slot="title"> Price </span>
    <span slot="body">
      <input
        type="number"
        id="price"
        min="0.01"
        max="1.5"
        step="0.01"
        bind:value={$store.price}
      />
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
