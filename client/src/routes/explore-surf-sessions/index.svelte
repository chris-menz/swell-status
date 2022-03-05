<script lang="ts">
    import Nav from "$lib/components/nav.svelte";
    import { SurfSession, Favorite, Comment, User } from "$lib/utils/models"
    import { surfSessions, getSurfSessions } from "$lib/stores/surfSessionsStore"
    import { getComments } from "$lib/stores/commentsStore"
    import { getFavorites } from "$lib/stores/favoritesStore"
    import { getUsers } from "$lib/stores/userStore";
    import SessionsFeed from "$lib/components/sessionsFeed.svelte";
    import FeedFilters from "$lib/components/feedFilters.svelte";
    import { onMount } from "svelte";
    import { fly } from "svelte/transition";

    let filteredSessionsArr: SurfSession[]

    let feedLoading: boolean;
    let feedLoaded: boolean;

    onMount(async () => {
        feedLoading = true;
        feedLoaded = false;
        const surfSessionResponse = await getSurfSessions()
        const commentsResponse = await getComments()
        const favoritesResponse = await getFavorites()
        const usersResponse = await getUsers()

        const success = 
            surfSessionResponse === "success" 
            && commentsResponse === "success"
            && favoritesResponse === "success"
            && usersResponse === "success"

        if (success){
            filteredSessionsArr = $surfSessions.filter(session => session.is_public == true)
        }

        feedLoading = false;
        feedLoaded = true;
        
    })

    function filterSessionsByRegion(region: string){
        filteredSessionsArr = $surfSessions.filter((session: SurfSession) => session.spot_region === region && session.is_public == true)
    }

    function filterSessionsBySpot(spot){
        filteredSessionsArr = $surfSessions.filter(session => session.surf_spot == spot && session.is_public == true)
    }
</script>

<main>
    <Nav />
    <div class="container" in:fly="{{ delay: 500, duration: 250 }}">
        <FeedFilters 
            on:filterRegion={e => {filterSessionsByRegion(e.detail.region)}} 
            on:filterSpot={e => filterSessionsBySpot(e.detail.spot)}
            on:removeFilters={() => filteredSessionsArr = $surfSessions.filter(session => session.is_public == true)}
        />
        {#if feedLoaded}
            <SessionsFeed {filteredSessionsArr}/>
        {/if}
    </div>
</main>

<style>
    *, *::before, *::after {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

    :global(body) {
        background-color: #161616;
    }

    .container {
        margin-top: 5em;
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
    }

    ::-webkit-scrollbar {
  		width: 10px;
	}

	::-webkit-scrollbar-track {
		background: #161616;
	}

	::-webkit-scrollbar-thumb {
		background: #888;
		border-radius: 5px;
	}

	/* Handle on hover */
	::-webkit-scrollbar-thumb:hover {
		background: #555;
	}

    @media (max-width: 950px){
        .container {
            flex-direction: column;
            align-items: center;
            justify-content: center;
            margin-right: 1em;
            margin-left: 1em;
        }
    }


</style>