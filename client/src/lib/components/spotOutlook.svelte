<script lang="ts">
    import { onMount } from "svelte";
    import { conditionsParser } from "$lib/utils/conditionsParser"
    import { breaks } from "$lib/utils/breaks"
    import { _16point } from "$lib/utils/_16point"
    import { createEventDispatcher } from 'svelte';
    import { getTideHeight } from "$lib/utils/tideUtils";
    import axios from "axios"
    import { fly } from "svelte/transition";
    import { endpoint } from "$lib/utils/endpoint"
    import IconRefreshDouble from "~icons/iconoir/refresh-double"

    const dispatch = createEventDispatcher();  

    export let spotName;
    export let spotRegion;

    let conditionsMap;
    let conditions;
	let weather;

    let surfReportGenerated = false;
    let surfReportLoading = false;

    
    let surf_break_object;

    let swellHeight;
	let swellPeriod;
	let swellDirection;

    let windSpeed;
    let windDirection;

    let tideHeight;

    onMount(() => {
        getConditions(spotName, spotRegion)
    })

    async function getConditions(surf_break_name, region) {
        surfReportLoading = true;
        surfReportGenerated = false;
        surf_break_object = breaks.filter((br) => surf_break_name == br.break_name && region == br.region)[0];
		
		const lat = surf_break_object.lat;
		const lng = surf_break_object.lng;

        const d = new Date();
		// get local date to look up correct tides
		const unixUTCTime = Math.floor(new Date(d.toUTCString()).getTime() / 1000);
			
		

		

        try {
                // get time zone difference in seconds between local and utc
            const timeZoneResponse = await axios.get(endpoint + "/timezone/" + lat + "/" + lng)
            const utcOffset = timeZoneResponse.data.gmtOffset
                // convert local unix time to utc unix time
            const localTimeInUnix = unixUTCTime + utcOffset

            // need to get tides starting from yesterday to calculate
            // tide values that occur at times before the first extreme of the day
            const localDateYesterday = new Date((localTimeInUnix - 86400) * 1000).toISOString();


            const conditions_response = await axios.get(`${endpoint}/sg/${lat}/${lng}`);
            conditionsMap = conditionsParser(conditions_response.data);

            const weather_response = await axios.get(`${endpoint}/weather/${lat}/${lng}`);
            weather = weather_response.data;

            const tides_response = await axios.get(`${endpoint}/tide/${lat}/${lng}/${localDateYesterday.split("T")[0]}`);
			tideHeight = getTideHeight(tides_response.data, unixUTCTime)
        } 
        catch (error) {
            console.log(error)
        }

        const key = `${d.getUTCFullYear()}-${d.getUTCMonth() < 9 ? '0' : ''}${d.getUTCMonth() + 1}-${d.getUTCDate() < 10 ? '0' : ''}${d.getUTCDate()}T${d.getUTCHours() < 10 ? "0" : ""}${d.getUTCHours()}:00:00+00:00`
		
		conditions = conditionsMap.get(key)

		if(+conditions.swell_height > conditions.wind_wave_height){
            swellHeight = conditions.swell_height
            swellPeriod = conditions.swell_period
            swellDirection = conditions.swell_direction 
        }
        else {
            swellHeight = conditions.wind_wave_height
            swellPeriod = conditions.wind_wave_period
            swellDirection = conditions.wind_wave_direction 
        }

        windSpeed = weather.wind.speed.toFixed(0) + (+weather.wind.speed.toFixed(0) == 1 ? "kt" : "kts")
        windDirection = _16point(weather.wind.deg)
		
		surfReportLoading = false;
        surfReportGenerated = true;
    }
</script>

<main in:fly="{{ delay: 300, duration: 200 }}">
    <div class="header">
        <div class="name">{spotName}</div>
        <div class="region">{spotRegion}</div>
    </div>
    {#if surfReportGenerated}
        <div class="conditions-container" >
            <div class="swell">
                <div class="subheader">Swell</div>
                <div>{swellHeight}ft @ {swellPeriod}s<br>{_16point(+swellDirection)}</div>  
            </div>
            <div class="wind">
                <div class="subheader">Wind</div>
                <div>{windSpeed}<br>{windDirection}</div>
                
            </div>
            <div class="tide">
                <div class="subheader">Tide</div>
                <div>{tideHeight.height} ft.<br>{tideHeight.state}</div>
            </div>
        </div>

        <button type="button" on:click={() => {
            dispatch('fullReportRequest', {
                surf_break: spotName, region: spotRegion
            });
        }}>
            Full Report
        </button>
    {/if}
    {#if surfReportLoading}
        <div class="loading-container" in:fly="{{ duration: 200 }}">
            <div class="loading">Loading</div>
            <div class="loading-icon"><IconRefreshDouble style="margin: -9.8px; padding: 0;"/></div>
        </div>
    {/if}
</main>

<style>
    * {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

    main {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        background-color: #2f2f2f;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        color: #f0f0f0;
        padding: 0.5em 1em;
        width: 550px;
    }

    .header {
        width: 150px;
    }

    .name {
        font-size: 1.5em;
        line-height: 1.2em;
        font-family: Tahoma, sans-serif;
    }

    .region {
        font-size: 1em;
        font-family: Tahoma, sans-serif;
        margin: 0.25em 0;
    }

    .subheader {
        font-size: 1.2em;
        margin-bottom: 0.3em;
        font-family: Verdana, sans-serif;
    }

    .conditions-container {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }

    .swell, .wind, .tide {
        width: 80px;
        text-align: center;
        font-size: 1em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        height: 40px;
        width: 130px;
        text-align: center;
        font-size: 1em;
        font-family: Verdana, sans-serif;
        border: none;
        border-radius: 5px;
        margin: 0.5em 0;
        margin-left: 00.25em;
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

    
    .loading-container {
		font-size: 2em;
		color: white;
		text-align: center;
		transition-duration: 100ms;
		display: flex;
		flex-direction: row;
		justify-content: center;
		width: 100%;
        height: 100%;
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
        height: 1.2em;
	}

    @keyframes spin {
        from {
            transform:rotate(0deg);
        }
        to {
            transform:rotate(360deg);
        }
    }

    @media (max-width: 600px){
        main {
            display: grid;
            grid-template-rows: repeat(2, 4em);
            width: 90vw;
        }

        .header, .conditions-container{
            grid-row: 1;
        }

        button {
            place-self: end;
            grid-column: 2;
            margin-top: 1em;
        }

        .name {
            font-size: 1.3em;
        }

        .region {
            font-size: 0.9em;
        }

        .loading-container {
            grid-row: 1;
            margin-right: 1em;
        }
    }

    @media (max-width: 480px){
        .swell, .wind, .tide {
            width: 70px;
            font-size: 1em;
        }

        .header {
            width: 120px;
        }

        button {
            font-size: 1em;
            height: 40px;
            width: 130px;
        }
    }

    @media (max-width: 400px){
        .swell, .wind, .tide {
            width: 65px;
        }
    }

    @media (max-width: 370px){
        .header {
            width: 90px;
        }
        
        button {
            font-size: 0.9em;
            height: 35px;
            width: 110px;
        }
    }
</style>