<script lang="ts">
	import { slide } from "svelte/transition";
	import { onMount } from "svelte";
	import {_16point} from "$lib/utils/_16point"
	import { convert24HrToAmPm } from "$lib/utils/dateTimeUtils"
	import IconArrowDown from "~icons/iconoir/arrow-down"

	export let days_ahead: string;
	export let localTimeConditionsMap: Map<string, Object>;
	export let tideExtremes;
	let tides_data;
	let tide_string_array = [];
	let only_3_tides = false;
	let conditionsHaveLoaded = false;
	
	// state of expanded table
	let isOpen = false;
	
	let weekdayIndex: number;
	
	const weekday = ["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"];
	const toggleTable = () => isOpen = !isOpen

	const d = new Date();

	let FiveAMConditions;
	let SixAMConditions;
	let SevenAMConditions;
	let EightAMConditions
	let NineAMConditions
	let TenAMConditions
	let ElevenAMConditions
	let TwelvePMConditions
	let OnePMConditions
	let TwoPMConditions
	let ThreePMConditions
	let FourPMConditions
	let FivePMConditions
	let SixPMConditions
	let SevenPMConditions
	let EightPMConditions

	onMount(() => {
		// gets index of correct weekday in array
		weekdayIndex = d.getUTCDay() + +days_ahead;
		if(weekdayIndex > 6){weekdayIndex -= 7;}
		
		// extracts tide data from correct day
		const date = new Date()
		date.setUTCDate(date.getUTCDate() + +days_ahead)
		tides_data = tideExtremes.get(date.toISOString().split("T")[0])

		// format tide data for display
		for(let i = 0; i < 4; ++i){
			if(tides_data[i]){
				const tide = tides_data[i]
				const time = convert24HrToAmPm(tide.date.substring(11, 16))
				tide_string_array[i] = `${tide.type}: ${Math.round(tide.height * 3.281 * 10) / 10} ft, ${time}`;
			}
			else{
				only_3_tides = true;
			}	
		}	

		const formattedDate = date.toISOString().split("T")[0]

		FiveAMConditions = localTimeConditionsMap.get(formattedDate + "T05:00:00.000Z")
		SixAMConditions = localTimeConditionsMap.get(formattedDate + "T06:00:00.000Z")
		SevenAMConditions = localTimeConditionsMap.get(formattedDate + "T07:00:00.000Z")
		EightAMConditions = localTimeConditionsMap.get(formattedDate + "T08:00:00.000Z")
		NineAMConditions = localTimeConditionsMap.get(formattedDate + "T09:00:00.000Z")
		TenAMConditions = localTimeConditionsMap.get(formattedDate + "T10:00:00.000Z")
		ElevenAMConditions = localTimeConditionsMap.get(formattedDate + "T11:00:00.000Z")
		TwelvePMConditions = localTimeConditionsMap.get(formattedDate + "T12:00:00.000Z")
		OnePMConditions = localTimeConditionsMap.get(formattedDate + "T13:00:00.000Z")
		TwoPMConditions = localTimeConditionsMap.get(formattedDate + "T14:00:00.000Z")
		ThreePMConditions = localTimeConditionsMap.get(formattedDate + "T15:00:00.000Z")
		FourPMConditions = localTimeConditionsMap.get(formattedDate + "T16:00:00.000Z")
		FivePMConditions = localTimeConditionsMap.get(formattedDate + "T17:00:00.000Z")
		SixPMConditions = localTimeConditionsMap.get(formattedDate + "T18:00:00.000Z")
		SevenPMConditions = localTimeConditionsMap.get(formattedDate + "T19:00:00.000Z")
		EightPMConditions = localTimeConditionsMap.get(formattedDate + "T20:00:00.000Z")

		conditionsHaveLoaded = true
	})
	

</script>
<main>
	{#if conditionsHaveLoaded}
	<div class="small-table table">
		<div class="date">
			{weekday[weekdayIndex]}
			{#if +days_ahead == 0}
				<div class="arrow"><IconArrowDown /></div>
			{/if}
		</div>
		
		
			<div class="forecast-data">
				<table>
					<tr>
						<th></th>
						<th>Swell</th>
						<th>Swell 2</th>
						<th>Wind</th>
					</tr>
					<tr>	
						<td>AM</td>
						<td>{EightAMConditions.swell_height} ft. @ {EightAMConditions.swell_period}s<br>{EightAMConditions.swell_direction}° {_16point(EightAMConditions.swell_direction)}</td>
						<td>
						{#if +EightAMConditions.secondary_swell_height >= +EightAMConditions.wind_wave_height && +EightAMConditions.secondary_swell_height > 1}
							{EightAMConditions.secondary_swell_height} ft. @ {EightAMConditions.secondary_swell_period}s<br>{EightAMConditions.secondary_swell_direction}° {_16point(EightAMConditions.secondary_swell_direction)}
						{/if}
						{#if +EightAMConditions.secondary_swell_height < +EightAMConditions.wind_wave_height && +EightAMConditions.wind_wave_height > 1}
							{EightAMConditions.wind_wave_height} ft. @ {EightAMConditions.wind_wave_period}s<br>{EightAMConditions.wind_wave_direction}° {_16point(EightAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(EightAMConditions.wind_speed).toFixed(0)} {Math.round(EightAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(EightAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>PM</td>
						<td>{FourPMConditions.swell_height} ft. @ {FourPMConditions.swell_period}s<br>{FourPMConditions.swell_direction}° {_16point(FourPMConditions.swell_direction)}</td>
						<td>
						{#if +FourPMConditions.secondary_swell_height >= +FourPMConditions.wind_wave_height && +FourPMConditions.secondary_swell_height > 1}
							{FourPMConditions.secondary_swell_height} ft. @ {FourPMConditions.secondary_swell_period}s<br>{FourPMConditions.secondary_swell_direction}° {_16point(FourPMConditions.secondary_swell_direction)}
						{/if}
						{#if +FourPMConditions.secondary_swell_height < +FourPMConditions.wind_wave_height && +FourPMConditions.wind_wave_height > 1}
							{FourPMConditions.wind_wave_height} ft. @ {FourPMConditions.wind_wave_period}s<br>{FourPMConditions.wind_wave_direction}° {_16point(FourPMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(FourPMConditions.wind_speed).toFixed(0)} {Math.round(FourPMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(FourPMConditions.wind_direction)}</td>
					</tr>
				</table>
				
			</div>
			
			<div class="tides-container">
				<div class="tides-header">Tides</div>
				{#each tide_string_array as tide}
					<div class="tide">{tide}</div>	
				{/each}
				{#if only_3_tides}
					<div class="tide"><br></div>
				{/if}
			</div>
			
			<button class="expand-button" on:click={() => toggleTable()} aria-expanded={isOpen}>
				Hourly Conditions
			</button>
			
		</div>
		{#if isOpen}
			
			<div class="expanded-table table" transition:slide="{{ duration: 400 }}">
				
				<table >
					<tr>
						<th></th>
						<th>Swell</th>
						<th>Swell 2</th>
						<th>Wind</th>
					</tr>
					<tr>	
						<td>5:00 AM</td>
						<td>{FiveAMConditions.swell_height} ft. @ {FiveAMConditions.swell_period}s<br>{FiveAMConditions.swell_direction}° {_16point(FiveAMConditions.swell_direction)}</td>
						<td>
						{#if +FiveAMConditions.secondary_swell_height >= +FiveAMConditions.wind_wave_height && +FiveAMConditions.secondary_swell_height > 1}
							{FiveAMConditions.secondary_swell_height} ft. @ {FiveAMConditions.secondary_swell_period}s<br>{FiveAMConditions.secondary_swell_direction}° {_16point(FiveAMConditions.secondary_swell_direction)}
						{/if}
						{#if +FiveAMConditions.secondary_swell_height < +FiveAMConditions.wind_wave_height && +FiveAMConditions.wind_wave_height > 1}
							{FiveAMConditions.wind_wave_height} ft. @ {FiveAMConditions.wind_wave_period}s<br>{FiveAMConditions.wind_wave_direction}° {_16point(FiveAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(FiveAMConditions.wind_speed).toFixed(0)} {Math.round(FiveAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(FiveAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>6:00 AM</td>
						<td>{SixAMConditions.swell_height} ft. @ {SixAMConditions.swell_period}s<br>{SixAMConditions.swell_direction}° {_16point(SixAMConditions.swell_direction)}</td>
						<td>
						{#if +SixAMConditions.secondary_swell_height >= +SixAMConditions.wind_wave_height && +SixAMConditions.secondary_swell_height > 1}
							{SixAMConditions.secondary_swell_height} ft. @ {SixAMConditions.secondary_swell_period}s<br>{SixAMConditions.secondary_swell_direction}° {_16point(SixAMConditions.secondary_swell_direction)}
						{/if}
						{#if +SixAMConditions.secondary_swell_height < +SixAMConditions.wind_wave_height && +SixAMConditions.wind_wave_height > 1}
							{SixAMConditions.wind_wave_height} ft. @ {SixAMConditions.wind_wave_period}s<br>{SixAMConditions.wind_wave_direction}° {_16point(SixAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(SixAMConditions.wind_speed).toFixed(0)} {Math.round(SixAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(SixAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>7:00 AM</td>
						<td>{SevenAMConditions.swell_height} ft. @ {SevenAMConditions.swell_period}s<br>{SevenAMConditions.swell_direction}° {_16point(SevenAMConditions.swell_direction)}</td>
						{#if SevenAMConditions.secondary_swell_height >= SevenAMConditions.wind_wave_height}
							<td>{SevenAMConditions.secondary_swell_height} ft. @ {SevenAMConditions.secondary_swell_period}s<br>{SevenAMConditions.secondary_swell_direction}° {_16point(SevenAMConditions.secondary_swell_direction)}</td>
						{/if}
						{#if SevenAMConditions.secondary_swell_height < SevenAMConditions.wind_wave_height}
							<td>{SevenAMConditions.wind_wave_height} ft. @ {SevenAMConditions.wind_wave_period}s<br>{SevenAMConditions.wind_wave_direction}° {_16point(SevenAMConditions.wind_wave_direction)}</td>
						{/if}
						<td>{(SevenAMConditions.wind_speed).toFixed(0)} {Math.round(SevenAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(SevenAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>8:00 AM</td>
						<td>{EightAMConditions.swell_height} ft. @ {EightAMConditions.swell_period}s<br>{EightAMConditions.swell_direction}° {_16point(EightAMConditions.swell_direction)}</td>
						<td>
						{#if +EightAMConditions.secondary_swell_height >= +EightAMConditions.wind_wave_height && +EightAMConditions.secondary_swell_height > 1}
							{EightAMConditions.secondary_swell_height} ft. @ {EightAMConditions.secondary_swell_period}s<br>{EightAMConditions.secondary_swell_direction}° {_16point(EightAMConditions.secondary_swell_direction)}
						{/if}
						{#if +EightAMConditions.secondary_swell_height < +EightAMConditions.wind_wave_height && +EightAMConditions.wind_wave_height > 1}
							{EightAMConditions.wind_wave_height} ft. @ {EightAMConditions.wind_wave_period}s<br>{EightAMConditions.wind_wave_direction}° {_16point(EightAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(EightAMConditions.wind_speed).toFixed(0)} {Math.round(EightAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(EightAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>9:00 AM</td>
						<td>{NineAMConditions.swell_height} ft. @ {NineAMConditions.swell_period}s<br>{NineAMConditions.swell_direction}° {_16point(NineAMConditions.swell_direction)}</td>
						<td>
						{#if +NineAMConditions.secondary_swell_height >= +NineAMConditions.wind_wave_height && +NineAMConditions.secondary_swell_height > 1}
							{NineAMConditions.secondary_swell_height} ft. @ {NineAMConditions.secondary_swell_period}s<br>{NineAMConditions.secondary_swell_direction}° {_16point(NineAMConditions.secondary_swell_direction)}
						{/if}
						{#if +NineAMConditions.secondary_swell_height < +NineAMConditions.wind_wave_height && +NineAMConditions.wind_wave_height > 1}
							{NineAMConditions.wind_wave_height} ft. @ {NineAMConditions.wind_wave_period}s<br>{NineAMConditions.wind_wave_direction}° {_16point(NineAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(NineAMConditions.wind_speed).toFixed(0)} {Math.round(NineAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(NineAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>10:00 AM</td>
						<td>{TenAMConditions.swell_height} ft. @ {TenAMConditions.swell_period}s<br>{TenAMConditions.swell_direction}° {_16point(TenAMConditions.swell_direction)}</td>
						<td>
						{#if +TenAMConditions.secondary_swell_height >= +TenAMConditions.wind_wave_height && +TenAMConditions.secondary_swell_height > 1}
							{TenAMConditions.secondary_swell_height} ft. @ {TenAMConditions.secondary_swell_period}s<br>{TenAMConditions.secondary_swell_direction}° {_16point(TenAMConditions.secondary_swell_direction)}
						{/if}
						{#if +TenAMConditions.secondary_swell_height < +TenAMConditions.wind_wave_height && +TenAMConditions.wind_wave_height > 1}
							{TenAMConditions.wind_wave_height} ft. @ {TenAMConditions.wind_wave_period}s<br>{TenAMConditions.wind_wave_direction}° {_16point(TenAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(TenAMConditions.wind_speed).toFixed(0)} {Math.round(TenAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(TenAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>11:00 AM</td>
						<td>{ElevenAMConditions.swell_height} ft. @ {ElevenAMConditions.swell_period}s<br>{ElevenAMConditions.swell_direction}° {_16point(ElevenAMConditions.swell_direction)}</td>
						<td>
						{#if +ElevenAMConditions.secondary_swell_height >= +ElevenAMConditions.wind_wave_height && +ElevenAMConditions.secondary_swell_height > 1}
							{ElevenAMConditions.secondary_swell_height} ft. @ {ElevenAMConditions.secondary_swell_period}s<br>{ElevenAMConditions.secondary_swell_direction}° {_16point(ElevenAMConditions.secondary_swell_direction)}
						{/if}
						{#if +ElevenAMConditions.secondary_swell_height < +ElevenAMConditions.wind_wave_height && +ElevenAMConditions.wind_wave_height > 1}
							{ElevenAMConditions.wind_wave_height} ft. @ {ElevenAMConditions.wind_wave_period}s<br>{ElevenAMConditions.wind_wave_direction}° {_16point(ElevenAMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(ElevenAMConditions.wind_speed).toFixed(0)} {Math.round(ElevenAMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(ElevenAMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>12:00 PM</td>
						<td>{TwelvePMConditions.swell_height} ft. @ {TwelvePMConditions.swell_period}s<br>{TwelvePMConditions.swell_direction}° {_16point(TwelvePMConditions.swell_direction)}</td>
						<td>
						{#if +TwelvePMConditions.secondary_swell_height >= +TwelvePMConditions.wind_wave_height && +TwelvePMConditions.secondary_swell_height > 1}
							{TwelvePMConditions.secondary_swell_height} ft. @ {TwelvePMConditions.secondary_swell_period}s<br>{TwelvePMConditions.secondary_swell_direction}° {_16point(TwelvePMConditions.secondary_swell_direction)}
						{/if}
						{#if +TwelvePMConditions.secondary_swell_height < +TwelvePMConditions.wind_wave_height && +TwelvePMConditions.wind_wave_height > 1}
							{TwelvePMConditions.wind_wave_height} ft. @ {TwelvePMConditions.wind_wave_period}s<br>{TwelvePMConditions.wind_wave_direction}° {_16point(TwelvePMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(TwelvePMConditions.wind_speed).toFixed(0)} {Math.round(TwelvePMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(TwelvePMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>1:00 PM</td>
						<td>{OnePMConditions.swell_height} ft. @ {OnePMConditions.swell_period}s<br>{OnePMConditions.swell_direction}° {_16point(OnePMConditions.swell_direction)}</td>
						<td>
						{#if +OnePMConditions.secondary_swell_height >= +OnePMConditions.wind_wave_height && +OnePMConditions.secondary_swell_height > 1}
							{OnePMConditions.secondary_swell_height} ft. @ {OnePMConditions.secondary_swell_period}s<br>{OnePMConditions.secondary_swell_direction}° {_16point(OnePMConditions.secondary_swell_direction)}
						{/if}
						{#if +OnePMConditions.secondary_swell_height < +OnePMConditions.wind_wave_height && +OnePMConditions.wind_wave_height > 1}
							{OnePMConditions.wind_wave_height} ft. @ {OnePMConditions.wind_wave_period}s<br>{OnePMConditions.wind_wave_direction}° {_16point(OnePMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(OnePMConditions.wind_speed).toFixed(0)} {Math.round(OnePMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(OnePMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>2:00 PM</td>
						<td>{TwoPMConditions.swell_height} ft. @ {TwoPMConditions.swell_period}s<br>{TwoPMConditions.swell_direction}° {_16point(TwoPMConditions.swell_direction)}</td>
						<td>
						{#if +TwoPMConditions.secondary_swell_height >= +TwoPMConditions.wind_wave_height && +TwoPMConditions.secondary_swell_height > 1}
							{TwoPMConditions.secondary_swell_height} ft. @ {TwoPMConditions.secondary_swell_period}s<br>{TwoPMConditions.secondary_swell_direction}° {_16point(TwoPMConditions.secondary_swell_direction)}
						{/if}
						{#if +TwoPMConditions.secondary_swell_height < +TwoPMConditions.wind_wave_height && +TwoPMConditions.wind_wave_height > 1}
							{TwoPMConditions.wind_wave_height} ft. @ {TwoPMConditions.wind_wave_period}s<br>{TwoPMConditions.wind_wave_direction}° {_16point(TwoPMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(TwoPMConditions.wind_speed).toFixed(0)} {Math.round(TwoPMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(TwoPMConditions.wind_direction)}</td>
					</tr>
					<tr>	
						<td>3:00 PM</td>
						<td>{ThreePMConditions.swell_height} ft. @ {ThreePMConditions.swell_period}s<br>{ThreePMConditions.swell_direction}° {_16point(ThreePMConditions.swell_direction)}</td>
						<td>
						{#if +ThreePMConditions.secondary_swell_height >= +ThreePMConditions.wind_wave_height && +ThreePMConditions.secondary_swell_height > 1}
							{ThreePMConditions.secondary_swell_height} ft. @ {ThreePMConditions.secondary_swell_period}s<br>{ThreePMConditions.secondary_swell_direction}° {_16point(ThreePMConditions.secondary_swell_direction)}
						{/if}
						{#if +ThreePMConditions.secondary_swell_height < +ThreePMConditions.wind_wave_height && +ThreePMConditions.wind_wave_height > 1}
							{ThreePMConditions.wind_wave_height} ft. @ {ThreePMConditions.wind_wave_period}s<br>{ThreePMConditions.wind_wave_direction}° {_16point(ThreePMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(ThreePMConditions.wind_speed).toFixed(0)} {Math.round(ThreePMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(ThreePMConditions.wind_direction)}</td>			
					</tr>
					<tr>	
						<td>4:00 PM</td>
						<td>{FourPMConditions.swell_height} ft. @ {FourPMConditions.swell_period}s<br>{FourPMConditions.swell_direction}° {_16point(FourPMConditions.swell_direction)}</td>
						<td>
						{#if +FourPMConditions.secondary_swell_height >= +FourPMConditions.wind_wave_height && +FourPMConditions.secondary_swell_height > 1}
							{FourPMConditions.secondary_swell_height} ft. @ {FourPMConditions.secondary_swell_period}s<br>{FourPMConditions.secondary_swell_direction}° {_16point(FourPMConditions.secondary_swell_direction)}
						{/if}
						{#if +FourPMConditions.secondary_swell_height < +FourPMConditions.wind_wave_height && +FourPMConditions.wind_wave_height > 1}
							{FourPMConditions.wind_wave_height} ft. @ {FourPMConditions.wind_wave_period}s<br>{FourPMConditions.wind_wave_direction}° {_16point(FourPMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(FourPMConditions.wind_speed).toFixed(0)} {Math.round(FourPMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(FourPMConditions.wind_direction)}</td>				
					</tr>
					<tr>	
						<td>5:00 PM</td>
						<td>{FivePMConditions.swell_height} ft. @ {FivePMConditions.swell_period}s<br>{FivePMConditions.swell_direction}° {_16point(FivePMConditions.swell_direction)}</td>
						<td>
						{#if +FivePMConditions.secondary_swell_height >= +FivePMConditions.wind_wave_height && +FivePMConditions.secondary_swell_height > 1}
							{FivePMConditions.secondary_swell_height} ft. @ {FivePMConditions.secondary_swell_period}s<br>{FivePMConditions.secondary_swell_direction}° {_16point(FivePMConditions.secondary_swell_direction)}
						{/if}
						{#if +FivePMConditions.secondary_swell_height < +FivePMConditions.wind_wave_height && +FivePMConditions.wind_wave_height > 1}
							{FivePMConditions.wind_wave_height} ft. @ {FivePMConditions.wind_wave_period}s<br>{FivePMConditions.wind_wave_direction}° {_16point(FivePMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(FivePMConditions.wind_speed).toFixed(0)} {Math.round(FivePMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(FivePMConditions.wind_direction)}</td>		
					</tr>
					<tr>	
						<td>6:00 PM</td>
						<td>{SixPMConditions.swell_height} ft. @ {SixPMConditions.swell_period}s<br>{SixPMConditions.swell_direction}° {_16point(SixPMConditions.swell_direction)}</td>
						<td>
						{#if +SixPMConditions.secondary_swell_height >= +SixPMConditions.wind_wave_height && +SixPMConditions.secondary_swell_height > 1}
							{SixPMConditions.secondary_swell_height} ft. @ {SixPMConditions.secondary_swell_period}s<br>{SixPMConditions.secondary_swell_direction}° {_16point(SixPMConditions.secondary_swell_direction)}
						{/if}
						{#if +SixPMConditions.secondary_swell_height < +SixPMConditions.wind_wave_height && +SixPMConditions.wind_wave_height > 1}
							{SixPMConditions.wind_wave_height} ft. @ {SixPMConditions.wind_wave_period}s<br>{SixPMConditions.wind_wave_direction}° {_16point(SixPMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(SixPMConditions.wind_speed).toFixed(0)} {Math.round(SixPMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(SixPMConditions.wind_direction)}</td>			
					</tr>
					<tr>	
						<td>7:00 PM</td>
						<td>{SevenPMConditions.swell_height} ft. @ {SevenPMConditions.swell_period}s<br>{SevenPMConditions.swell_direction}° {_16point(SevenPMConditions.swell_direction)}</td>
						<td>
						{#if +SevenPMConditions.secondary_swell_height >= +SevenPMConditions.wind_wave_height && +SevenPMConditions.secondary_swell_height > 1}
							{SevenPMConditions.secondary_swell_height} ft. @ {SevenPMConditions.secondary_swell_period}s<br>{SevenPMConditions.secondary_swell_direction}° {_16point(SevenPMConditions.secondary_swell_direction)}
						{/if}
						{#if +SevenPMConditions.secondary_swell_height < +SevenPMConditions.wind_wave_height && +SevenPMConditions.wind_wave_height > 1}
							{SevenPMConditions.wind_wave_height} ft. @ {SevenPMConditions.wind_wave_period}s<br>{SevenPMConditions.wind_wave_direction}° {_16point(SevenPMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td>{(SevenPMConditions.wind_speed).toFixed(0)} {Math.round(SevenPMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(SevenPMConditions.wind_direction)}</td>				
					</tr>
					<tr>	
						<td style="border-bottom: none;">8:00 PM</td>
						<td style="border-bottom: none;">{EightPMConditions.swell_height} ft. @ {EightPMConditions.swell_period}s<br>{EightPMConditions.swell_direction}° {_16point(EightPMConditions.swell_direction)}</td>
						<td style="border-bottom: none;">
						{#if +EightPMConditions.secondary_swell_height >= +EightPMConditions.wind_wave_height && +EightPMConditions.secondary_swell_height > 1}
							{EightPMConditions.secondary_swell_height} ft. @ {EightPMConditions.secondary_swell_period}s<br>{EightPMConditions.secondary_swell_direction}° {_16point(EightPMConditions.secondary_swell_direction)}
						{/if}
						{#if +EightPMConditions.secondary_swell_height < +EightPMConditions.wind_wave_height && +EightPMConditions.wind_wave_height > 1}
							{EightPMConditions.wind_wave_height} ft. @ {EightPMConditions.wind_wave_period}s<br>{EightPMConditions.wind_wave_direction}° {_16point(EightPMConditions.wind_wave_direction)}
						{/if}
						</td>
						<td style="border-bottom: none;">{(EightPMConditions.wind_speed).toFixed(0)} {Math.round(EightPMConditions.wind_speed) == 1 ? "kt" : "kts"}<br>{_16point(EightPMConditions.wind_direction)}</td>				
					</tr>
				</table>
			</div>
		{/if}
	{/if}
</main>

<style>
	* {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
		text-align: center;
	}

	main {
		margin: 0 1em 0 0;
	}

	th {
		font-size: 1em;
		font-family: Verdana;
		font-weight: normal;
	}

	.table {
		background-color: #313131;
	}

	.small-table, .expanded-table {
		width: 330px;
	}
	
	.small-table table{
		border-top: 2px solid grey;
	}

	.expanded-table {
		border-top: 2px solid grey;
	}

	.forecast-data {
		display: flex;
		flex-direction: row;
	}

	table {
		margin: auto;
		padding: 0.5em;
	}

	.date {
		position: relative;
		padding: 0.5em 0;
		font-size: 1.5em;
		font-family: Verdana, sans-serif;
		text-shadow: 1px 1px 2px #000000;
	}

	.arrow {
		position: absolute;
		right: 4px;
		top: 8px;
		transform: rotate(270deg);
	}


	th {
		font-family: Verdana, sans-serif;
		padding-bottom: 00.2em;
	}

	td {
		padding: 0.2em 0.3em;
		font-family: Helvetica, sans-serif;
		font-weight:lighter;
	}

	.expanded-table td {
		border-bottom: 1px solid grey;
	}


	.expand-button {
		margin: 1em auto;
		padding: 0.5em;
		position: relative;
		background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
		font-family: verdana;
		width: 60%;
		text-align: center;
        font-size: 1em;
        border: none;
        border-radius: 5px;
		cursor: pointer;
		box-shadow: rgba(0, 0, 0, 0.07) 0px 1px 2px, rgba(0, 0, 0, 0.07) 0px 2px 4px, rgba(0, 0, 0, 0.07) 0px 4px 8px, rgba(0, 0, 0, 0.07) 0px 8px 16px, rgba(0, 0, 0, 0.07) 0px 16px 32px, rgba(0, 0, 0, 0.07) 0px 32px 64px;
		z-index: 1;
		transition-duration: 200ms;
	}

	.expand-button::before{
		position: absolute;
		content: "";
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		border-radius: 5px;
		background: linear-gradient(to right bottom, rgb(59, 31, 110),rgb(57, 39, 122)0);
		transition: opacity 0.2s;
		z-index: -1;
		opacity: 0;
	}

	.expand-button:hover::before{
		opacity: 1;
	}

	button:hover {
        color: rgb(182, 182, 182);
    }

	.tide {
		line-height: 1.3em;
		font-family: Helvetica, sans-serif;
		font-weight:lighter;
	}

	.tides-header {
		padding: 0 0 0.3em 0;
		font-size: 1.2em;
		font-family: Verdana, sans-serif;
		text-shadow: 1px 1px 2px #000000;
	}
	
</style>