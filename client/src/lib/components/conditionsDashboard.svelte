<script lang="ts">
    import axios from "axios"
    import { breaks, regions }  from "$lib/utils/breaks"
	import { conditionsParser } from "$lib/utils/conditionsParser"
	import { getTideHeight, tideParser } from "$lib/utils/tideUtils"
	import { convertMapToLocalTime } from "$lib/utils/convertMapKeysToLocalTime"
	import DailyForecast from "$lib/components/dailyForecast.svelte"
	import Swell from "$lib/components/swell.svelte"
	import Tide from "$lib/components/tide.svelte"
	import Wind from "$lib/components/wind.svelte"
	import Weather from "$lib/components/weather.svelte"
	import { onMount } from "svelte";
	import { fly } from "svelte/transition";
	import { savedSpots, saveNewSpot, deleteSpot } from "$lib/stores/savedSpotStore"
	import { user, isAuthenticated } from "$lib/stores/userStore"
	import { endpoint } from "$lib/utils/endpoint"
	import IconRefreshDouble from "~icons/iconoir/refresh-double"

	// api responses
	let conditionsMap: Map<string, object>;
	let localTimeConditionsMap: Map<string, object>;
	let conditions: Object;
	let weather: Object;
    let tideExtremes;
	let tideHeight;
	let localDate;

	// state vars
	let has_selected_region = false;
	let displayErrorMustSelect = false;
    let surfReportGenerated = false;
    let surfReportLoading = false;
	let spotIsSaved = false;


	type SurfSpot = {
        break_name: string;
        location: string;
        region: string;
        lat: string;
        lng: string;
    }
    
    let surf_break_object: SurfSpot;

    export let region: string;
    export let surf_break_name: string;
	let surf_break_names: string[];

	$: surf_break_names = breaks.filter(surfBreak => surfBreak.region == region).map(surfBreak => surfBreak.break_name).sort()

	onMount(() => {
		handleBreakChange()
	})

	async function handleBreakChange() {
		surfReportGenerated = false;
        surfReportLoading = true;
		spotIsSaved = false;
        surf_break_object = breaks.filter((br) => surf_break_name == br.break_name && region == br.region)[0];
		const d = new Date();

		// get local date to look up correct tides


		const unixUTCTime = Math.floor(new Date(d.toUTCString()).getTime() / 1000);
			
		// get time zone difference in seconds between local and utc
		const timeZoneResponse = await axios.get(endpoint + "/timezone/" + surf_break_object.lat + "/" + surf_break_object.lng)
		const utcOffset = timeZoneResponse.data.gmtOffset

		// convert utc unix time to local unix time
		const localTimeInUnix = unixUTCTime + utcOffset

		// convert from unix to iso format
		localDate = new Date(localTimeInUnix * 1000).toISOString();

		// need to get tides starting from yesterday to calculate
		// tide values that occur at times before the first extreme of the day
		const localDateYesterday = new Date((localTimeInUnix - 86400) * 1000).toISOString();

		if($savedSpots.length > 0){
			const spotSavedArr = $savedSpots.filter((br) => surf_break_name == br.spot_name && region == br.spot_region)
			spotIsSaved = (spotSavedArr.length > 0 ? true : false)
		}
		
		
		const lat = surf_break_object.lat;
		const lng = surf_break_object.lng;

        try {
            const conditions_response = await axios.get(`${endpoint}/sg/${lat}/${lng}`);
            conditionsMap = conditionsParser(conditions_response.data);
			localTimeConditionsMap = convertMapToLocalTime(utcOffset, conditionsMap)
            
            const weather_response = await axios.get(`${endpoint}/weather/${lat}/${lng}`);
            weather = weather_response.data;

            const tides_response = await axios.get(`${endpoint}/tide/${lat}/${lng}/${localDateYesterday.split("T")[0]}`);
            tideExtremes = tideParser(tides_response.data)
			tideHeight = getTideHeight(tides_response.data, unixUTCTime)
        } catch (error) {
            console.log(error)
        }


		const key = `${d.getUTCFullYear()}-${d.getUTCMonth() < 9 ? '0' : ''}${d.getUTCMonth() + 1}-${d.getUTCDate() < 10 ? '0' : ''}${d.getUTCDate()}T${d.getUTCHours() < 10 ? "0" : ""}${d.getUTCHours()}:00:00+00:00`
		
		conditions = conditionsMap.get(key)

		surfReportLoading = false;
        surfReportGenerated = true;
		
        // for the menu
		has_selected_region = true;
		surf_break_names = breaks.filter(surfBreak => surfBreak.region == region).map(surfBreak => surfBreak.break_name).sort();
	}
</script>

<main>
    {#if surfReportLoading}
		<div class="loading-container" in:fly="{{ duration: 200 }}">
			<div class="loading">Loading Surf Data</div>
			<div class="loading-icon"><IconRefreshDouble style="margin: -9.8px; padding: 0;"/></div>
		</div>
	{/if}
    {#if surfReportGenerated}
    <div class="conditions-container" in:fly="{{ delay: 320, duration: 200 }}">

        <div class="conditions-header">
            <div class="break-name">
                {surf_break_object.break_name}
            </div>
            <div class="break-location">
                {surf_break_object.location}
            </div>
			<div class="save-button-container">
				{#if $user && !spotIsSaved}
					<button class="save-button" on:click={() => {
						saveNewSpot(surf_break_object.break_name, surf_break_object.region, $user);
						// $savedSpots = [ ...$savedSpots]
						spotIsSaved = true;
					}}>
						Save
					</button>
				{/if}
				{#if $user && spotIsSaved}
					<button type="button" class="save-button" on:click={() => {
						const spotId = $savedSpots.filter(
							(spot) =>
							spot.spot_name == surf_break_object.break_name &&
							spot.spot_region == surf_break_object.region &&
							spot.user_id == $user.id
						)[0].id;
						deleteSpot(spotId, $user);
						spotIsSaved = false;
					}}>
						Remove
					</button>
				{/if}
			</div>
			
			

            <div class="header-menu">
				
				<!-- fix, use reactive statement -->
                <select class="region-select-header" bind:value={region} on:change={() => {surf_break_names = breaks.filter(surfBreak => surfBreak.region == region).map(surfBreak => surfBreak.break_name).sort(); has_selected_region = true; surf_break_name = undefined}}>
                    <option value="" disabled selected>Select Region</option>
                    {#each regions as region_option}
                        <option value="{region_option}">{region_option}</option>
                    {/each}
                </select>
                
                <select class="break-select-header" bind:value={surf_break_name}>
                    <option value="" disabled selected>Select Surf Spot</option>
                    {#if has_selected_region}
                        {#each surf_break_names as break_option}
                            <option value="{break_option}">{break_option}</option>
                        {/each}
                    {/if}
                </select>

                <button type="button" class="search-button" on:click={() => {
                    if(surf_break_name && region) {
                        handleBreakChange();
                        displayErrorMustSelect = false;
                    }
                    else{
                        displayErrorMustSelect = true
                    }     
                }}>
                    Get Conditions
                </button>
        
                {#if displayErrorMustSelect}
                    <div class="error">Must select break to search</div>
                {/if}
            </div>
        </div>

        <div class="live-conditions-container">
            <div class="live-conditions-header">Live Conditions</div>
                <div class="live-conditions">
                    <div class="grid-item swell">
                        <Swell {conditions}/>
                    </div>
                    <div class="grid-item wind">
                        <Wind {weather}/>
                    </div>
                    <div class="grid-item tide">
                        <Tide {tideExtremes} {localDate} {tideHeight}/>
					</div>
                    <div class="grid-item weather">
                        <Weather {weather} {conditions}/>
                    </div>
                </div>
        </div>

        <div class="forecast">
            <div class="forecast-header">
				Forecast
			</div>
			<div class="forecast-comps-container">
				<DailyForecast days_ahead=0 {localTimeConditionsMap} {tideExtremes}/>
				<DailyForecast days_ahead=1 {localTimeConditionsMap} {tideExtremes}/>
				<DailyForecast days_ahead=2 {localTimeConditionsMap} {tideExtremes}/>
				<DailyForecast days_ahead=3 {localTimeConditionsMap} {tideExtremes}/>
				<DailyForecast days_ahead=4 {localTimeConditionsMap} {tideExtremes}/>
				<DailyForecast days_ahead=5 {localTimeConditionsMap} {tideExtremes}/>
				<DailyForecast days_ahead=6 {localTimeConditionsMap} {tideExtremes}/>
			</div>
		</div>

    </div>
    {/if}
</main>

<style>

	*, *::before, *::after {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}
    .conditions-container {
		margin: 5em 1em;
        position: relative;
		color: #f0f0f0;
		display: flex;
		align-items: flex-start;
		justify-content: center;
	}
	
	.live-conditions {
		background-color: #313131;
		box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
		display: grid;
		grid-template-columns: repeat(3, 12em);
		grid-template-rows: repeat(2, 11em);
	}

	.grid-item {
		text-align: center;
		padding: 1rem;
	}

	.weather {
		grid-column-start: 2;
	}

	/* header styles */
	
	.conditions-header {
		width: 350px;
		margin-right: 1em;
		position: relative;
		background-color: #313131;
		box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
		font-family: Verdana, sans-serif;
		margin-bottom: 15px;
	}

	.break-name {
		font-size: 2.8em;
		line-height: 1.1em;
		font-family: Georgia;
		padding: 0.5rem 1rem 0.5rem 1rem;
		font-family: Verdana, sans-serif;
	}

	.break-location {
		font-size: 1.4em;
		padding: 0rem 1rem 0.5rem 1rem;
		font-family: Verdana, sans-serif;
	}

	.header-menu {
		position: relative;
		display: flex;
		width: 80%;
		margin: auto;
		padding-bottom: 1em;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.region-select-header, .break-select-header {
        background: #111111;
        color: rgb(240, 234, 234);
        width: 100%;
        padding: 6px;
        font-size: 1.5em;
        text-align: center;
		margin-top: 0.5em;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		transition-duration: 300ms;
    }

	.break-select-header {
		margin-bottom: 1em;
	}

	.search-button, .save-button {
        /* background-color: #6e38cc; */
		position: relative;
		background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        min-height: 5vh;
        max-width: 84%;
        text-align: center;
        font-size: 1.2em;
		font-family: verdana, sans-serif;
		box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
		border: none;
        border-radius: 5px;
        margin: 0 0 0.5em 0;
        cursor: pointer;
		transition-duration: 300ms;
		z-index: 1;
    }

	.search-button::before, .save-button::before{
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

	.search-button:hover::before, .save-button:hover::before {
		opacity: 1;
	}

	
	.region-select-header:hover, .break-select-header:hover {
		background-color: #030303;
	}

	button:hover {
        color: rgb(182, 182, 182);
    }
	
	.save-button {
		margin-top: 0.3em;
		margin-left: 0.8em;
	}

	.error {
        color: white;
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
    }

	select:defined {
        border: none;
        outline: none;
    }

	/* forecast styles */
	
	.forecast {
        position: relative;
		margin: 0 0 0 1em;
    }

	.forecast-header, .live-conditions-header {
		font-size: 3em;
		font-family: Verdana, sans-serif;
		text-align: center;
		background-color: #313131;
		box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
		padding: 0.5em 1em; 
		margin-bottom: 0.4em;
	}

	.forecast-comps-container {
		display: flex;
		flex-direction: row;
		max-width: 25vw;
		/* overflow-y: scroll; */
        overflow-y: hidden;
		box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
	}

    .loading-container {
		font-size: 2em;
		color: white;
		text-align: center;
		margin: 8em auto;
		transition-duration: 100ms;
		display: flex;
		flex-direction: row;
		justify-content: center;
		width: 100%;
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

	::-webkit-scrollbar {
  		height: 10px;
	}

	::-webkit-scrollbar-track {
		/* background: #313131; */
		background: #111111;
	}

	::-webkit-scrollbar-thumb {
		background: #888;
		height: 5px;
		border-radius: 5px;
	}

	
	::-webkit-scrollbar-thumb:hover {
		background: #555;
	} 
	

    @media (max-width: 500px) {
		.conditions-container {
			max-width: 90vw;
			margin: 4em auto 0 auto;
			flex-direction: column;
		}


		.live-conditions {
			grid-template-columns: repeat(2, 12em);
			grid-template-rows: repeat(2, 18em);
			margin-bottom: 1em;
		}

		.conditions-header {
			width: 90vw;
			margin: auto;
			padding-top: 0.5em;
			margin-bottom: 1em;
		}
		
		.break-name {
			text-align: center;
		}

		.break-location {
			text-align: center;
		}

		
		.header-menu {
			position: relative;
			justify-content: center;
			align-items: center;
			margin: auto;
			width: 70%;
			top: 0;
			right: 0;
			padding-bottom: 1em;
		}



		.forecast {
			margin: auto;
		}

		.forecast-comps-container {
			margin: auto;
			max-width: 90vw;
			margin-bottom: 1em;
		}
	}

	@media (max-width: 1280px) {
		.live-conditions {
			/* grid-template-columns: repeat(2, 12em); */
			grid-template-columns: auto 1fr auto;
			grid-template-rows: repeat(2, 13em);
		}

		.search-button {
			font-size: 1.2em;
		}
	}

	@media (max-width: 1150px) {
		.live-conditions {
			grid-template-columns: auto 1fr auto;
			grid-template-rows: repeat(2, 13em);
		}
	}

	@media (max-width: 1040px) {
		.live-conditions {
			grid-template-columns: repeat(2, 10em);
			grid-template-rows: repeat(2, 18em);
		}

		.forecast-header, .live-conditions-header {
			font-size: 2.2em;
		}

		.tide {
			grid-row: 1;
			grid-column: 2;
		}

		.wind {
			grid-column: 1;
		}
	}

	@media (max-width: 800px) {

		.conditions-container {
			max-width: 90vw;
			margin: 5em auto 0 auto;
			flex-direction: column;
		}


		.break-name {
			text-align: center;
		}

		.break-location {
			text-align: center;
		}

		.save-button-container {
			width: 100%;
			display: flex;
			justify-content: center;
		}

		.save-button {
			margin-left: 0;
		}

		.live-conditions {
			grid-template-columns: repeat(2, 1fr);
			grid-template-rows: repeat(2, 18em);
			margin-bottom: 1em;
			width: 90vw;
			margin: 0;
		}

		.conditions-container {
			flex-direction: column;
		}

		.forecast {
			margin: 1em 0;
		}


		.forecast-comps-container {
			max-width: 90vw;
			margin-bottom: 1em;
		}

		.conditions-header {
			width: 90vw;
			margin: auto;
			padding-top: 0.5em;
			margin-bottom: 1em;
		}
		
		.search-button {
			font-size: 1.2em;
		}

		.search-button:hover, .region-select-header:hover, .break-select-header:hover {
            outline: none;
        }

        .search-button:active {
            background-color: rgba(10, 5, 39, 0.336);
        }
	}

	@media (max-width: 450px) {
		.live-conditions {
			grid-template-columns: repeat(2, 1fr);
			grid-template-rows: repeat(2, 1fr);
		}
	}
</style>

