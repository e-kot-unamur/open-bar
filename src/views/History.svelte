<script>
    import Card from "../components/Card.svelte";
    import Container from "../components/Container.svelte";
    import Separator from "../components/Separator.svelte";
    import store from "../store";
</script>

<Container>
    <Separator>
        <span slot="title">History</span>
        <span slot="body">
            {#if !$store.history?.length}
                Nothing to show here yet
            {/if}
        </span>
    </Separator>
    {#each $store.history.reverse() as { date, targetId, numberOfBars }}
        <Card>
            <!-- FIXME : change users to a Map with the user id as the key ! -->
            <span slot="title">{$store.users[targetId].name}</span>
            <span class="card-body" slot="body">
                <p>{new Date(date).toLocaleString()}</p>
                <p>{numberOfBars}</p>
            </span>
        </Card>
    {/each}
</Container>

<style>
    .card-body {
        width: 100%;
        height: min-content;
        display: flex;
        justify-content: space-between;
    }
</style>
