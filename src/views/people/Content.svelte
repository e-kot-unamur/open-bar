<script>
  import Separator from "../../components/Separator.svelte";
  import store from "../../store";
</script>

<Separator>
  <span slot="title">People</span>
  <span slot="body">
    {#if $store.users?.length}
      Press on the desired participant to add a bar, or long press to remove one
    {:else}
      Start by adding someone first !
    {/if}
  </span>
</Separator>
