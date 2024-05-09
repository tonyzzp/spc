<script lang="ts">
    import Login from "./Login.svelte";
    import Dashboard from "./Dashboard.svelte";
    import { route } from "./route";
    import Reg from "./Reg.svelte";

    let routePath = $route;
    $: {
        console.info("App.routePath", routePath);
    }
    route.subscribe((v) => {
        routePath = v;
    });
</script>

{#if routePath == "/login"}
    <Login />
{:else if routePath == "/reg"}
    <Reg />
{:else if routePath == "/dash"}
    <Dashboard />
{/if}
