<script>
  import { createEventDispatcher } from "svelte";
  import { fade } from "svelte/transition";

  // TODO : refactor me bro

  export let show = false;
  export let type = "text";

  let name = "";

  const dispatch = createEventDispatcher();

  $: if (!show) name = "";
</script>

{#if show}
  <div in:fade class="background" on:click={() => dispatch("cancel")}>
    <div class="container" on:click|stopPropagation>
      {#if type === "text"}
        <label for="name">Name :</label>
        <input type="text" id="name" maxlength="15" bind:value={name} />
        <button
          disabled={!(name.length > 2)}
          on:click={() => dispatch("confirm", name)}>Ok</button
        >
      {:else if type === "confirm"}
        <slot name="confirm-text" />
        <button on:click={() => dispatch("confirm")}>
          <slot name="confirm-btn" />
        </button>
      {/if}
    </div>
  </div>
{/if}

<style>
  .background {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
  }

  .container {
    width: max-content;
    max-width: 85%;
    min-height: max-content;
    border-radius: 1rem;
    box-shadow: 0 0 0.4em var(--primary);

    margin: auto;
    padding: 1rem;
    color: #000;
    background-color: #fff;

    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  label {
    display: block;
  }

  input {
    margin: 0.4rem 0;
    background-color: #f5f5f5;
    width: 10rem;
  }
</style>
