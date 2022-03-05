<script lang="ts">
	import { user, logout } from "$lib/stores/userStore"
	import { slide } from "svelte/transition";
    import { goto } from "$app/navigation";
import { getSavedSpots } from "$lib/stores/savedSpotStore";
	
	let open = false;
</script>

<main>
	<nav class="nav-container">
		<div class="logo">
			<a href="/">SwellStatus</a>	
			
		</div>
		<div href="" class="toggle-button" on:click={()=> open = !open} >
			<span class="bar"></span>
			<span class="bar"></span>
			<span class="bar"></span>
		</div>
		<ul class="nav-list">
			<li><a href="/surf-reports">Surf Reports</a></li>
			<li><a href="/my-surf-sessions">My Surf Sessions</a></li>
			<li><a href="/explore-surf-sessions">Explore Surf Sessions</a></li>
			{#if $user}
				<li on:click={async () => {
					await logout()
					getSavedSpots($user)
					goto("/")
				}}><a href="/surf-reports">Logout</a></li>
			{/if}
			{#if !$user}
				<li><a href="/login">Login/Signup</a></li>
			{/if}

		</ul>
		<ul class="burger-nav-list" style="display: {open ? "flex" : "none"}" transition:slide="{{ duration: 400 }}">
			<li><a href="/surf-reports">Surf Reports</a></li>
			<li><a href="/my-surf-sessions">My Surf Sessions</a></li>
			<li><a href="/explore-surf-sessions">Explore Surf Sessions</a></li>
			{#if $user}
				<li on:click={async () => {
					await logout()
					goto("/")
				}}><a href="/surf-reports">Logout</a></li>
			{/if}
			{#if !$user}
				<li><a href="/login">Login/Signup</a></li>
			{/if}
		</ul>
	</nav>
</main>

<style>
    *, *::before, *::after {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

	.nav-container {
		font-family: sans-serif;
		position: fixed;
		width: 100%;
		top: 0;
		left: 0;
		z-index: 999;
		background-image: linear-gradient(to right bottom, #6a37c2, #49329e);
		/* background-color: #6e38cc; */
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	

	.nav-list {
		list-style: none;
		position: relative;
		display: flex;
		margin-right: 10px;
	}


	nav ul li a {
		font-size: 1.2em;
		padding: 1rem;
		display: block;
	}



	a {
		background-color: transparent;
		text-decoration: none;
	}

	.toggle-button {
		position: absolute;
		top: 0.75em;
		right: 0.75em;
		display: none;
		flex-direction: column;
		justify-content: space-between;
		width: 40px;
		height: 28px;
		cursor: pointer;
	}

	.bar {
		background-color: white;
		border-radius: 10px;
		width: 100%;
		height: 3px;
	}


	.nav-list a, .nav-list a:visited, .burger-nav-list a, .burger-nav-list a:visited {
		color: white;
		background-color: transparent;
		text-decoration: none;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		transition-duration: 250ms;
	}

	a:active {
		color: grey;
	}


	.nav-list a:hover, .burger-nav-list a:hover {
		border-radius: 5px;
		background-color: rgba(10, 5, 39, 0.336);
		color: rgb(182, 182, 182);
		text-decoration: none;
		transition-duration: 250ms;
	}
	
	.logo{
		font-size: 2.5em;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		margin: 00.5rem;
		text-shadow: 2px 2px 4px #000000;
	}

	.logo a, .logo a:visited {
		color: rgb(240, 234, 234);
	}

	.burger-nav-list {
		display: none;
		list-style: none;
		
	}
	@media (max-width: 700px) {

		.logo {
			font-size: 2em;
		}
		
		.toggle-button {
			display: flex;
		}

		.nav-list {
			display: none;
		}

		.nav-container {
			flex-direction: column;
			align-items: flex-start;
		}

		.burger-nav-list {
			flex-direction: column;
			width: 100%;
			list-style: none;
		}

		.nav-list li {
			text-align: center;
		}

	}
	
</style>