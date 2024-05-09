<script lang="ts">
    import { api } from "./api";
    import { goto } from "./route";

    let iptUser = "";
    let iptPassword = "";
    let iptPassword2 = "";
    let err = "";

    const onConfrimClick = async () => {
        if (!iptUser || !iptPassword) {
            err = "请输入用户名密码";
            return;
        }
        if (iptPassword != iptPassword2) {
            err = "密码不一致";
            return;
        }
        let result = await api.reg(iptUser, iptPassword);
        if (result == null || result.code) {
            err = result?.msg || "登录失败";
            return;
        }
        console.info("注册成功", result);
        alert("注册成功");
        goto("/login");
    };

    const onPasswordChange = () => {
        setTimeout(() => {
            iptPassword = iptPassword.trim();
            iptPassword2 = iptPassword2.trim();
            console.info(iptPassword, iptPassword2);
            if (iptPassword != iptPassword2 && iptPassword != "") {
                err = "密码不一致";
            } else {
                err = "";
            }
        }, 1);
    };

    const onLoginClick = async () => {
        goto("/login");
    };
</script>

<h1 class="text-center mt-3 mb-3">注册</h1>

<main>
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
                on:input={onPasswordChange}
                type="password"
                class="form-control"
                maxlength="20"
                id="ipt_password"
                bind:value={iptPassword}
            />
        </div>
        <div class="mb-2">
            <label for="ipt_password2" class="form-label">重复密码</label>
            <input
                on:input={onPasswordChange}
                type="password"
                class="form-control"
                maxlength="20"
                id="ipt_password2"
                bind:value={iptPassword2}
                on:keypress={(data) => {
                    console.info(data);
                    if (data.key == "Enter") {
                        onConfrimClick();
                    }
                }}
            />
        </div>
        <p class="text-danger">{err}&nbsp;</p>
        <button
            class="btn btn-primary mt-2"
            type="button"
            on:click={onConfrimClick}>注册</button
        >
        <div class="d-flex justify-content-end">
            <button type="button" class="btn btn-link" on:click={onLoginClick}
                >登录</button
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
