<script lang="ts">
    import { Modal } from "bootstrap";
    import { init, type EChartsType } from "echarts";
    import { SunburstChart } from "echarts/charts";
    import {
        TitleComponent,
        ToolboxComponent,
        TooltipComponent,
    } from "echarts/components";
    import * as echarts from "echarts/core";
    import { CanvasRenderer } from "echarts/renderers";
    import type { datastore } from "./datastore";

    const ICON_SHARE =
        "M11 2.5a2.5 2.5 0 1 1 .603 1.628l-6.718 3.12a2.499 2.499 0 0 1 0 1.504l6.718 3.12a2.5 2.5 0 1 1-.488.876l-6.718-3.12a2.5 2.5 0 1 1 0-3.256l6.718-3.12A2.5 2.5 0 0 1 11 2.5z";

    export let data: datastore.Item[];
    export let showOther: boolean;
    export let percent: boolean;

    let chartDom: HTMLElement;
    let chart: EChartsType;
    let shareDom: HTMLElement;
    let shareImgContent: string;
    let shareDialog: Modal;

    const formatNumber = (v: number) => {
        if (v >= 1000000) {
            return (v / 1000000).toFixed(2) + " 百万";
        } else if (v >= 10000) {
            return (v / 10000).toFixed(1) + " 万";
        } else if (v >= 1000) {
            return (v / 1000).toFixed(1) + " 千";
        } else {
            return v.toString();
        }
    };

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

    $: if (shareDom) {
        shareDialog = new Modal(shareDom);
        shareDialog.hide();
    }

    $: if (chartDom && data && showOther != null) {
        if (!chart) {
            chart = init(chartDom);
        }
        renderChart();
    }
</script>

<div bind:this={chartDom} class="chart_dom mt-3"></div>
<div class="modal" tabindex="-1" bind:this={shareDom}>
    <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">分享</h5>
                <button
                    type="button"
                    class="btn-close"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                ></button>
            </div>
            <div class="modal-body">
                <p>长按图片保存或者分享</p>
                <div
                    class="d-flex flex-column justify-content-center align-items-center"
                >
                    <img src={shareImgContent} class="w-100" alt="chart" />
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .chart_dom {
        width: 100%;
        aspect-ratio: 1;
        border-width: 1px;
        border-color: black;
        border-style: solid;
        min-width: 300px;
    }
</style>
