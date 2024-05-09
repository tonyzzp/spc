import "bootstrap"
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap-icons/font/bootstrap-icons.css"
import App from "./App.svelte"
import { api } from "./api"
import { goto } from "./route"

console.info("location.hash", location.hash)
let target = ""
if (location.hash != "#/login" && location.hash != "#/reg") {
  let token = localStorage.getItem("spc-token")
  console.info("token", token)
  if (token) {
    api.setToken(token)
    target = "/dash"
  }
} else if (location.hash == "#/login") {
  target = "/login"
} else if (location.hash == "#/reg") {
  target = "/reg"
}
if (!target) {
  target = "/login"
}
goto(target)

const app = new App({
  target: document.getElementById('app')!,
})

export default app
