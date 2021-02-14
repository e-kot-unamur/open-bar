<script>
  import Card from "../../components/Card.svelte";
  import Input from "../../components/Input.svelte";
  import Separator from "../../components/Separator.svelte";
  import store from "../../store.js";

  let showInput = false;
  let participantOption = false;

  function handleClick(keepParticipants) {
    showInput = true;
    participantOption = keepParticipants;
  }
</script>

{#if $store.users?.length}
  <Separator>
    <span slot="title">Reset application</span>
    <span slot="body">Long click on the desired option</span>
  </Separator>

  <Card important on:longclick={() => handleClick(true)}>
    <div class="card-title" slot="title">
      Reset all debts but keep participants
    </div>
  </Card>

  <Card important on:longclick={() => handleClick(false)}>
    <div class="card-title" slot="title">
      Reset all debts and delete all participants
    </div>
  </Card>

  <Input
    type="confirm"
    on:confirm={() => {
      store.reset(participantOption);
      showInput = false;
    }}
    on:cancel={() => (showInput = false)}
    show={showInput}
  >
    <div slot="confirm-text">
      <h3>Warning !</h3>
      <span>
        You're about to reset the application, are you sure ? Prior state will
        be kept within a file on the server.
      </span>
    </div>
    <div slot="confirm-btn">Yes I am</div>
  </Input>
{/if}

<style>
  .card-title {
    width: 100%;
    text-align: center;
  }
  h3 {
    font-weight: 400;
    text-align: center;
  }
</style>
