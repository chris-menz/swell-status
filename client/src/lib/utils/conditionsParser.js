export function conditionsParser(response) {
  let conditions = new Map();

  const conditions_data = response.hours;

  for (let i = 0; i < conditions_data.length; ++i) {
    conditions.set(conditions_data[i].time, {
      swell_height: (conditions_data[i].swellHeight.sg * 3.281).toFixed(1),
      swell_period: conditions_data[i].swellPeriod.noaa.toFixed(0),
      swell_direction: conditions_data[i].swellDirection.sg.toFixed(0),
      wind_speed: conditions_data[i].windSpeed.sg * (1.98).toFixed(0),
      wind_direction: conditions_data[i].windDirection.sg.toFixed(0),
      water_temperature: (
        (conditions_data[i].waterTemperature.sg * 9) / 5 +
        32
      ).toFixed(0),
      wave_height: (conditions_data[i].waveHeight.sg * 3.281).toFixed(1),
      wave_period: conditions_data[i].wavePeriod.sg.toFixed(0),
      wave_direction: conditions_data[i].waveDirection.sg.toFixed(0),
      wind_wave_height: (conditions_data[i].windWaveHeight.sg * 3.281).toFixed(
        1
      ),
      wind_wave_period: conditions_data[i].windWavePeriod.sg.toFixed(0),
      wind_wave_direction: conditions_data[i].windWaveDirection.sg.toFixed(0),
      secondary_swell_height: (
        conditions_data[i].secondarySwellHeight.sg * 3.281
      ).toFixed(1),
      secondary_swell_period:
        conditions_data[i].secondarySwellPeriod.sg.toFixed(0),
      secondary_swell_direction:
        conditions_data[i].secondarySwellDirection.sg.toFixed(0),
    });
  }
  return conditions;
}
