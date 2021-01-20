<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";

  // TODO: Connect with backend
  const participants = {
    Hugo: 10,
    Piras: 3,
    Baetsle: 2,
    Pierre: 2,
    Quentain: 100,
    Thibaut: 3,
    Basile: 1,
    Evan: 3,
  };
</script>

<Container>
  {#each Object.entries(participants) as [participant, bars]}
    <Card
      on:click={() => participants[participant]++}
      on:longClick={() => {
        if (participants[participant] > 0) participants[participant]--;
      }}
    >
      <span slot="title">{participant}</span>
      <span slot="body">
        {#each Array((bars - (bars % 5)) / 5) as _}
          <span class="tally-marks">IIII</span>
        {/each}
        <span>{"I".repeat(bars % 5)}</span>
      </span>
    </Card>
  {/each}
</Container>

<style>
  .tally-marks {
    text-decoration: line-through;
    padding: 0.2rem;
  }
</style>
