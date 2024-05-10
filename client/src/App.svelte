<script lang="ts">
    import Login from "./Login.svelte";
    import Dashboard from "./Dashboard.svelte";
    import { goto, route } from "./route";
    import Reg from "./Reg.svelte";
    import ChangePassword from "./ChangePassword.svelte";

    let routePath = $route;
    $: {
        console.info("App.routePath", routePath);
    }
    route.subscribe((v) => {
        routePath = v;
    });

    window.addEventListener("hashchange", (e) => {
        console.info("window.hashchange", e);
        let index = e.newURL.indexOf("#");
        if (index > -1) {
            let hash = e.newURL.substring(index + 1);
            goto(hash);
        }
    });
</script>

<main>
    {#if routePath == "/login"}
        <Login />
    {:else if routePath == "/reg"}
        <Reg />
    {:else if routePath == "/dash"}
        <Dashboard />
    {:else if routePath == "/changePassword"}
        <ChangePassword />
    {/if}
</main>
