
const KEY = "stock-portfolio-chart-data"

export namespace datastore {

    let data: { [name: string]: Item }

    export interface Item {
        name: string
        value: number
        type: string
    }

    export function init() {
        let s = localStorage.getItem(KEY)
        if (s && s.length) {
            try {
                data = JSON.parse(s)
            } catch (e) {
                localStorage.removeItem(KEY)
                data = {}
            }
        } else {
            data = {}
        }
    }

    function save() {
        localStorage.setItem(KEY, JSON.stringify(data))
    }

    export function all() {
        return Object.values(data).sort((a, b) => {
            if (a.type == b.type) {
                return b.value - a.value
            }
            return a.type.localeCompare(b.type)
        })
    }

    export function put(item: Item) {
        data[item.name] = item
        save()
    }

    export function remove(name: string) {
        delete data[name]
        save()
    }

}