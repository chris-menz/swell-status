// tide parser takes in array of ~28 tide extremes for the next 7 days
// and returns a map which has a key of the date in YYYY-MM-DD
// and value of an array of tides for that day.

// this allows easy lookup for the 7 day forecast and current day conditions

export function tideParser(tideData: TideModel) {
  let tideMap = new Map<string, [Extreme]>();

  for (let i = 0; i < tideData.extremes.length; ++i) {
    const dateSubstring = tideData.extremes[i].date.substring(0, 10);

    if (tideMap.has(dateSubstring)) {
      let extremesArr = tideMap.get(dateSubstring);
      extremesArr.push(tideData.extremes[i]);
      tideMap.delete(dateSubstring);
      tideMap.set(dateSubstring, extremesArr);
    } else {
      let extremesArr: [Extreme] = [tideData.extremes[i]];
      tideMap.set(dateSubstring, extremesArr);
    }
  }

  return tideMap;
}

// takes in array of tide extremes and a time and
// returns the estimated tide height at that time

// done by getting two surrounding extremes

export function getTideHeight(tideData: TideModel, utcTimeInUnix: number) {
  // get the two surrounding extremes. .dt is the time in unix for that extreme
  let extremeOne;
  let extremeTwo;

  let heightOne;
  let heightTwo;

  for (let i = 0; i < tideData.extremes.length; ++i) {
    if (tideData.extremes[i].dt > utcTimeInUnix) {
      extremeOne = tideData.extremes[i - 1];
      extremeTwo = tideData.extremes[i];
      heightOne = Math.round(extremeOne.height * 3.281 * 10) / 10;
      heightTwo = Math.round(extremeTwo.height * 3.281 * 10) / 10;

      // break loop
      i = 10000;
    }
  }

  // create trig equation from two surrounding extremes that gets tide height in feet
  // given a time in unix

  const period = (extremeTwo.dt - extremeOne.dt) * 2;
  const a = (heightOne - heightTwo) / 2;
  const b = extremeOne.dt;
  const c =
    Math.abs((heightOne - heightTwo) / 2) +
    (heightOne > heightTwo ? heightTwo : heightOne);
  const d = (2 * Math.PI) / period;

  const state = extremeTwo.type == "High" ? "Incoming" : "Outgoing";

  // compute and convert from meters to feet

  return {
    height: +(
      c +
      a *
        (Math.sqrt(Math.abs(Math.cos(d * (utcTimeInUnix - b)))) *
          Math.sign(Math.cos(d * (utcTimeInUnix - b))))
    ).toFixed(1),
    state,
  };

  // return {
  //   height: a * Math.cos(d * (utcTimeInUnix - b)) + c,
  //   state,
  // };

  //   return +(
  //     (c +
  //       a *
  //         (Math.sqrt(Math.abs(Math.cos(d * (utcTimeInUnix - b)))) *
  //           Math.sign(Math.cos(d * (utcTimeInUnix - b))))) *
  //     3.281
  //   ).toFixed(1);

  // return {
  //   height: +(
  //     ((c +
  //       a *
  //         (Math.sqrt(Math.abs(Math.cos(d * (utcTimeInUnix - b)))) *
  //           Math.sign(Math.cos(d * (utcTimeInUnix - b))))) *
  //       3.281 +
  //       (a * Math.cos(d * (utcTimeInUnix - b)) + c) * 3.281) /
  //     2
  //   ).toFixed(1),
  //   state,
  // };
}

interface TideModel {
  status: number;
  callCount: number;
  copyright: string;
  requestLat: number;
  requestLon: number;
  responseLat: number;
  responseLon: number;
  atlas: string;
  station: string;
  timezone: string;
  requestDatum: string;
  responseDatum: string;
  extremes: Extreme[];
}

interface Extreme {
  dt: number;
  date: string;
  height: number;
  type: string;
}
