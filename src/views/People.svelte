<script>
  import Card from "../components/Card.svelte";
  import Container from "../components/Container.svelte";

  // TODO: Connect with backend
  let participants = [
    { name: "Hugo", bars: 10 },
    { name: "Piras", bars: 3 },
    { name: "Baetsle", bars: 3 },
    { name: "Pierre", bars: 3 },
    { name: "Quentain", bars: 100 },
    { name: "Basile", bars: 3 },
    { name: "Evan", bars: 3 },
    { name: "Thibaut", bars: 3 },
  ];

  function addParticipant(_, name) {
    participants = [
      ...participants,
      { name: name ?? "New participant", bars: 0 },
    ];
  }
</script>

<Container>
  {#each participants as { name, bars }}
    <Card
      on:click={() => bars++}
      on:longClick={() => {
        if (bars > 0) bars--;
      }}
    >
      <span slot="title">{name}</span>
      <span slot="body">
        {#each Array((bars - (bars % 5)) / 5) as _}
          <span class="tally-marks">IIII</span>
        {/each}
        <span>{"I".repeat(bars % 5)}</span>
      </span>
    </Card>
  {/each}
</Container>
<div class="add-btn" on:click={addParticipant}>+</div>

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
