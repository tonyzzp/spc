<script lang="ts">
    import { Toast } from "bootstrap";
    import ChartBar from "../components/ChartBar.svelte";
    import ChartSunburst from "../components/ChartSunburst.svelte";
    import Navi from "../components/Navi.svelte";
    import { api } from "../common/api";
    import { datastore } from "../common/datastore";
    import { goto } from "../common/route";
    import { formatNumber, sortData } from "../common/utils";

    let data: datastore.Item[] = [];
    let totalValue = 0;
    let iptName = "";
    let iptType = "";
    let iptValue = "";
    let showOther = localStorage.getItem("spc-show-other") == "true";
    let showAddArea = false;
    let err = "";
    let toastDom: HTMLElement;
    let toast: Toast;

    const save = async () => {
        let ret = await api.save(JSON.stringify(data));
        if (ret == null || ret.code) {
            err = ret.msg || "保存失败";
            toast.show();
        }
    };

    $: if (toastDom) {
        toast = new Toast(toastDom);
        toast.hide();
    }

    $: {
        localStorage.setItem("spc-show-other", showOther ? "true" : "false");
    }

    $: if (data != null && data.length > 0) {
        totalValue = data.reduce((rtn, current) => rtn + current.value, 0);
    }

    const onAddClick = () => {
        iptName = "";
        iptType = "";
        iptValue = "";
        showAddArea = !showAddArea;
    };

    const onAddConfirmClick = () => {
        console.info(iptName, iptType, iptValue);
        if (iptName == "" || iptType == "" || !parseInt(iptValue)) {
            err = "请输入数据";
            toast.show();
            return;
        }
        let index = data.findIndex((v) => v.name == iptName);
        if (index > -1) {
            data.splice(index, 1);
        }
        data = [
            ...data,
            {
                name: iptName,
                type: iptType,
                value: parseInt(iptValue),
            },
        ];
        sortData(data);
        iptName = "";
        iptType = "";
        iptValue = "";
        showAddArea = false;
        save();
    };

    const onCancelClick = () => {
        showAddArea = false;
    };

    const onEditClick = (item: datastore.Item) => {
        console.info("edit", item);
        iptName = item.name;
        iptType = item.type;
        iptValue = item.value.toString();
        showAddArea = true;
    };

    const onDelClick = (item: datastore.Item) => {
        console.info("del", item);
        let ok = confirm(`删除 ${item.name}, ${item.value} ?`);
        if (ok) {
            let index = data.findIndex((v) => v.name == item.name);
            if (index != -1) {
                data.splice(index, 1);
                data = [...data];
            }
        }
        save();
    };

    const onKeyPress = (e: KeyboardEvent) => {
        console.info(e);
        if (e.key == "Enter") {
            onAddConfirmClick();
        }
    };

    const onLogoutClick = () => {
        api.setToken("");
        localStorage.removeItem("spc-token");
        goto("/login");
    };

    const initView = async () => {
        try {
            data = await datastore.load();
            sortData(data);
        } catch (e) {
            err = `${e}`;
            toast.show();
        }
    };

    initView();
</script>

<Navi />
<h1>我的资产</h1>
<div class="toast-container">
    <div class="toast show text-bg-danger" bind:this={toastDom}>
        <div class="d-flex">
            <div class="toast-body">{err}</div>
            <button
                type="button"
                class="btn-close btn-close-white me-2 m-auto"
                data-bs-dismiss="toast"
            ></button>
        </div>
    </div>
</div>

<div class="btn-bar mb-3 align-items-center">
    <span class="text-secondary fs-5">user: {api.getUserName()}</span>
    <button class="btn btn-info" on:click={onAddClick}>添加资产</button>
    <button class="btn btn-secondary" on:click={onLogoutClick}>退出登录</button>
    <div class="form-check">
        <input
            class="form-check-input"
            type="checkbox"
            id="flexCheckChecked"
            bind:checked={showOther}
        />
        <label class="form-check-label" for="flexCheckChecked"
            >低于5%归为其他</label
        >
    </div>
</div>

<div class="border m-2 p-2 w-100" class:d-none={!showAddArea}>
    <div class="mb-1">
        <label for="ipt_name" class="form-label">名称</label>
        <input
            type="text"
            class="form-control"
            id="ipt_name"
            bind:value={iptName}
        />
    </div>
    <div class="mb-1">
        <label for="ipt_type" class="form-label">类别</label>
        <input
            type="text"
            class="form-control"
            id="ipt_type"
            bind:value={iptType}
        />
    </div>
    <div class="mb-1">
        <label for="ipt_value" class="form-label">总金额</label>
        <input
            type="number"
            class="form-control"
            id="ipt_value"
            bind:value={iptValue}
            on:keypress={onKeyPress}
        />
    </div>
    <div class="d-flex flex-row">
        <button class="btn btn-primary w-75" on:click={onAddConfirmClick}
            >确定</button
        >
        <button class="btn btn-secondary w-25 ms-3" on:click={onCancelClick}
            >取消</button
        >
    </div>
</div>

<table class="table table-striped align-middle">
    <thead>
        <tr>
            <th>名称</th>
            <th>金额</th>
            <th>分类</th>
            <!-- <th>修改</th>
            <th>删除</th> -->
            <th>操作</th>
        </tr>
    </thead>
    <tbody>
        {#each data as item}
            <tr>
                <td>{item.name}</td>
                <td>{item.value.toLocaleString("en-us")}</td>
                <td>{item.type}</td>
                <td>
                    <div class="dropdown">
                        <button
                            class="btn btn-secondary dropdown-toggle"
                            type="button"
                            data-bs-toggle="dropdown">操作</button
                        >
                        <ul class="dropdown-menu">
                            <li>
                                <button
                                    class="dropdown-item"
                                    type="button"
                                    on:click={onEditClick.bind(null, item)}
                                    ><i class="bi bi-pencil-square"></i> 修改</button
                                >
                            </li>
                            <li>
                                <button
                                    class="dropdown-item"
                                    type="button"
                                    on:click={onDelClick.bind(null, item)}
                                    ><i class="bi bi-trash"></i> 删除</button
                                >
                            </li>
                        </ul>
                    </div>
                </td>
            </tr>
        {/each}
    </tbody>
</table>
<b class="d-block text-end w-100">总额: {formatNumber(totalValue)}</b>
<br />

<ChartSunburst {data} percent={false} {showOther} />
<ChartSunburst {data} percent={true} {showOther} />
<ChartBar {data} percent={false} {showOther} />
<ChartBar {data} percent={true} {showOther} />

<style>
    .btn-bar {
        display: flex;
        flex-direction: row;
        gap: 10px;
    }

    table {
        min-width: 300px;
    }
</style>
