import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [svelte()],
	build: {
		rollupOptions: {
			output: {
				manualChunks: function (id, meta) {
					console.info(id)
					let index = id.indexOf(".pnpm/")
					if (index > -1) {
						let index2 = id.indexOf("/", index + 6)
						if (index2 > -1) {
							let rtn = id.substring(index + 6, index2)
							if (rtn) {
								if (rtn.indexOf("@") == 0) {
									rtn = rtn.substring(1)
								}
								return rtn
							}
						}
					}
					index = id.indexOf("node_modules/")
					if (index > -1) {
						let index2 = id.indexOf("/", index + 13)
						if (index2 > -1) {
							let rtn = id.substring(index + 13, index2)
							if (rtn) {
								if (rtn.indexOf("@") == 0) {
									rtn = rtn.substring(1)
								}
								return rtn
							}
						}
					}
				}
			}
		}
	},
	server: {
		proxy: {
			"^/api/.*": "http://localhost:80"
		},
		host: "0.0.0.0",
		port: 10000,
	}
})
