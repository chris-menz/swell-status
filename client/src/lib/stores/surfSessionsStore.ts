import { Writable, writable } from "svelte/store";
import { SurfSession } from "../utils/models";
import axios from "axios";
import { axiosConfig } from "../utils/axiosConfig";
import { endpoint } from "../utils/endpoint";

export const surfSessions: Writable<SurfSession[]> = writable([]);

export async function getSurfSessions() {
  try {
    const response = await axios.get(endpoint + "/surfsession", axiosConfig);
    if (response.data.message == "success") {
      surfSessions.update(
        (surfSessions) => (surfSessions = response.data.surf_sessions)
      );
    }
  } catch (error) {
    console.log(error);
  }
  return "success";
}

export async function deleteSurfSession(id) {
  try {
    const response = await axios.delete(
      endpoint + "/auth/surfsession/" + id,
      axiosConfig
    );
    const response2 = await axios.get(endpoint + "/surfsession", axiosConfig);
    if (response2.data.message == "success") {
      surfSessions.update(
        (surfSessions) => (surfSessions = response2.data.surf_sessions)
      );
    }
    if (response2.data.message == "no sessions found") {
      surfSessions.update((surfSessions) => (surfSessions = []));
    }
  } catch (error) {
    console.log(error);
  }
}
