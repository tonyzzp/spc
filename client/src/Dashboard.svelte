<script lang="ts">
    import { SunburstChart } from "echarts/charts";
    import {
        TitleComponent,
        ToolboxComponent,
        TooltipComponent,
    } from "echarts/components";
    import * as echarts from "echarts/core";
    import { init, type EChartsType } from "echarts/core";
    import { CanvasRenderer } from "echarts/renderers";
    import { api } from "./api";
    import { datastore } from "./datastore";
    import { goto } from "./route";
    import type { etypes } from "./types";
    import { Toast } from "bootstrap";
    import Navi from "./Navi.svelte";

    let chartDom1: HTMLElement;
    let chartDom2: HTMLElement;
    let chart1: EChartsType;
    let chart2: EChartsType;

    let data: datastore.Item[] = [];
    let iptName = "";
    let iptType = "";
    let iptValue = "";
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
            data = data.splice(index, 1);
        }
        data = [
            ...data,
            {
                name: iptName,
                type: iptType,
                value: parseInt(iptValue),
            },
        ];
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
        } catch (e) {
            err = `${e}`;
            toast.show();
        }
    };

    const formatNumber = (v: number) => {
        if (v >= 1000000) {
            return (v / 1000000).toFixed(2) + " 百万";
        } else if (v >= 10000) {
            return (v / 10000).toFixed(1) + " 万";
        } else if (v >= 1000) {
            return (v / 1000).toFixed(1) + " 千";
        }
    };

    const renderChart = (chart: EChartsType, percent: boolean) => {
        let formatter = function (v: any) {
            if (percent) {
                return v.name + "," + (v.data.value * 100).toFixed(1) + "%";
            } else {
                return `${v.name},${formatNumber(v.value)}`;
            }
        };
        let option: etypes.Option = {
            title: {
                show: true,
                text: percent ? "资产占比" : "资产详情",
                left: "center",
            },
            tooltip: {
                show: false,
            },
            toolbox: {
                show: true,
                feature: {
                    saveAsImage: {
                        show: true,
                    },
                    dataView: {
                        show: true,
                        readOnly: true,
                    },
                },
            },
            series: {
                type: "sunburst",
                emphasis: {
                    focus: "none",
                },
                label: {
                    show: true,
                    formatter: formatter,
                },
                nodeClick: false,
                data: [],
            },
        };
        let total = 0;
        let group: { [type: string]: datastore.Item[] } = {};
        data.forEach((stock) => {
            let list = group[stock.type];
            if (!list) {
                list = [];
                group[stock.type] = list;
            }
            list.push(stock);
            total += stock.value;
        });

        for (let type in group) {
            let list = group[type];
            let item: etypes.Item = {
                name: type,
                children: list.map((v) => {
                    return {
                        name: v.name,
                        value: percent ? v.value / total : v.value,
                    };
                }),
            };
            option.series.data.push(item);
        }
        chart.setOption(option as any);
    };

    echarts.use([
        SunburstChart,
        TitleComponent,
        TooltipComponent,
        ToolboxComponent,
        CanvasRenderer,
    ]);
    window.addEventListener("resize", () => {
        chart1?.resize();
        chart2?.resize();
    });

    initView();

    $: if (chartDom1 && data) {
        if (!chart1) {
            chart1 = init(chartDom1);
        }
        renderChart(chart1, false);
        if (!chart2) {
            chart2 = init(chartDom2);
        }
        renderChart(chart2, true);
    }
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
</div>

<div class="border m-2 p-2" class:d-none={!showAddArea}>
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

<table class="table table-striped table-hover align-middle">
    <thead>
        <tr>
            <th>名称</th>
            <th>金额</th>
            <th>分类</th>
            <th>修改</th>
            <th>删除</th>
        </tr>
    </thead>
    <tbody>
        {#each data as item}
            <tr>
                <td>{item.name}</td>
                <td>{item.value.toLocaleString("en-us")}</td>
                <td>{item.type}</td>
                <td>
                    <button
                        class="btn btn-light"
                        on:click={onEditClick.bind(null, item)}
                        ><i class="bi bi-pencil-square"></i></button
                    >
                </td>
                <td>
                    <button
                        class="btn btn-light"
                        on:click={onDelClick.bind(null, item)}
                    >
                        <i class="bi bi-trash"></i>
                    </button>
                </td>
            </tr>
        {/each}
    </tbody>
</table>

<div bind:this={chartDom1} class="chart_dom mt-3"></div>
<div bind:this={chartDom2} class="chart_dom mt-3"></div>

<style>
    .chart_dom {
        width: 100%;
        aspect-ratio: 1;
        border-width: 1px;
        border-color: black;
        border-style: solid;
        min-width: 300px;
    }

    .btn-bar {
        display: flex;
        flex-direction: row;
        gap: 10px;
    }

    table {
        min-width: 300px;
    }
</style>
