<script lang="ts">
    import AddSession from "$lib/components/addSession.svelte";
    import MySurfSessionsFeed from "$lib/components/mySurfSessionsFeed.svelte";
    import Nav from '$lib/components/nav.svelte';
    import { onMount } from 'svelte';
    import { fly } from "svelte/transition";

    let mySessionsViewActive
    let logSessionViewActive

    const toggleMySessionsView = () => {
		mySessionsViewActive = true
        logSessionViewActive = false
	}

	const toggleLogSessionView = () => {
        mySessionsViewActive = false
        logSessionViewActive = true
    }

    onMount(async () => {
        toggleMySessionsView()
    })

</script>

<main class="small-screen-view">
    <Nav />
    <div class="tabs">
        <div class="tab" style="background: {mySessionsViewActive ? "linear-gradient(to right bottom, #6a37c2, #49329e)" : "none"}" on:click={() => toggleMySessionsView()}>My Sessions</div>
        <div class="tab"  style="background: {logSessionViewActive ? "linear-gradient(to right bottom, #6a37c2, #49329e)" : "none"}" on:click={() => toggleLogSessionView()}>Log a Session</div>
    </div>
    <div class="container" in:fly="{{ delay: 350, duration: 200 }}">
        
        {#if logSessionViewActive}
            <AddSession />
        {/if}
        {#if mySessionsViewActive}
             <MySurfSessionsFeed />  
        {/if}
    </div>
</main>

<main class="large-screen-view">
    <Nav />
    <div class="container" in:fly="{{ delay: 380, duration: 200 }}">
        <AddSession />
        <MySurfSessionsFeed />  
    </div>
    
</main>

<style>
    /* fix responsiveness */
	* {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

	:global(body) {
        background-color: #161616;
    }

    .small-screen-view {
        display: none;
    }

    .tabs {
		display: none;
		top: 3.90em;
		left: 0%;
		position: fixed;
		background-color: #2b2b2b;
		width: 100vw;
		height: 50px;
		cursor: pointer;
	}

	.tab {
		color: white;
		width: 50%;
		display: grid;
		place-content: center;
		text-transform: uppercase;
        font-family: Verdana, Geneva, Tahoma, sans-serif;
	}

    .container {
        margin: 5em 0 0 0;
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
    }


    @media (max-width: 1000px) {
        .large-screen-view {
            display: none;
        }

        .small-screen-view {
			display:inherit; 
		}

        .tabs {
			display: flex;
		}
        
        .container {
            flex-direction: column;
            margin: 8em 0 0 0;
        }

    }

    @media (max-width: 700px) {
        .tabs {
			top: 3.4em;
		}
    }

</style>