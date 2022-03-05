export const time_options = [
  "5:00 AM",
  "5:30 AM",
  "6:00 AM",
  "6:30 AM",
  "7:00 AM",
  "7:30 AM",
  "8:00 AM",
  "8:30 AM",
  "9:00 AM",
  "9:30 AM",
  "10:00 AM",
  "10:30 AM",
  "11:00 AM",
  "11:30 AM",
  "12:00 PM",
  "12:30 PM",
  "1:00 PM",
  "1:30 PM",
  "2:00 PM",
  "2:30 PM",
  "3:00 PM",
  "3:30 PM",
  "4:00 PM",
  "4:30 PM",
  "5:00 PM",
  "5:30 PM",
  "6:00 PM",
  "6:30 PM",
  "7:00 PM",
  "7:30 PM",
  "8:00 PM",
  "8:30 PM",
];

export function convertAmPmToIsoTime(time) {
  switch (time) {
    case "5:00 AM":
    case "5:30 AM":
      return "T05:00:00+00:00";
    case "6:00 AM":
    case "6:30 AM":
      return "T06:00:00+00:00";
    case "7:00 AM":
    case "7:30 AM":
      return "T07:00:00+00:00";
    case "8:00 AM":
    case "8:30 AM":
      return "T08:00:00+00:00";
    case "9:00 AM":
    case "9:30 AM":
      return "T09:00:00+00:00";
    case "10:00 AM":
    case "10:30 AM":
      return "T10:00:00+00:00";
    case "11:00 AM":
    case "11:30 AM":
      return "T11:00:00+00:00";
    case "12:00 PM":
    case "12:30 PM":
      return "T12:00:00+00:00";
    case "1:00 PM":
    case "1:30 PM":
      return "T13:00:00+00:00";
    case "2:00 PM":
    case "2:30 PM":
      return "T14:00:00+00:00";
    case "3:00 PM":
    case "3:30 PM":
      return "T15:00:00+00:00";
    case "4:00 PM":
    case "4:30 PM":
      return "T16:00:00+00:00";
    case "5:00 PM":
    case "5:30 PM":
      return "T17:00:00+00:00";
    case "6:00 PM":
    case "6:30 PM":
      return "T18:00:00+00:00";
    case "7:00 PM":
    case "7:30 PM":
      return "T19:00:00+00:00";
    case "8:00 PM":
    case "8:30 PM":
      return "T20:00:00+00:00";
    default:
      return "Error";
  }
}

export function convert24HrToAmPm(str) {
  let convertedStr;
  if (+str.substring(0, 2) < 12) {
    convertedStr = str + " AM";
    if (str[0] == "0") {
      convertedStr = convertedStr.slice(1);
    }
    if (convertedStr.charAt(0) === "0") {
      convertedStr = "12" + convertedStr.substring(1);
    }
  } else if (+str.substring(0, 2) === 12) {
    convertedStr = +str.substring(0, 2) + str.slice(2) + " PM";
  } else {
    convertedStr = +str.substring(0, 2) - 12 + str.slice(2) + " PM";
  }

  return convertedStr;
}
