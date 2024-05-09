import { api } from "./api"

const KEY = "stock-portfolio-chart-data"

export namespace datastore {

    export interface Item {
        name: string
        value: number
        type: string
    }

    export async function load(): Promise<Item[]> {
        let result = await api.load()
        if (result == null || result.code) {
            throw result.msg || "加载数据失败"
        }
        if (result.data.content == null || result.data.content == "") {
            return []
        }
        try {
            return JSON.parse(result.data.content)
        } catch (e) {
            throw "服务器上保存的数据格式不正确"
        }
    }

    export async function save(items: Item[]) {
        return api.save(JSON.stringify(items))
    }
}