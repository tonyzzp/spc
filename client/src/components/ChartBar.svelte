<script lang="ts" context="module">
    function createChart(
        chart: EChartsType,
        dom: HTMLElement,
        data: datastore.Item[],
        showOther: boolean,
        percent: boolean,
        formatter: any,
        onShareClick: (extension: any) => void,
    ) {
        if (!chart) {
            chart = init(dom);
        }
        let newData = data.map((v) => {
            return { ...v };
        });
        let total = newData.reduce((rtn, current) => rtn + current.value, 0);
        if (showOther) {
            let otherTotal = 0;
            for (let i = newData.length - 1; i >= 0; i--) {
                let item = newData[i];
                if (item.value / total < 0.05) {
                    newData.splice(i, 1);
                    otherTotal += item.value;
                }
            }
            if (otherTotal) {
                newData.push({
                    name: "others",
                    type: "others",
                    value: otherTotal,
                });
            }
        }
        newData.sort((a, b) => a.value - b.value);
        if (percent) {
            newData.forEach((v) => {
                v.value = v.value / total;
            });
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
        chart.setOption({
            backgroundColor: "white",
            animation: false,
            title: {
                show: true,
                text: percent ? "资产占比" : "资产详情",
                left: "center",
            },
            yAxis: {
                type: "category",
                data: newData.map((v) => v.name),
            },
            xAxis: {
                type: "value",
            },
            tooltip: {
                show: false,
            },
            toolbox: {
                show: true,
                feature: feature,
            },
            series: [
                {
                    type: "bar",
                    data: newData.map((v) => v.value),
                    label: {
                        show: true,
                        formatter: formatter,
                    },
                },
            ],
        });
        return chart;
    }
</script>

<script lang="ts">
    import { init, type EChartsType } from "echarts/core";
    import type { datastore } from "../common/datastore";
    import { ICON_SHARE, formatNumber } from "../common/utils";
    import ShareDailog from "./ShareDailog.svelte";

    export let data: datastore.Item[];
    export let percent: boolean;
    export let showOther: boolean;

    let chartDom: HTMLElement;
    let chart: EChartsType;
    let imgContent = "";
    let shareDialog: ShareDailog;

    let formatter = function (v: any) {
        if (percent) {
            return (v.value * 100).toFixed(1) + "%";
        } else {
            return `${formatNumber(v.value)}`;
        }
    };

    $: if (chartDom && data && showOther != null) {
        chart = createChart(
            chart,
            chartDom,
            data,
            showOther,
            percent,
            formatter,
            (extension) => {
                let dom = document.createElement("canvas");
                dom.width = 1000;
                dom.height = 1000;
                dom.style.backgroundColor = "white";
                createChart(
                    null,
                    dom,
                    data,
                    showOther,
                    percent,
                    formatter,
                    null,
                );
                setTimeout(() => {
                    imgContent = dom.toDataURL();
                    shareDialog.show();
                }, 100);
            },
        );
    }
</script>

<div bind:this={chartDom} class="chart_dom mt-3"></div>
<ShareDailog {imgContent} bind:this={shareDialog} />
