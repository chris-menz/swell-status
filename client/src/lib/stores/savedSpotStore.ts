import { Writable, writable } from "svelte/store";
import { SavedSpot, User } from "../utils/models";
import axios from "axios";
import { axiosConfig } from "../utils/axiosConfig";
import { endpoint } from "../utils/endpoint";

export const savedSpots: Writable<SavedSpot[]> = writable([]);

export async function getSavedSpots(user: User) {
  if (user) {
    const response = await axios.get(
      endpoint + "/spot/user/" + user.id,
      axiosConfig
    );
    if (response.data.message == "Success" && response.data.spots) {
      savedSpots.update((savedSpots) => (savedSpots = response.data.spots));
      return;
    }
  }
  savedSpots.update((savedSpots) => (savedSpots = []));
}

export async function saveNewSpot(
  spot_name: string,
  spot_region: string,
  user: User
) {
  const newSavedSpot = {
    spot_name,
    spot_region,
    user_id: user.id,
  };
  const response = await axios.post(
    endpoint + "/auth/spot",
    newSavedSpot,
    axiosConfig
  );
  getSavedSpots(user);
}

export async function deleteSpot(spotId: number, user: User) {
  const response = await axios.delete(
    endpoint + "/auth/spot/" + spotId,
    axiosConfig
  );
  getSavedSpots(user);
}
