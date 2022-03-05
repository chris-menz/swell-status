<script>
	import { _16point } from '$lib/utils/_16point'
	import IconArrowDownCircled from "~icons/iconoir/arrow-down-circled"
	import { fly } from "svelte/transition";
	
	export let conditions;
	
	let displayTooltip = false
	const tooltip = "Swell is a measure of ocean energy that tells us how big the waves will be. Ex: 3 feet at 15 seconds means that every 15 seconds on average, 3 feet of energy is hitting the surf spot. \n\nThe shape of the ocean floor at the surf spot also affects wave height and shape. Combining swell information (and tide/wind conditions) with local knowledge of your surf spot can give you a fairly accurate estimate of wave size and shape"
	
</script>

<main>
	<div class="card">
		<div class="mobile-header-container"
			on:click={() => displayTooltip = !displayTooltip}>

			<div class="header">
				Swell
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
				Swell
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
		

		<div class="swell">{conditions.swell_height} ft. @ {conditions.swell_period}s<br>{_16point(+conditions.swell_direction)} {conditions.swell_direction}°</div>
		<IconArrowDownCircled style="font-size: 3em; transform: rotate({+conditions.swell_direction}deg); margin-bottom: 0.2em" />

		{#if +conditions.wind_wave_height > 1  && +conditions.wind_wave_height > +conditions.secondary_swell_height}
			<div class="swell">{conditions.wind_wave_height} ft. @ {conditions.wind_wave_period}s<br>{_16point(+conditions.wind_wave_direction)} {conditions.wind_wave_direction}°</div>
			<IconArrowDownCircled style="font-size: 3em; transform: rotate({+conditions.wind_wave_direction}deg); margin-bottom: 0.2em" />
		{/if}

		{#if +conditions.secondary_swell_height > 1 && +conditions.wind_wave_height < +conditions.secondary_swell_height}
			<div class="swell">{conditions.secondary_swell_height} ft. @ {conditions.secondary_swell_period}s<br>{_16point(+conditions.secondary_swell_direction)} {conditions.secondary_swell_direction}°</div>
			<IconArrowDownCircled style="font-size: 3em; transform: rotate({+conditions.secondary_swell_direction}deg); margin-bottom: 0.2em" />
		{/if}
		
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

	.swell {
		font-size: 1.2em;
		padding: 0.3em;
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