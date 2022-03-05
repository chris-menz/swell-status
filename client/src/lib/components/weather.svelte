<script>
	export let weather;	
	export let conditions;
	import { fly } from "svelte/transition";


	let displayTooltip = false
	const tooltip = "Wetsuit choice comes down to preference. The colder the water is, the thicker the suit needs to be. Gloves, booties and hoods also become essential below certain temperatures"
	
</script>

<main>
    <div class="card">

		<div class="mobile-header-container"
			on:click={() => displayTooltip = !displayTooltip}>

			<div class="header">
				Weather
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
				Weather
			</div>

			<div class="tooltip">
				{#if displayTooltip}
					<div class="tooltip-message" transition:fly="{{ delay: 50, duration: 200 }}">
						{tooltip}
					</div>
				{/if}
				<div class="tooltip-toggle">
					{displayTooltip ? "X" : "?"}
				</div>
				
			</div>

			
			
		</div>

		
		<div class="weather-container">
			<div class="weather">
				<div>Water: {conditions.water_temperature}°f</div>
				<div>Air: {weather.main.temp.toFixed(0)}°f</div>
				<div>{weather.weather[0].main}</div>
			</div>
			
			<!-- <img src={url} alt="weather"> -->
		</div>
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
	

	.weather {
		padding: 0.3em 0;
		font-size: 1.2em;
		line-height: 1.2em;
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
		right: -30px;
		width: 250px;
		height: 100px;
		overflow-y: scroll;
		transition-duration: 300ms;
	}

	.tooltip-toggle {
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
		font-size: 0.9em;
		margin-left: 4px;
	}

	.wetsuit {
		color: white;
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