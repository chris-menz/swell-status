import { Writable, writable } from "svelte/store";
import { User } from "../utils/models";
import axios from "axios";
import { axiosConfig } from "../utils/axiosConfig";
import { endpoint } from "../utils/endpoint";
import { getSavedSpots, savedSpots } from "./savedSpotStore";

// user represents logged in user
export const user: Writable<User> = writable(null);

export const users: Writable<User[]> = writable([]);

export async function isAuthenticated() {
  const response = await axios.get(
    endpoint + "/auth/getcurruserid",
    axiosConfig
  );
  const data = response.data;

  if (data.message == "user is logged in") {
    const userres = await axios.get(`${endpoint}/user/${data.id}`);
    user.update((user) => (user = userres.data));
    return true;
  } else {
    user.update((user) => (user = null));
    return false;
  }
}

export async function getUsers() {
  try {
    const response = await axios.get(endpoint + "/user", axiosConfig);
    users.update((users) => (users = response.data));
  } catch (error) {
    console.log(error);
  }
  return "success";
}

export async function attemptLogin(credentials) {
  let message: string;
  try {
    const response = await axios.post(
      endpoint + "/login",
      credentials,
      axiosConfig
    );
    if (response.data.message == "login successful") {
      user.update((user) => (user = response.data.user));
    }
    message = response.data.message;
  } catch (error) {
    console.log(error);
  }
  return message;
}

export async function logout() {
  const response = await axios.post(
    endpoint + "/auth/logout",
    null,
    axiosConfig
  );
  savedSpots.update((savedSpots) => (savedSpots = []));
  user.update((user) => (user = null));
}
