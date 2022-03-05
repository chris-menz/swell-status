<script>
    import { _16point } from '$lib/utils/_16point'
	import IconArrowDown from "~icons/iconoir/arrow-down"
	import { fly } from "svelte/transition";
	
	export let weather;

	let displayTooltip = false
	const tooltip = "Wind has a pronounced effect on the shape of the waves. Ideal is light wind under 5kts or offshore wind that blows from the shore out to the ocean, cleaning up the waves. Onshore wind that blows from the ocean to the shore over 10kts will crumble the waves by knocking them over before they get a chance to form, and will chop up the oceans's surface as well"
	
</script>

<main>
	<div class="card">

		<div class="mobile-header-container"
			on:click={() => displayTooltip = !displayTooltip}>

			<div class="header">
				Wind
			</div>

			<div class="tooltip">
				{#if displayTooltip}
					<div class="tooltip-message" transition:fly="{{ delay: 50, duration: 200 }}">{tooltip}</div>
				{/if}
				<div class="tooltip-toggle">
					{displayTooltip ? "X" : "?"}
				</div>
				
			</div>
			
		</div>

			<div class="bigscreen-header-container"
				on:mouseenter={() => displayTooltip = true}
				on:mouseleave={() => displayTooltip = false}
				on:click={() => displayTooltip = !displayTooltip}>

			<div class="header">
				Wind
			</div>

			<div class="tooltip">
				{#if displayTooltip}
					<div class="tooltip-message" transition:fly="{{ delay: 50, duration: 200 }}">{tooltip}</div>
				{/if}
				<div class="tooltip-toggle">
					{displayTooltip ? "X" : "?"}
				</div>
				
			</div>
				
		</div>

		<div class="wind">{weather.wind.speed.toFixed(0)} {+weather.wind.speed.toFixed(0) == 1 ? "kt" : "kts"}
		{#if weather.wind.gust}
			<br>{weather.wind.gust.toFixed(0)} {+weather.wind.gust.toFixed(0) == 1 ? "kt" : "kts"} Gusts
		{/if}<br>{_16point(weather.wind.deg)} {weather.wind.deg}°</div>
		<IconArrowDown style="font-size: 3.5em; transform: rotate({weather.wind.deg}deg); margin-top: -0.1em " />
	</div>
    
    
</main>

<style>
	
	* {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

	.card {
		max-width: 250px;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.bigscreen-header-container, .mobile-header-container {
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	.header {
		font-size: 2em;
		font-family: Verdana, sans-serif;
		text-decoration: underline;
		cursor: pointer;
		transition-duration: 250ms;
	}


	.wind {
		font-size: 1.3em;
		padding-top: 0.2em;
		font-family: Helvetica, sans-serif;
		font-weight:lighter;
	}

	.tooltip {
		position: relative;
		cursor: pointer;
	}

	.tooltip-message {
		position: absolute;
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
		color: rgb(211, 211, 211);
		font-size: 0.9em;
		line-height: 1.25em;
		background-color: #111111;
		border-radius: 5px;
		padding: 0.5em 1em;
		bottom: 30px;
		left: -100px;
		width: 250px;
		height: 100px;
		overflow-y: scroll;
		transition-duration: 300ms;
	}

	.tooltip-toggle {
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
		font-size: 0.9em;
		margin-left: 2px;
	}

	@media (max-width: 768px){
		.mobile-header-container {
			display: flex;
		}

		.bigscreen-header-container {
			display: none;
		}
	}

	@media (min-width: 769px){
		.mobile-header-container {
			display: none;
		}

		.bigscreen-header-container {
			display: flex;
		}
	}

</style>