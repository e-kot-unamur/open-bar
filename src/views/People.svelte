<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";
  import Input from "../components/Input.svelte";
  import Separator from "../components/Separator.svelte";
  import store from "../store";

  let showInput = false;
</script>

<Container>
  {#if $store.users && $store.users.length}
    {#each $store.users as { id, name, debt }}
      <Card
        on:click={() => store.updateDebt(id, debt + 1)}
        on:longClick={() => {
          if (debt > 0) store.updateDebt(id, debt - 1);
        }}
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
  {:else}
    <Separator>
      <span slot="title">There is nothing here yet...</span>
      <span slot="body">Start by adding someone first !</span>
    </Separator>
  {/if}
</Container>

<div class="add-btn" on:click={() => (showInput = true)}>+</div>
<Input
  on:confirm={(e) => {
    store.addParticipant(e.detail);
    showInput = false;
  }}
  show={showInput}
/>

<style>
  .tally-marks {
    text-decoration: line-through;
    padding: 0.2rem;
  }

  .add-btn {
    position: fixed;
    bottom: 1rem;
    right: 1rem;
    z-index: 100;
    background-color: var(--primary);
    width: 2.5rem;
    height: 2.5rem;
    line-height: 2.5rem;
    text-align: center;
    border-radius: 100%;
    user-select: none;
  }
</style>
