<script>
  import Card from "../../components/Card.svelte";
  import Input from "../../components/Input.svelte";
  let showInput = false;
</script>

<Card important on:longclick={() => (showInput = true)}>
  <div slot="title">Reset all debts</div>
</Card>

<Input
  type="confirm"
  on:confirm={() => console.log("todo, implement special event")}
  on:cancel={() => (showInput = false)}
  show={showInput}
>
  <h3 slot="confirm-text">
    You're about to delete all debts, are you sure ?
  </h3>
  <div slot="confirm-btn">Yes I am</div>
</Input>

<style>
  div {
    width: 100%;
    text-align: center;
  }
  h3 {
    font-weight: 400;
    text-align: center;
  }
</style>
