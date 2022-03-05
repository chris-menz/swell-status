<script lang="ts">
    import { _16point } from '$lib/utils/_16point.js';
    import { convertMapToLocalTime } from '$lib/utils/convertMapKeysToLocalTime';
    import { getTideHeight } from '$lib/utils/tideUtils';
    import { onMount } from 'svelte';
    import { breaks, regions } from '$lib/utils/breaks';
    import axios from 'axios'
    import { fly } from "svelte/transition";
    import { getSurfSessions } from '$lib/stores/surfSessionsStore';
    import { axiosConfig} from "$lib/utils/axiosConfig"
    import { conditionsParser } from "$lib/utils/conditionsParser"
    import { time_options, convertAmPmToIsoTime } from "$lib/utils/dateTimeUtils"
    import { user } from "$lib/stores/userStore"
    import { endpoint } from "$lib/utils/endpoint"
    import IconRefreshDouble from "~icons/iconoir/refresh-double"


    // prevents user from getting conditions without inputting all parameters
    let isDisabled = true;
    $: isDisabled = !(surf_break_selected && region && date && time);

    let displayErrorBtnDisabled = false;

    // state
    let has_got_conditions = false;
    let conditionsLoading = false;

    // data for when and where session happened; used to fetch conditions for that time and location
    let surf_break_selected: String;
    let region: String;
    let date: Date;
    let time: String;
    let is_public: string;

    // updates when getPastConditions is called to reflect conditions at specific time and location that user requested
    let conditions;
    let tides;

    let swell;
    let wind;
    let tide;
    let surf_spot;
    let spot_region;
    let spot_location;

    // description of surf session will be inputted by user
    let description: String = '';

    // populates on Mount with last 7 dates
    let last7: Date[] = [];

    onMount(() => {
        last7 = last7Days();
    })

    function last7Days () {
        var result: Date[] = [];
        for (var i=0; i<7; i++) {
            var d = new Date();
            d.setDate(d.getDate() - i);
            result.push(d)
        }
        return(result);
    }

    async function addSurfSession(){
        const surf_session = {
            datetime: date.toDateString(),
            surf_spot,
            spot_region,
            spot_location,
            swell,
            wind,
            tide,
            description,
            is_public: is_public === "true"
        }

        const response = await axios.post(endpoint + "/auth/surfsession", surf_session, axiosConfig);
        getSurfSessions()

        description = ""
        has_got_conditions = false;
    }

    async function getPastConditions(date: Date, time: String) {
        conditionsLoading = true;
        has_got_conditions = false;
        
        // function must get data from 3 day range in order
        // to account for switch from local time to utc
        // example: a request for 02-12 6:00 pst is 02-11 23:00 in UTC

        let d = new Date(date)

        const startDate = new Date(d.setDate(d.getDate() - 1))
        const endDate = new Date(d.setDate(d.getDate() + 3))

        // converts Date object into string that fits requirement of the API
        const parsedStartDateString = `${startDate.getFullYear()}-${startDate.getMonth() < 9 ? '0' : ''}${startDate.getMonth() + 1}-${startDate.getDate() < 10 ? '0' : ''}${startDate.getDate()}`;
        const parsedEndDateString = `${endDate.getFullYear()}-${endDate.getMonth() < 9 ? '0' : ''}${endDate.getMonth() + 1}-${endDate.getDate() < 10 ? '0' : ''}${endDate.getDate()}`;
        const parsedDateString = `${date.getFullYear()}-${date.getMonth() < 9 ? '0' : ''}${date.getMonth() + 1}-${date.getDate() < 10 ? '0' : ''}${date.getDate()}`;
        

        // gets surf break object that matches user inputted name and region
        const surf_break_object = breaks.filter((br) => surf_break_selected == br.break_name && region == br.region)[0];

        const apiString = `${endpoint}/sg/${surf_break_object.lat}/${surf_break_object.lng}/${parsedStartDateString}/${parsedEndDateString}`;
        const response = await axios.get(apiString);
        
        const parsedConditions = conditionsParser(response.data)
        

        // convert local time at surf spot to utc time
        // because surf data from api is in UTC

        // combine date and time in ISO format
        const dateInIsoFormat = parsedDateString + convertAmPmToIsoTime(time)
    
        
        // get time zone difference in seconds between local and utc
        const timeZoneResponse = await axios.get(endpoint + "/timezone/" + surf_break_object.lat + "/" + surf_break_object.lng)
        const utcOffset = timeZoneResponse.data.gmtOffset
        
        // convert conditions to local time and assign to local conditions object
        conditions = convertMapToLocalTime(utcOffset, parsedConditions).get(dateInIsoFormat.slice(0, -6) + ".000Z")

        const tides_response = await axios.get(`${endpoint}/tide/${surf_break_object.lat}/${surf_break_object.lng}/${parsedStartDateString}`);

        const localTimeInUnix = Math.floor(new Date(new Date(dateInIsoFormat).toUTCString()).getTime() / 1000)
        const unixUTCTime = localTimeInUnix - utcOffset
        
        tides = getTideHeight(tides_response.data, unixUTCTime)


        // assign conditions to local variables
        swell = `${conditions.swell_height} ft. @ ${conditions.swell_period}s ${_16point(conditions.swell_direction)}`

        if(+conditions.secondary_swell_height > 1) {
            swell += `, ${conditions.secondary_swell_height} ft. @ ${conditions.secondary_swell_period}s ${_16point(conditions.secondary_swell_direction)}`
        }
        if(+conditions.wind_wave_height > 1) {
            swell += `, ${conditions.wind_wave_height} ft. @ ${conditions.wind_wave_period}s ${_16point(conditions.wind_wave_direction)}`
        }

        wind = `${(conditions.wind_speed).toFixed(0)}kts ${_16point(conditions.wind_direction)}`
        tide = `${tides.height} ft, ${tides.state}`

        surf_spot = surf_break_object.break_name
        spot_location = surf_break_object.location
        spot_region = surf_break_object.region

        console.log(spot_location)
        
        // update state
        has_got_conditions = true;   
        conditionsLoading = false; 
    }

    // surf_break_names dynamically updates based on chosen region, then displayed in dropdown menu
    let has_selected_region = false;
    let surf_break_names = []

    function handleRegionChange(selected) {
        surf_break_names = breaks.filter(surfBreak => surfBreak.region == selected).map(surfBreak => surfBreak.break_name); 
        has_selected_region = true
    }
</script>

<main>
    <div class="log-new-session-container">
        <div class="log-session-header header">Log a Session</div>
            <div class="new-session-form">
                
                <div class="select-container">
                    <div class="instructions">
                        First get the conditions, then add a description {#if !$user}<br><br>You are not logged in. You can still view past conditions, but you will not be able to save a session{/if}
                    </div>
                    <select class="date-select" bind:value={date}>
                        <option value="" disabled selected>Select Date</option>
                        {#each last7 as date_option}
                            <option value={date_option}>{date_option.toDateString()}</option>
                        {/each}
                    </select>
                    <select class="time-select" bind:value={time}>
                        <option value="" disabled selected>Select Time</option>
                        {#each time_options as time_option}
                            <option value={time_option}>{time_option}</option>
                        {/each}
                    </select> 
                    <select class="region-select" bind:value={region} on:change={() => handleRegionChange(region)}>
                        <option value="" disabled selected>Select Region</option>
                        {#each regions as region}
                            <option value="{region}">{region}</option>
                        {/each}
                    </select>
                    <select class="break-select" bind:value={surf_break_selected}>
                        <option value="" disabled selected>Select Break</option>
                        {#if has_selected_region}
                            {#each surf_break_names as break_option}
                                <option value="{break_option}">{break_option}</option>
                            {/each}
                        {/if}
                    </select>
                
                    <button type="button" class="get-conditions-btn" on:click={() => {
                        if(isDisabled){
                            displayErrorBtnDisabled = true;
                        }
                        else {
                            getPastConditions(date, time);
                            displayErrorBtnDisabled = false;
                        }   
                    }}>
                        Get Conditions
                    </button>
                
                    {#if displayErrorBtnDisabled}
                        <div class="btn-disabled-error">Must select all options to get conditions</div>
                    {/if}
                </div>

                {#if conditionsLoading}
                    <div class="loading-container" in:fly="{{ duration: 200 }}">
                        <div class="loading">Loading Surf Data</div>
                        <div class="loading-icon"><IconRefreshDouble style="margin: -9.8px; padding: 0;"/></div>
                    </div>
                {/if}
                
                {#if has_got_conditions}
                    <div class="session-info-container" in:fly="{{ delay: 350, duration: 200 }}">
                        <div class="session-info swell">
                            Swell: {swell}
                        </div>
                        <div class="session-info wind">
                            Wind: {wind}
                        </div>
                        <div class="session-info tide">
                            Tide: {tide}
                        </div>
                        <div class="session-info description-input-wrapper">
                            <textarea type="text" class="description-input" placeholder="Add a description" bind:value={description}></textarea>
                        </div>
                        <div>
                            <select bind:value={is_public}>
                                <option value="false">Private</option>
                                <option value="true">Public</option>
                            </select>
                            <div class="private-session-message">Private sessions are only viewable to you<br>Public sessions are viewable to anyone</div>
                        </div>
                        <div>
                            <button class="add-session-btn" on:click={() => addSurfSession()}>
                                Add Session
                            </button>
                        </div>
                    </div>   
                {/if}
            </div>
        </div>
</main>

<style>
    * {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}


    .header {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
       font-size: 3em;
        font-family: Verdana, sans-serif;
        padding: 0.5em 1em;
        margin: 0 0 0.5em 0;
        color: #f0f0f0;;
        text-align: center;
    }

    .log-new-session-container {
        display: flex;
        flex-direction: column;
        width: 100%;
        align-items: center;
    }

    .new-session-form {
        padding: 0 1em;
    }


    .select-container {
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px; 
        display: grid;
        grid-template-rows: repeat(2, 1fr);
        grid-template-columns: repeat(2, 14em);
        align-items: center;
        padding: 1em;
    }

    .instructions {
        color: white;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.1em;
        margin-bottom: 0.5em;
        /* margin-left: 1em;*/
        margin-right: 1em; 
        margin-top: 0.5em;
        margin-left: 0em;
        font-size: 1.2em;
        text-align: center;
        grid-row: 3;
    }

    select {
        font-size: 1.3em;
        max-width: 90%;
        margin: 0 0 0.5em 0;
        background-color: #1b1b1b;
        border: none;
        color: white;
        padding: 0.3em;
        border-radius: 5px;
        cursor: pointer;
        transition-duration: 300ms;
    }

    select:defined {
        outline: none;
    }

    select:hover {
        background-color: #030303;
    }

    .loading-container {
		font-size: 2em;
		color: white;
		text-align: center;
		margin: 3em auto;
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

    .session-info-container {
        background-color: #313131;
        padding: 1em;
        padding-right: 1em;
        margin-top: 1em;
        margin-bottom: 1em;
        color: white;
        max-width: 480px;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
    }

    .session-info {
        margin-bottom: 0.5em;
        font-size: 1.3em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
    }

    .description-input {
        width: 100%;
        padding: 0.4em;
        font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
        font-size: 1em;
        resize: none;
        background: #111111;
        color: white;
        border: none;
        border-radius: 5px;
        transition-duration: 300ms;
    }

    .description-input:defined {
        border: none;
        outline: none;
    }

    .description-input:hover {
        background: #030303;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        text-align: center;
        font-family: Verdana;
        font-size: 1.1em;
        color: rgb(240, 234, 234);
        max-width: 80%;
        padding: 0.5em 1em;
        border-radius: 5px;
        margin-top: 0.5em;
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

    .get-conditions-btn {
        align-self: flex-start;
        justify-self: center;
    }

    .btn-disabled-error {
        padding-top: 5px;
        color: white;
		font-family: Helvetica, sans-serif;
		font-weight: lighter;
        text-align: center;
    }

    .private-session-message {
        margin-bottom: 00.25em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        line-height: 1.5em;
    }

    @media (max-width: 1000px) {
        .header {
            display: none;
        }

        .instructions {
            grid-row: 3;
            margin-top: 1em;
            margin-left: 0em;
        }

        .session-info-container {
            max-width: none;
            width: 100%;
        }
    }

    @media (max-width: 700px) {
        
        .log-new-session-container {
            align-items: center;
            justify-content: center;
        }

        .select-container {
            width: 85vw;
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        select {
            text-align: center;
            width: 70%;
        }

        .session-info-container {
            width: 85vw;
            font-size: 1em;
        }

        .get-conditions-btn, .add-session-btn {
            font-size: 1.2em;
            max-width: 70%;
            margin-top: 0.5em;
            padding: 0.5em;
            place-self: center;
        }

        .instructions {
            margin-top: 0;
        }
    }
</style>