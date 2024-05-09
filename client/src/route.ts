import { writable } from "svelte/store"

export let route = writable("/")

export function goto(path: string) {
    route.set(path)
    location.hash = `#${path}`
}