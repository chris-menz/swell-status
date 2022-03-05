<script lang="ts">
    import { SurfSession } from "$lib/utils/models"
    import { deleteSurfSession, getSurfSessions, surfSessions } from "$lib/stores/surfSessionsStore"
    import { fly, slide } from "svelte/transition"

    export let session: SurfSession;

    let displayAreYouSure = false

</script>

<main>
    <div class="surf-session-container" in:fly="{{ delay: 300, duration: 200 }}">
        <div class="header-container">
            <div class="session-spot">{session.surf_spot}</div>
            <div class="date">{session.datetime}</div>
        </div>
        <div class="session-location">{session.spot_location}</div>
        
        
        <div class="session-conditions">
            <div class="swell">Swell: {session.swell}</div>
            <div class="wind">Wind: {session.wind}</div>
            <div class="tide">Tide: {session.tide}</div>
        </div>
        <div class="session-description">{session.description}</div>
        {#if session.surfboard_id}
            <div class="session-surfboard"></div>
        {/if} 
        <button  on:click={() => displayAreYouSure = !displayAreYouSure}>Delete</button>
        {#if displayAreYouSure}
            <div class="are-you-sure" transition:slide="{{ duration: 400 }}">
                <div class="confirm-deletion-message">Confirm deletion (not reversable)</div>
                <button on:click={async () => await deleteSurfSession(session.id)} class="delete-btn">Delete Session</button>
                <button on:click={() => displayAreYouSure = false}>Keep Session</button>
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

     .surf-session-container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
       
         margin-bottom: 1em;
        padding: 0.5em 1em;
        color: white;
        width: 450px;
    }

    .header-container {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 0.2em;
    }

    .session-spot {
        font-size: 2em;
        padding-right: 00.75em;
        font-family: Verdana, sans-serif;
    }

    .date {
        font-size: 1em;
        font-family: Verdana, sans-serif;
        text-align: center;
        width: 123px;
        align-self: flex-start;
    }

    .session-location {
        font-size: 1.4em;
        margin-bottom: 0.4em;
        font-family: Verdana, sans-serif;
    }

    .session-conditions {
        margin-bottom: 00.6em;
        font-size: 1.1em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
    }

    .swell, .wind {
        margin-bottom: 0.2em;
    }

    .session-description {
        background-color: #111111;
        border-radius: 5px;
        padding: 0.4em 1em;
        margin-bottom: 0.8em;
        font-family: Verdana, sans-serif;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        padding: 0.5em 1em;
        text-align: center;
        font-size: 1em;
        font-family: Verdana, sans-serif;
        border: none;
        border-radius: 5px;
        margin-bottom: 0.5em;
        transition-duration: 300ms;
        cursor: pointer;
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

    .confirm-deletion-message {
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        margin-top: 0.25em;
        margin-bottom: 0.5em;
    }

    .delete-btn {
        margin-right: 0.25em
    }

    @media (max-width: 600px) {
        .surf-session-container {
            width: 85vw;
        }

        .header-container {
            align-items: flex-start;
            margin-bottom: 0;
        }

        .session-spot {
            font-size: 1.5em;
        }

        .date {
            font-size: 0.9em;
        }

        .session-location {
            font-size: 1.2em;
        }

        .session-conditions {
            font-size: 0.95em;
        }

        .session-description {
            font-size: 00.95em;
        }
    }
   
</style>