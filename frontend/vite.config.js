import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vitest/config";

import { loadEnv } from "vite";

export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd());
    const apiBaseUrl = env.VITE_API_BASE_URL;
    const frontBaseUrl = env.VITE_FRONT_BASE_URL;
    console.log({
        apiBaseUrl,
        frontBaseUrl,
    });

    return {
        plugins: [sveltekit()],
        test: {
            include: ["src/**/*.{test,spec}.{js,ts}"],
        },
        server: {
            proxy: {
                "/api": {
                    target: apiBaseUrl,
                    changeOrigin: true,
                    rewrite: (path) => path.replace(/^\/api/, ""),
                },
            },
        },
    };
});
