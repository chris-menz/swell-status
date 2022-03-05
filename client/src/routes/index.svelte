<script>
    import Nav from "$lib/components/nav.svelte";
    import { user } from "$lib/stores/userStore";
    import { goto } from "$app/navigation";
    import { fly } from "svelte/transition";
</script>

<main>
    <Nav />
    <div class="container" in:fly="{{ delay: 400, duration: 250 }}">
       <div class="header">
            Welcome to SwellStatus.com{#if $user}, {$user.first_name}{/if}!
        </div> 
        
        
        <div class="section">
            <div class="button">
                <button on:click={() => goto("/surf-reports")}>
                    Surf Reports
                </button>
            </div>
            <div class="content">
                View surf conditions (live + 7-day forecast) for any mainland United States surf spot
            </div>
        </div>
        <div class="section">
            <div class="button">
                <button on:click={() => goto("/my-surf-sessions")}>
                    My Surf Sessions
                </button>
            </div>
            <div class="content">
                Log your sessions to keep track of what conditions generate good waves at your local surf spots
            </div>
        </div>
        <div class="section">
            <div class="button">
                <button on:click={() => goto("/explore-surf-sessions")}>
                    Explore
                </button>
            </div>
            <div class="content">
                Check out what other surfers in your area are doing on the Explore page
            </div>
        </div>
        {#if !$user}
            <div class="section">
                <div class="button">
                    <button on:click={() => goto("/login")}>
                        Login/Signup
                    </button>
                </div>
                <div class="content">
                    New user? Create an account to save surf spots and log your surf sessions
                </div>
            </div>
        {/if}
    
    </div>
    
</main>

<style>
	:global(body) {
        background-color: #161616;
        transition: background-color 0.3s
    }

    .container {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin: 6em auto;
        width: 65vw;
    }

    .header {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        font-size: 2.5em;
        font-family: Verdana, sans-serif;
        padding: 0.5em 1em;
        margin: 0 0 0.75em 0;
        color: #f0f0f0;
        text-align: center;
    }

    .section {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        margin: 0 0 1.25em 0;
        padding: 1em 1em;
        display: flex;
    }

    .content {
        font-size: 2em;
        font-family: Helvetica, sans-serif, sans-serif, sans-serif;
        font-weight: lighter;
        color: #f0f0f0;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        text-align: center;
        font-family: Verdana;
        font-size: 1.4em;
        color: rgb(240, 234, 234);
        width: 210px;
        margin-right: 1em;
        padding: 0.5em 1em;
        border-radius: 5px;
        border: none;
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

    @media (max-width: 700px) {

        .container {
            width: 90vw;
            margin-top: 4.5em;
        }
        .header{
            font-size: 2em;
            margin-bottom: 1em;
        }
        .section {
            flex-direction: column;
            align-items: center;
            justify-content: center;
            margin-bottom: 1em;
        }

        .content {
            font-size: 1.5em;
            text-align: center;
        }

        button {
            margin-bottom: 00.5em;
            margin-right: 0%;
        }
    }

    @media (max-width: 375px){
        .header {
            font-size: 1.6em;
        }
    }

</style>