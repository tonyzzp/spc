import "bootstrap"
import "bootstrap-icons/font/bootstrap-icons.css"
import "bootstrap/dist/css/bootstrap.css"
import { use } from "echarts"
import { BarChart, SunburstChart } from "echarts/charts"
import { GridComponent } from "echarts/components"
import "./app.css"
import App from "./App.svelte"
import { api } from "./common/api"
import { goto } from "./common/route"


use([
  GridComponent,
  BarChart,
  SunburstChart,
])


console.info("location.hash", location.hash)

let token = localStorage.getItem("spc-token") || ""
api.setToken(token)
let hash = location.hash
if (hash) {
  if (hash == "#/dash" && !token) {
    goto("/login")
  } else {
    goto(hash.substring(1))
  }
} else {
  goto("/login")
}

const app = new App({
  target: document.getElementById('app')!,
})

export default app
