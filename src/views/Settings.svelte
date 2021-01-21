<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";
  import Separator from "../components/Separator.svelte";
  import store from "../store";

  $: if ($store.price === null) $store.price = 0.5;
  $: console.log($store.price);
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

  <Separator>
    <span slot="title">Mark a debt as paid</span>
    <span slot="body">Long click on the concerned participant</span>
  </Separator>

  {#each $store.debts as { name, bars }}
    <Card on:longClick={() => (bars = 0)}>
      <span slot="title">
        {name}
      </span>
      <span slot="body">
        {(bars * $store.price).toFixed(2)}€
      </span>
    </Card>
  {/each}
</Container>
