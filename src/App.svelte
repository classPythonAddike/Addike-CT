<script>
    import Solution from './Solution.svelte'
</script>

<div class="container">
    {#each Array(10) as i}
        <Solution/>
    {/each}
</div>

<style>
.container {
    display: flex;
    flex-wrap: wrap;
}
</style>