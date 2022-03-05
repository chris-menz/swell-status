<script lang="ts">
    import SpotOutlook from "./spotOutlook.svelte";
    import { user, isAuthenticated } from "$lib/stores/userStore"
    import { onMount } from "svelte";
    import { getSavedSpots, savedSpots } from "$lib/stores/savedSpotStore";
    import { goto } from "$app/navigation";

	let savedSpotsLoaded: boolean;

    onMount(async () => {
		savedSpotsLoaded = false

        if(await isAuthenticated()){
			getSavedSpots($user)
        } else {
            $savedSpots = []
        }
        
        savedSpotsLoaded = true
	})
</script>

<main>
    <div class="container">
        <div class="header">
            My Spots
        </div>
        {#if savedSpotsLoaded}
            <div class="spots-container">
                {#if $user && $savedSpots.length > 0}
                    {#each $savedSpots as spot}
                        <div class="spot"><SpotOutlook spotName={spot.spot_name} spotRegion={spot.spot_region} on:fullReportRequest /></div>
                    {/each}
                {/if}
                {#if $savedSpots.length == 0}
                    <div class="no-spots-saved-container">
                        <div class="no-spots-saved">You haven't saved any spots yet<br><br>Save spots for quick access to conditions</div>
                    </div>
                {/if}
                {#if !$user}
                    <div class="not-logged-in-container">
                        <div class="not-logged-in">You are not logged in.<br>Log in to save surf spots</div>
                        <button class="login-button" on:click={() => goto("/login")}>Login/Signup</button>
                    </div>
                {/if}
            </div>
        {/if}
        
    </div>
</main>

<style>
    * {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

    .container {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .header {
        color: white;
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        text-align: center;
        padding: 0.5em 1em;
        font-size: 3em;
        font-family: Verdana, sans-serif;
        margin-bottom: 0.5em;
    }

    .spots-container {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .spot {
        margin-bottom: 0.7em;
    }

    .no-spots-saved-container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 1em;
        width: 475px;
        margin-bottom: 1em;
    }

    .no-spots-saved {
        font-size: 1.5em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.2em;
        margin-right: 1em;
		color: white;
        text-align: center;
    }

    .not-logged-in-container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 1em;
        width: 475px;
    }

    .not-logged-in {
        font-size: 1.5em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.2em;
        margin-right: 1em;
		color: white;
        text-align: center;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        text-align: center;
        font-family: Verdana;
        font-size: 1.1em;
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        border-radius: 5px;
        border: none;
        outline: none;
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


     @media (max-width: 950px){
        .header {
            display: none;
        }

        .spots-container {
            margin-top: 3em;
        }
     }

     @media (max-width: 700px){
        .no-spots-saved-container {
            width: 90%;
        }

        .not-logged-in-container {
            width: 90%;
        }
     }

     @media (max-width: 500px) {
        .not-logged-in-container {
            flex-direction: column;
        }

        .not-logged-in {
            margin-right: 0;
            margin-bottom: 00.5em;
        }
     }
</style>