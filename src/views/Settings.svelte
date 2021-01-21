<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";
  import Separator from "../components/Separator.svelte";
  import store from "../store";

  let price = $store.price;

  $: if (price === null) {
    price = $store.price;
  } else {
    store.updatePrice($store.price);
  }
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
        bind:value={price}
      />
    </span>
  </Card>

  <Separator>
    <span slot="title">Mark a debt as paid</span>
    <span slot="body">Long click on the concerned participant</span>
  </Separator>

  {#each $store.debts as { id, name, debt }}
    <Card on:longClick={() => store.updateDebt(id, 0)}>
      <span slot="title">
        {name}
      </span>
      <span slot="body">
        {(debt * $store.price).toFixed(2)}€
      </span>
    </Card>
  {/each}
</Container>
