export const ICON_SHARE =
    "M11 2.5a2.5 2.5 0 1 1 .603 1.628l-6.718 3.12a2.499 2.499 0 0 1 0 1.504l6.718 3.12a2.5 2.5 0 1 1-.488.876l-6.718-3.12a2.5 2.5 0 1 1 0-3.256l6.718-3.12A2.5 2.5 0 0 1 11 2.5z";

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