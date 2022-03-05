<script>
	import { onMount } from 'svelte'
	import { convert24HrToAmPm } from '$lib/utils/dateTimeUtils'
	import { fly } from "svelte/transition";
	
	export let tideExtremes;
	export let tideHeight;
	export let localDate;
	let curr_tide_data = tideExtremes.get(localDate.toString().split("T")[0]);;

	let tide_string_array = [];

	let displayTooltip = false
	const tooltip = "Tides affect every surf spot differently. Some spots work in all tides, while others are very tide sensitive. In general, higher tides tend to slow down the surf as there is more water in between the wave and the ocean floor, while low tides speed up the surf."
	
	onMount(async () => {
		for(let i = 0; i < curr_tide_data.length; ++i){
			const tide = curr_tide_data[i]
			const time = convert24HrToAmPm(tide.date.substring(11, 16))
			tide_string_array[i] = `${tide.type}: ${Math.round(tide.height * 3.281 * 10) / 10} ft, ${time}`;
		}
	})

	
</script>

<main>
	<div class="card">
		<div class="mobile-header-container"
			on:click={() => displayTooltip = !displayTooltip}>

				<div class="header">
					Tide
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
					Tide
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

		<div class="tides">
			<div class="tide-height">{tideHeight.height} ft, {tideHeight.state}</div>
			{#each tide_string_array as tide}
				<div class="tide">{tide}</div>
			{/each}
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
	.tides {
		line-height: 1.3em;
		font-size: 1em;
		padding: 0.4em 0;
		font-family: Helvetica, sans-serif;
		font-weight:lighter;
	}

	.tide-height {
		font-size: 1.3em;
		padding-bottom: 00.1em;
	}

	.tide {
		padding: 0.1em;
		font-size: 1em;
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
		right: -70px;
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

	@media(max-width: 500px) {
		.tide {
			font-size: 00.9em;
		}
	}
</style>