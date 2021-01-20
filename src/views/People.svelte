<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";

  // TODO: Connect with backend
  const participants = {
    Hugo: 10,
    Piras: 3,
    Baetsle: 2,
    Pierre: 2,
    Quentain: 56,
    Thibaut: 3,
    Basile: 1,
    Evan: 3,
  };

  const temp = () => {
    
  };
</script>

<Container>
  {#each Object.entries(participants) as [participant, bars]}
    <Card on:click={() => participants[participant]++}>
      <span slot="title">{participant}</span>
      <span slot="body">
        {"|".repeat(bars)}
      </span>
    </Card>
  {/each}
</Container>
