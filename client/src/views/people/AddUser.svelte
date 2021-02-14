<script>
  import Input from "../../components/Input.svelte";
  import store from "../../store";

  let showInput = false;

  function addParticipant(e) {
    store.addParticipant(e.detail);
    showInput = false;
  }
</script>

<div class="add-btn" on:click={() => (showInput = true)}>+</div>
<Input
  type="text"
  on:confirm={addParticipant}
  on:cancel={() => (showInput = false)}
  show={showInput}
/>

<style>
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
