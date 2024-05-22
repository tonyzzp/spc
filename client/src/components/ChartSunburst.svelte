<script lang="ts">
    import { init, type EChartsType } from "echarts";
    import { SunburstChart } from "echarts/charts";
    import {
        TitleComponent,
        ToolboxComponent,
        TooltipComponent,
    } from "echarts/components";
    import * as echarts from "echarts/core";
    import { CanvasRenderer } from "echarts/renderers";
    import ShareDailog from "./ShareDailog.svelte";
    import type { datastore } from "../common/datastore";
    import { ICON_SHARE, formatNumber } from "../common/utils";

    export let data: datastore.Item[];
    export let showOther: boolean;
    export let percent: boolean;

    let chartDom: HTMLElement;
    let chart: EChartsType;
    let shareImgContent: string;
    let shareDialog: ShareDailog;

    const renderChart = () => {
        let formatter = function (v: any) {
            if (percent) {
                return v.name + "," + (v.data.value * 100).toFixed(1) + "%";
            } else {
                return `${v.name},${formatNumber(v.value)}`;
            }
        };
        let option: any = {
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
                    myShareImage: {
                        show: true,
                        title: "分享",
                        icon: `path://${ICON_SHARE}`,
                        onclick: (_: any, extension: any) => {
                            shareImgContent = extension.getDataURL();
                            shareDialog.show();
                        },
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
            list.push({ ...stock });
            total += stock.value;
        });
        if (showOther) {
            let otherValue = 0;
            let others: datastore.Item[] = [];
            for (let type in group) {
                let list = group[type];
                for (let i = list.length - 1; i >= 0; i--) {
                    let item = list[i];
                    if (item.value / total < 0.05) {
                        list.splice(i, 1);
                        others.push(item);
                        otherValue += item.value;
                    }
                }
            }
            if (otherValue) {
                group.others = [
                    {
                        name: "others",
                        type: "others",
                        value: otherValue,
                    },
                ];
            }
        }

        for (let type in group) {
            let list = group[type];
            let item: any = {
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
        chart.setOption(option);
    };

    echarts.use([
        SunburstChart,
        TitleComponent,
        TooltipComponent,
        ToolboxComponent,
        CanvasRenderer,
    ]);
    window.addEventListener("resize", () => {
        chart?.resize();
    });

    $: if (chartDom && data && showOther != null) {
        if (!chart) {
            chart = init(chartDom);
        }
        renderChart();
    }
</script>

<div bind:this={chartDom} class="chart_dom mt-3"></div>
<ShareDailog imgContent={shareImgContent} bind:this={shareDialog} />
