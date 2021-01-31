<script>
  import Card from "../../components/Card.svelte";
  import Separator from "../../components/Separator.svelte";
  import store from "../../store";

  function updatePrice(float) {
    store.updatePrice(parseFloat(($store.price + float).toFixed(2)));
  }
</script>

<Separator>
  <span slot="title">Settings</span>
</Separator>

<Card>
  <span slot="title"> Price </span>
  <span slot="body">
    <button on:click={() => updatePrice(-0.01)}>-</button>
    {$store.price.toFixed(2)}€
    <button on:click={() => updatePrice(0.01)}>+</button>
  </span>
</Card>

<style>
  button {
    width: 5rem;
    margin: 10px;
  }
</style>
