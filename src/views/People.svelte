<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";
  import Input from "../components/Input.svelte";
  import Separator from "../components/Separator.svelte";
  import store from "../store";

  let showInput = false;
</script>

<Container>
  <Separator>
    <span slot="title">People</span>

    <span slot="body">
      {#if $store.users?.length}
        Press on the desired participant to add a bar, or long press to remove
        one
      {:else}
        Start by adding someone first !
      {/if}
    </span>
  </Separator>
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
</Container>

<div class="add-btn" on:click={() => (showInput = true)}>+</div>
<Input
  on:confirm={(e) => {
    store.addParticipant(e.detail);
    showInput = false;
  }}
  on:cancel={() => (showInput = false)}
  show={showInput}
/>

<style>
  .tally-marks {
    text-decoration: line-through;
    padding: 0.2rem;
  }

  .add-btn {
    position: fixed;
    margin: 1rem;
    bottom: 0;
    z-index: 100;
    background-color: var(--primary);
    width: 2.5rem;
    height: 2.5rem;
    line-height: 2.3rem;
    text-align: center;
    border-radius: 100%;
    user-select: none;
  }
</style>
