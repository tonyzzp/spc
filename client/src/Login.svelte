<script lang="ts">
    import { api } from "./api";
    import { goto } from "./route";

    let iptUser = "";
    let iptPassword = "";
    let err = "";

    const onLoginClick = async () => {
        if (!iptUser || !iptPassword) {
            err = "请输入用户名密码";
            return;
        }
        let result = await api.login(iptUser, iptPassword);
        if (result == null || result.code) {
            err = result?.msg || "登录失败";
            return;
        }
        console.info("登录成功", result.data);
        api.setToken(result.data.token);
        localStorage.setItem("spc-token", result.data.token);
        goto("/dash");
    };

    const onRegClick = async () => {
        goto("/reg");
    };
</script>

<main>
    <h1 class="text-center mt-3 mb-3">登录</h1>
    <form class="border border-2 rounded-2 p-2 m-2">
        <div class="mb-2">
            <label for="ipt_user" class="form-label">用户名</label>
            <input
                type="text"
                class="form-control"
                maxlength="20"
                id="ipt_user"
                bind:value={iptUser}
            />
        </div>
        <div class="mb-2">
            <label for="ipt_password" class="form-label">密码</label>
            <input
                type="password"
                class="form-control"
                maxlength="20"
                id="ipt_password"
                bind:value={iptPassword}
                on:keypress={(data) => {
                    console.info(data);
                    if (data.key == "Enter") {
                        onLoginClick();
                    }
                }}
            />
        </div>
        <p class="text-danger">{err}&nbsp;</p>
        <button
            class="btn btn-primary mt-2"
            type="button"
            on:click={onLoginClick}>登录</button
        >
        <div class="d-flex justify-content-end">
            <button type="button" class="btn btn-link" on:click={onRegClick}
                >注册</button
            >
        </div>
    </form>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    form {
        min-width: 500px;
        max-width: 800px;
        display: flex;
        flex-direction: column;
    }
</style>
