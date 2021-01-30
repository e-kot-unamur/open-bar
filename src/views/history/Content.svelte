<script>
  import Card from "../../components/Card.svelte";
  import store from "../../store";
</script>

{#each $store.history
  ?.slice()
  .reverse() ?? [] as { date, targetId, numberOfBars }}
  <Card>
    <!-- FIXME : change users to a Map with the user id as the key ! -->
    <span slot="title">{$store.users[targetId].name}</span>
    <span class="card-body" slot="body">
      <p>{new Date(date).toLocaleString()}</p>
      <p>{numberOfBars}</p>
    </span>
  </Card>
{/each}

<style>
  .card-body {
    width: 100%;
    height: min-content;
    display: flex;
    justify-content: space-between;
  }
</style>
