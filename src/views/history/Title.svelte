<script>
  import Separator from "../../components/Separator.svelte";
  import store from "../../store";
</script>

<Separator>
  <span slot="title">History</span>
  <span slot="body">
    {#if !$store.history?.length}
      Nothing to show here yet
    {/if}
  </span>
</Separator>
