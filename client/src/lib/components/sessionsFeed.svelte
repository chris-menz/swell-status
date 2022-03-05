<script lang="ts"> 
    import { SurfSession } from "$lib/utils/models"
    import SurfSessionPost from "$lib/components/surfSessionPost.svelte"
    import { getSurfSessions } from "$lib/stores/surfSessionsStore"
    import { onMount } from "svelte";

    export let filteredSessionsArr: SurfSession[];

    onMount(async () => {
        await getSurfSessions()
    })

</script>

<main>
    <div class="feed">
        {#each filteredSessionsArr as session}
            <SurfSessionPost {session} />
        {/each}
        {#if filteredSessionsArr.length == 0}
            <div class="no-sessions-found-container">
                <div class="no-sessions-found">We couldn't find any sessions with those filters</div>
            </div>
        {/if}
    </div>
</main>

<style>
     *, *::before, *::after {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

    .no-sessions-found-container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 1em;
        width: 450px;
        margin-bottom: 1em;
    }

    .no-sessions-found {
        font-size: 1.5em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.2em;
        text-align: center;
        margin-right: 1em;
		color: white
    }

    @media (max-width: 700px) {
        .no-sessions-found-container {
            width: 100%;
        }
    }
</style>
