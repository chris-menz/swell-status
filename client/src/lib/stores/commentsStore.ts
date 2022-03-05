import { Writable, writable } from "svelte/store";
import { Comment } from "../utils/models";
import axios from "axios";
import { axiosConfig } from "../utils/axiosConfig";
import { endpoint } from "../utils/endpoint";

export const comments: Writable<Comment[]> = writable([]);

export async function getComments() {
  try {
    const response = await axios.get(endpoint + "/comment", axiosConfig);
    comments.update((comments) => (comments = response.data));
  } catch (error) {
    console.log(error);
  }
  return "success";
}

export async function postComment(comment) {
  try {
    const response = await axios.post(
      endpoint + "/auth/comment",
      comment,
      axiosConfig
    );
    const response2 = await axios.get(endpoint + "/comment", axiosConfig);
    comments.update((comments) => (comments = response2.data));
  } catch (error) {
    console.log(error);
  }
}
