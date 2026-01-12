import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import devtoolsJson from 'vite-plugin-devtools-json';

export default defineConfig({
	plugins: [sveltekit(), tailwindcss(), devtoolsJson()],

	server: {
		proxy: {
			'^/(api|_)/': { target: 'http://localhost:8090', changeOrigin: true },
		},
	},
});
