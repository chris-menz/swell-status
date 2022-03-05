// takes a time zone offset and a map of conditions with dates as keys
// and converts keys from utc to local time
export function convertMapToLocalTime(utcOffset, conditionsMap) {
  let localTimeConditionsMap = new Map();
  for (let [key, value] of conditionsMap) {
    const unixUTCTime = Math.floor(new Date(key).getTime() / 1000);
    const localTimeInUnix = unixUTCTime + utcOffset;
    const localDate = new Date(localTimeInUnix * 1000).toISOString();
    localTimeConditionsMap.set(localDate, value);
  }
  return localTimeConditionsMap;
}
