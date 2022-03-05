<script lang="ts">
    import { fly } from "svelte/transition";
    import PersonalSurfSession from "$lib/components/personalSurfSession.svelte";
    import { user, } from "$lib/stores/userStore"
    import { onMount } from "svelte";
    import { getSurfSessions, surfSessions } from "$lib/stores/surfSessionsStore"
    import { goto } from "$app/navigation";
    import axios from "axios"
    import { axiosConfig } from "$lib/utils/axiosConfig"
    import { endpoint } from "$lib/utils/endpoint"
    import IconRefreshDouble from "~icons/iconoir/refresh-double"

    let surf_sessions = [];

    let sessionsLoading = false;
    let sessionsLoaded = false;

    onMount(async () => {
        sessionsLoaded = false
        sessionsLoading = true

        const response = await axios.get(endpoint + "/auth/getcurruserid", axiosConfig)
        const data = response.data;
		
        if(data.message == "user is logged in"){
            const userres = await axios.get(`${endpoint}/user/${data.id}`)
            $user = userres.data
            const response2 = await getSurfSessions()
            if(response2 === "success"){
                surf_sessions = $surfSessions.filter(session => session.creator_id == $user.id)
            }  
        }
        else {
            $user = null
            surf_sessions = []
        }


        sessionsLoaded = true
        sessionsLoading = false
    })

    $: if($user){
        surf_sessions = $surfSessions.filter(session => session.creator_id == $user.id)
    }

</script>

<main>
    <div class="my-sessions-container">
        <div class="my-sessions-header header">My Sessions</div>
        {#if sessionsLoading}
            <div class="loading-container" in:fly="{{ duration: 200 }}">
                <div class="loading">Loading Sessions</div>
                <div class="loading-icon"><IconRefreshDouble style="margin: -9.8px; padding: 0;"/></div>
            </div>
        {/if}
        {#if sessionsLoaded}
            
            <div class="sessions-container">
                {#each surf_sessions as session}
                    <PersonalSurfSession {session} />
                {/each}
            </div>
        {/if}
        {#if sessionsLoaded && surf_sessions.length == 0}
            <div class="no-sessions-yet-container" in:fly="{{ delay: 350, duration: 200 }}">
                <div class="no-sessions-yet">You haven't saved any sessions yet<br><br>Save sessions to track what conditions work well at your local spots</div>
            </div>
        {/if}
        {#if !$user && sessionsLoaded}
            <div class="not-logged-in-container" in:fly="{{ delay: 350, duration: 200 }}">
                <div class="not-logged-in">You are not logged in.<br>Log in to save sessions</div>
                <button class="login-button" on:click={() => goto("/login")}>Login/Signup</button>
            </div>
        {/if}
    </div>
</main>


<style>

    .loading-container {
		font-size: 2em;
		color: white;
		text-align: center;
		transition-duration: 100ms;
		display: flex;
		flex-direction: row;
		justify-content: center;
		width: 100%;
        margin-top: 3em;
	}

	.loading {
		padding-right: 0.75em;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		font-weight: lighter;
	}
	.loading-icon {
		animation-name: spin;
		animation-duration: 1.5s;
		animation-iteration-count: infinite;
		animation-timing-function: linear;
	}

    @keyframes spin {
        from {
            transform:rotate(0deg);
        }
        to {
            transform:rotate(360deg);
        }
    }
    .header {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        font-size: 3em;
        font-family: Verdana, Geneva, Tahoma, sans-serif;
        padding: 0.5em 1em;
        margin: 0 0 0.5em 0;
        color: #f0f0f0;;
        text-align: center;
    }

    .my-sessions-container {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .sessions-container {
        display: flex;
        flex-direction: column;
        align-items: center;
    }


    .not-logged-in-container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 1em;
        width: 450px;
    }

    .not-logged-in {
        font-size: 1.5em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.2em;
        text-align: center;
        margin-right: 1em;
		color: white;
    }

    .no-sessions-yet-container {
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

    .no-sessions-yet {
        font-size: 1.5em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.2em;
        text-align: center;
        margin-right: 1em;
		color: white
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

    @media (max-width: 1000px) {
        .header {
            display: none;
        }
    }

    @media (max-width: 700px) {
        .header {
            width: 80vw;
        }

        .not-logged-in-container {
            width: 85%;
        }

        .no-sessions-yet-container {
            max-width: none;
            width: 85%;
        }
    }
    @media (max-width: 500px) {
        .header {
            font-size: 2.5em;
        }

        .not-logged-in-container {
            flex-direction: column;
        }

        .not-logged-in {
            margin-right: 0;
            margin-bottom: 0.5em;
        }
    }

    @media (max-width: 400px) {
        .header {
            font-size: 2.2em;
        }
    }
</style>