<script lang="ts">
    import Navi from "./Navi.svelte";
    import { api } from "./api";
    import { goto } from "./route";

    let iptOldPassword = "";
    let iptPassword1 = "";
    let iptPassword2 = "";
    let err = "";

    const onConfirmClick = async () => {
        iptOldPassword = iptOldPassword.trim();
        iptPassword1 = iptPassword1.trim();
        iptPassword2 = iptPassword2.trim();
        if (!iptOldPassword || !iptPassword1 || !iptPassword2) {
            err = "输入不完整";
            return;
        }
        if (iptPassword1 != iptPassword2) {
            err = "新密码不一致";
            return;
        }
        let result = await api.changePassword(iptOldPassword, iptPassword1);
        if (result == null || result.code) {
            err = result?.msg || "登录失败";
            return;
        }
        console.info("修改成功", result);
        alert("修改成功");
        goto("/dash");
    };
</script>

<Navi />
<h1>修改密码</h1>
<div class="border border-2 rounded-2 p-2 m-2 w-100">
    <div class="mb-2">
        <label for="ipt_oldpassword" class="form-label">旧密码</label>
        <input
            type="password"
            class="form-control"
            maxlength="20"
            id="ipt_oldpassword"
            bind:value={iptOldPassword}
        />
    </div>
    <div class="mb-2">
        <label for="ipt_password1" class="form-label">新密码</label>
        <input
            type="password"
            class="form-control"
            maxlength="20"
            id="ipt_password1"
            bind:value={iptPassword1}
        />
    </div>
    <div class="mb-2">
        <label for="ipt_password2" class="form-label">重复新密码</label>
        <input
            type="password"
            class="form-control"
            maxlength="20"
            id="ipt_password2"
            bind:value={iptPassword2}
            on:keypress={(data) => {
                console.info(data);
                if (data.key == "Enter") {
                    onConfirmClick();
                }
            }}
        />
    </div>
    <p class="text-danger">{err}&nbsp;</p>
    <button class="btn btn-primary mt-2" type="button" on:click={onConfirmClick}
        >确定</button
    >
</div>
