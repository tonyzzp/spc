<script lang="ts">
    import { init, type EChartsType } from "echarts";
    import type { datastore } from "./datastore";

    export let data: datastore.Item[];
    export let percent: boolean;
    export let showOther: boolean;

    let chartDom: HTMLElement;
    let chart: EChartsType;

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

    let formatter = function (v: any) {
        if (percent) {
            return (v.value * 100).toFixed(1) + "%";
        } else {
            return `${formatNumber(v.value)}`;
        }
    };

    $: if (chartDom && data && showOther != null) {
        if (!chart) {
            chart = init(chartDom);
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
        chart.setOption({
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
    }
</script>

<div bind:this={chartDom} class="chart_dom mt-3"></div>

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
