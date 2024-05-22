<script lang="ts">
    import Login from "./pages/Login.svelte";
    import Dashboard from "./pages/Dashboard.svelte";
    import { goto, route } from "./common/route";
    import Reg from "./pages/Reg.svelte";
    import ChangePassword from "./pages/ChangePassword.svelte";

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
    {#if routePath == "/reg"}
        <Reg />
    {:else if routePath == "/dash"}
        <Dashboard />
    {:else if routePath == "/changePassword"}
        <ChangePassword />
    {:else}
        <Login />
    {/if}
</main>
