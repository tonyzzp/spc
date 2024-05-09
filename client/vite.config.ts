import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [svelte()],
	build: {
		rollupOptions: {
			output: {
				manualChunks: function (id, meta) {
					console.info(id, meta)
					{
						let index = id.indexOf(".pnpm/")
						if (index > -1) {
							let index2 = id.indexOf("/", index + 6)
							if (index2 > -1) {
								let rtn = id.substring(index + 6, index2)
								if (rtn) {
									return rtn
								}
							}
						}
					}
				}
			}
		}
	},
	server: {
		proxy: {
			"^/api/.*": "http://localhost:8080"
		}
	}
})
