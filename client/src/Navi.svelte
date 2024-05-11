<script lang="ts">
    import { route } from "./route";

    let routePath = $route;
</script>

<ul class="nav nav-tabs nav-fill mb-3">
    <li class="nav-item">
        <a class="nav-link" href="/#/dash" class:active={routePath == "/dash"}
            ><i class="bi bi-house"></i></a
        >
    </li>
    <li class="nav-item">
        <a
            class="nav-link"
            href="/#/changePassword"
            class:active={routePath == "/changePassword"}
            ><i class="bi bi-key"></i></a
        >
    </li>
</ul>
