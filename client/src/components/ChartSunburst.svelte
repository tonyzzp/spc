<script lang="ts" context="module">
    function createChart(
        chart: EChartsType,
        dom: HTMLElement,
        data: datastore.Item[],
        showOther: boolean,
        percent: boolean,
        onShareClick: (extension: any) => void,
    ) {
        let formatter = function (v: any) {
            if (percent) {
                return v.name + "," + (v.data.value * 100).toFixed(1) + "%";
            } else {
                return `${v.name},${formatNumber(v.value)}`;
            }
        };
        if (!chart) {
            chart = init(dom);
        }
        let feature: any = {
            // saveAsImage: {
            //     show: true,
            // },
            // dataView: {
            //     show: true,
            //     readOnly: true,
            // },
        };
        if (onShareClick) {
            feature.myShareImage = {
                show: true,
                title: "分享",
                icon: `path://${ICON_SHARE}`,
                onclick: (_: any, extension: any) => {
                    onShareClick(extension);
                },
            };
        }
        let option: any = {
            backgroundColor: "white",
            animation: false,
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
                feature: feature,
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
        return chart;
    }
</script>

<script lang="ts">
    import { init, type EChartsType } from "echarts";
    import type { datastore } from "../common/datastore";
    import { ICON_SHARE, formatNumber } from "../common/utils";
    import ShareDailog from "./ShareDailog.svelte";

    export let data: datastore.Item[];
    export let showOther: boolean;
    export let percent: boolean;

    let chartDom: HTMLElement;
    let chart: EChartsType;
    let shareImgContent: string;
    let shareDialog: ShareDailog;

    const renderChart = () => {
        chart = createChart(chart, chartDom, data, showOther, percent, () => {
            let dom = document.createElement("canvas");
            dom.width = 1000;
            dom.height = 1000;
            createChart(null, dom, data, showOther, percent, null);
            setTimeout(() => {
                shareImgContent = dom.toDataURL();
                shareDialog.show();
            }, 100);
        });
    };

    window.addEventListener("resize", () => {
        chart?.resize();
    });

    $: if (chartDom && data && showOther != null) {
        renderChart();
    }
</script>

<div bind:this={chartDom} class="chart_dom mt-3"></div>
<ShareDailog imgContent={shareImgContent} bind:this={shareDialog} />
