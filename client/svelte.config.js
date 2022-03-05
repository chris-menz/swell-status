import preprocess from "svelte-preprocess";
import adapter from "@sveltejs/adapter-auto";
import Icons from "unplugin-icons/vite";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: preprocess(),

  kit: {
    adapter: adapter(),
    vite: {
      plugins: [
        Icons({
          compiler: "svelte",
        }),
      ],
    },
  },
};

export default config;
