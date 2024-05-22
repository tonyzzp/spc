export namespace etypes {

    export interface Item {
        name: string
        value?: number
        children?: Item[]
    }

    export type Data = Item[]

    export interface Label {
        show: boolean
        /**
         * {b}: {c}
         */
        formatter?: any
    }

    export interface Series {
        type: "sunburst",
        nodeClick?: any
        data: Data
        label: Label
        emphasis: {
            focus: string,
        },
    }

    export interface Tooltip {
        show: boolean
    }

    export interface Option {
        title: {
            show: true,
            text: string,
            left: string | number,
        },
        tooltip?: Tooltip
        toolbox?: {
            show: boolean,
            feature?: {
                saveAsImage?: {
                    show: true,
                },
                dataView?: {
                    show: true,
                    readOnly: true,
                },
            },
        },
        series: Series
    }
}