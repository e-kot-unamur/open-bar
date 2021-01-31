<script>
  import Card from "../../components/Card.svelte";
  import store from "../../store";
</script>

{#each $store.users ?? [] as { id, name, debt }}
  <Card
    on:shortclick={() => store.updateDebt(id, debt + 1)}
    on:longclick={() => debt > 0 && store.updateDebt(id, debt - 1)}
  >
    <span slot="title">{name} ({(debt * $store.price).toFixed(2)}€)</span>
    <span slot="body">
      {#each Array((debt - (debt % 5)) / 5) as _}
        <span class="tally-marks">IIII</span>
      {/each}
      <span>{"I".repeat(debt % 5)}</span>
    </span>
  </Card>
{/each}

<style>
  .tally-marks {
    text-decoration: line-through;
    padding: 0.2rem;
  }
</style>
