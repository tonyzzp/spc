import { BarChart, SunburstChart } from "echarts/charts";
import { GridComponent, TitleComponent, ToolboxComponent, TooltipComponent } from "echarts/components";
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import type { datastore } from "./datastore";

export const ICON_SHARE =
    "M11 2.5a2.5 2.5 0 1 1 .603 1.628l-6.718 3.12a2.499 2.499 0 0 1 0 1.504l6.718 3.12a2.5 2.5 0 1 1-.488.876l-6.718-3.12a2.5 2.5 0 1 1 0-3.256l6.718-3.12A2.5 2.5 0 0 1 11 2.5z";

export const ICON_SAVE =
    "M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5 M7.646 11.854a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708L8.5 10.293V1.5a.5.5 0 0 0-1 0v8.793L5.354 8.146a.5.5 0 1 0-.708.708z"


export function formatNumber(v: number) {
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

export function sortData(data: datastore.Item[]) {
    data.sort((a, b) => b.value - a.value)
}


export function initEcharts() {
    use([
        CanvasRenderer,
        SunburstChart,
        BarChart,
        TitleComponent,
        TooltipComponent,
        ToolboxComponent,
        GridComponent,
    ])
}