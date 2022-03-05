<script lang="ts">
    import { breaks, regions } from '$lib/utils/breaks';
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();   

    let has_selected_region = false;
    let surf_break_names = []
    
    let region;
    let surf_break;

    let displayErrorMustSelect = false
    let displaySelectRegionFirstMessage = false

    $: surf_break_names = breaks.filter(surfBreak => surfBreak.region == region).map(surfBreak => surfBreak.break_name).sort()

</script>

<main>
    <div class="menu">
        <div class="header">
            Live Conditions<br>& 7-day Forecast
        </div>
        <select class="region-select" bind:value={region} on:change={() => has_selected_region = true}  on:click={() => {
            displayErrorMustSelect = false
            displaySelectRegionFirstMessage = false;
        }}>
            <option value="" disabled selected>Select Region</option>
            {#each regions as region}
                <option value="{region}">{region}</option>
            {/each}
        </select>

        <select class="break-select" bind:value={surf_break} on:click={() => {
            displayErrorMustSelect = false
            if(!region){
                displaySelectRegionFirstMessage = true;
            }
        }}>
            <option value="" disabled selected>Select Surf Spot</option>
            {#if has_selected_region}
                {#each surf_break_names as break_option}
                    <option value="{break_option}">{break_option}</option>
                {/each}
            {/if}
        </select>

        <button type="button" on:click={() => {
            if(surf_break && region) {
                dispatch('breakChange', {surf_break, region});
                displayErrorMustSelect = false;
            }
            else{
                displayErrorMustSelect = true
            }     
        }}>
            Get Conditions
        </button>

        {#if displayErrorMustSelect}
            <div class="error">Must select break to search</div>
        {/if}
        {#if displaySelectRegionFirstMessage}
            <div class="error">Select a region first</div>
        {/if}
    </div>
</main>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        font-family: sans-serif;
        margin: 0;
        padding: 0;
    }

    .menu {
        background-color: #313131;
        padding: 1em;
        display: flex;
        flex-direction: column;
        justify-content: center;
		align-items: center;
        width: 30vw;
        min-height: 25vh;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
    }

    .header {
        color: rgb(240, 234, 234);
        font-size: 2.5em;
        font-family: Verdana, sans-serif;
        text-shadow: 1px 1px 2px #000000;
        text-align: center;
        padding-bottom: 0.5em;
        margin: 0;
    }
    
    .region-select, .break-select {
        background-color: #111111;
        font-family: Verdana, Helvetica, sans-serif;
        color: rgb(240, 234, 234);
        border: none;
        border-radius: 5px;
        width: 85%;
        text-align: center;
        margin-bottom: 0.5em;
        padding: 0.5em;
        font-size: 1.5em;
        cursor: pointer;
        transition-duration: 300ms;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        min-height: 5vh;
        width: 70%;
        text-align: center;
        font-size: 1.5em;
        font-family: Verdana, sans-serif;
        border: none;
        border-radius: 5px;
        margin: 0.5em 0 0.5em 0;
        cursor: pointer;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        transition-duration: 300ms;
        z-index: 1;
    }

    button::before{
		position: absolute;
		content: "";
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		border-radius: 5px;
		background: linear-gradient(to right bottom, rgb(59, 31, 110),rgb(57, 39, 122)0);
		transition: opacity 0.3s;
		z-index: -1;
		opacity: 0;
	}

    button:hover::before {
		opacity: 1;
	}

    button:hover {
        color: rgb(182, 182, 182);
    }

    .region-select:hover, .break-select:hover {
        background-color: #030303;
    }

    select:defined {
        border: none;
        outline: none;
    }

    .error {
        color: white;
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
        text-align: center;
    }
    
    @media (max-width: 1252px) {
        
        /* .menu {
            padding: 1.5em;
        } */

        .break-select, .region-select {
            font-size: 1.3em;
            width: 85%;
        }

        button {
            font-size: 1.3em;
        }
    }
    
    @media (max-width: 1100px) {
        .menu {
            width: 35vw;
            padding: 1em;
            padding-top: 00.75em;
        }

        .header {
            font-size: 2em;
        }
    }
    

    
    @media (max-width: 800px) {
        .menu {
            width: 45vw;
            min-height: 30vw;
        }

        .break-select, .region-select {
            font-size: 1.3em;
            font-family: Roboto;
            width: 90%;
        }

        button {
            width: 75%;
            font-size: 1.1em;
        }

        button:hover, .region-select:hover, .break-select:hover {
            outline: none;
        }

        button:active {
            background-color: rgba(10, 5, 39, 0.336);
        }
    }

    @media (max-width: 600px) {
        .menu {
            width: 60vw;
        }
    }

    @media (max-width: 500px) {
        .menu {
            width: 82vw;
        }
    }

    
    
</style>