import { Writable, writable } from "svelte/store";
import { Favorite } from "../utils/models";
import axios from "axios";
import { axiosConfig } from "../utils/axiosConfig";
import { endpoint } from "../utils/endpoint";

export const favorites: Writable<Favorite[]> = writable([]);

export async function getFavorites() {
  try {
    const response = await axios.get(endpoint + "/favorite", axiosConfig);
    favorites.update((favorites) => (favorites = response.data));
  } catch (error) {
    console.log(error);
  }
  return "success";
}

export async function addFavorite(sessionId: number, userId: number) {
  try {
    const newFav = {
      surf_session_id: sessionId,
      user_id: userId,
    };
    const response = await axios.post(
      endpoint + "/auth/favorite",
      newFav,
      axiosConfig
    );
    getFavorites();
  } catch (error) {
    console.log(error);
  }
}

export async function removeFavorite(id: number) {
  try {
    const response = await axios.delete(
      endpoint + "/auth/favorite/" + id,
      axiosConfig
    );
    getFavorites();
  } catch (error) {
    console.log(error);
  }
}
