<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { breaks, regions } from "$lib/utils/breaks" 

    const dispatch = createEventDispatcher();

    let region
    let spot

    let has_selected_region = false

    let surf_break_names

    let disabled = false
    let displayDisbaledMessage = false
    let displaySelectRegionFirstMessage = false

    $: disabled = !spot && !region

    $: if(has_selected_region){
        surf_break_names = breaks.filter(surfSpot => region == surfSpot.region).map(surfBreak => surfBreak.break_name).sort()
    }
</script>

<main>
    <div class="container">
        <div class="select-container">
            <select class="region-select-header" bind:value={region} on:change={() => {
                    has_selected_region = true; 
                    spot = undefined; 
                    displayDisbaledMessage = false
                    displaySelectRegionFirstMessage = false;
                }
            }
                on:click={() => {
                    displayDisbaledMessage = false
                    displaySelectRegionFirstMessage = false;
                }
            }>
                <option value="" disabled selected>Select Region</option>
                {#each regions as region_option}
                    <option value="{region_option}">{region_option}</option>
                {/each}
            </select>
            
            <select class="break-select-header" bind:value={spot} on:click={() => {
                    displayDisbaledMessage = false
                    if(!region){
                        displaySelectRegionFirstMessage = true;
                    }
                }
            }>
                <option value="" disabled selected>Select Surf Spot</option>
                {#if has_selected_region}
                    {#each surf_break_names as break_option}
                        <option value="{break_option}">{break_option}</option>
                    {/each}
                {/if}
            </select>
        </div>
        
        <button on:click={() => {
                if(!disabled){
                   if(region){
                        dispatch("filterRegion", {region})
                    }
                    if(spot){
                        dispatch("filterSpot", {spot})
                    } 
                } else {
                    displayDisbaledMessage = true
                }
            }
        } type="button">Filter</button>

        <button on:click={() => {
            dispatch("removeFilters")
        }}>
            Remove Filters
        </button>

        {#if displayDisbaledMessage}
            <div class="messages">Select a region or spot first</div>
        {/if}
        {#if displaySelectRegionFirstMessage}
            <div class="messages">Select a region first</div>
        {/if}
    </div>
</main>

<style>
      *, *::before, *::after {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

    .container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        width: 450px;
        padding: 1em;
        margin-bottom: 1.5em;
    }

    .select-container {
        display: flex;
        flex-direction: column;
    }

    select {
        background-color: #111111;
        font-family: Verdana, sans-serif;
        color: rgb(240, 234, 234);
        border: none;
        border-radius: 5px;
        width: 60%;
        text-align: left;
        margin-bottom: 0.5em;
        padding: 0.5em;
        font-size: 1.3em;
        cursor: pointer;
        transition-duration: 300ms;
        margin-right: 1em;
    }

    select:defined {
        outline: none;
    }

    select:hover {
        background-color: #030303;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        text-align: center;
        font-size: 1em;
        font-family: Verdana, sans-serif;
        border: none;
        border-radius: 5px;
        margin: 0.5em 0.5em 0.5em 0;
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

    .messages {
        color: white;
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
        text-align: center;
    }

    @media (max-width: 950px){
        .container {
            width: 300px;
        }

        select {
            width: 100%;
        }
    }
</style>