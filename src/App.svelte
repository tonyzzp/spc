<script lang="ts">
    import { BarChart, LineChart, SunburstChart } from "echarts/charts";
    import {
        DatasetComponent,
        GridComponent,
        TitleComponent,
        ToolboxComponent,
        TooltipComponent,
        TransformComponent,
    } from "echarts/components";
    import * as echarts from "echarts/core";
    import { init, type EChartsType } from "echarts/core";
    import { LabelLayout, UniversalTransition } from "echarts/features";
    import { CanvasRenderer } from "echarts/renderers";
    import { datastore } from "./datastore";
    import type { etypes } from "./types";

    let chartDom1: HTMLElement;
    let chartDom2: HTMLElement;
    let chart1: EChartsType;
    let chart2: EChartsType;

    let data: datastore.Item[] = [];
    let iptName = "";
    let iptType = "";
    let iptValue = "";
    let showAddArea = false;

    const onAddClick = () => {
        iptName = "";
        iptType = "";
        iptValue = "";
        showAddArea = !showAddArea;
    };

    const onAddConfirmClick = () => {
        console.info(iptName, iptType, iptValue);
        if (iptName == "" || iptType == "" || !parseInt(iptValue)) {
            alert("请输入数据");
            return;
        }
        datastore.put({
            name: iptName,
            type: iptType,
            value: parseInt(iptValue),
        });
        reloadData();
        iptName = "";
        iptType = "";
        iptValue = "";
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
            datastore.remove(item.name);
            reloadData();
        }
    };

    const initView = () => {
        datastore.init();
        reloadData();
    };

    const reloadData = () => {
        data = datastore.all();
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

    datastore.init();
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

<main>
    <button class="btn btn-info" on:click={onAddClick}
        >{showAddArea ? "取消" : "添加"}</button
    >

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
            />
        </div>
        <button
            class="btn btn-primary w-100 mb-1 mt-1"
            on:click={onAddConfirmClick}>添加</button
        >
    </div>

    <table class="table table-striped">
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
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .chart_dom {
        width: 80%;
        max-width: 600px;
        aspect-ratio: 1;
        border-width: 1px;
        border-color: black;
        border-style: solid;
    }
</style>
