<script>
  import Card from "../../components/Card.svelte";
  import Separator from "../../components/Separator.svelte";
  import store from "../../store";
</script>

<Separator>
  <span slot="title">Mark a debt as paid</span>
  <span slot="body">
    {#if $store.users?.length}
      Long click on the concerned participant
    {:else}
      Start by adding someone first !
    {/if}
  </span>
</Separator>

{#each $store.users ?? [] as { id, name, debt }}
  <Card on:longclick={() => store.updateDebt(id, 0)}>
    <span slot="title">
      {name}
    </span>
    <span slot="body">
      {(debt * $store.price).toFixed(2)}€
    </span>
  </Card>
{/each}
