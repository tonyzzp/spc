import md5 from "md5-ts"

const SIGN_KEY = "15998569"

export namespace api {

    let _token = ""

    export interface LoginParam {
        userName: string
        password: string
    }

    export interface ChangePasswordParam {
        oldPassword: string,
        newPassword: string
    }

    export interface BaseResult {
        code: number
        msg: string
    }

    export interface LoginResult extends BaseResult {
        data: {
            token: string
        }
    }

    export interface RegResult extends BaseResult {
    }

    export interface SaveResult extends BaseResult {

    }

    export interface LoadResult extends BaseResult {
        data: {
            content: string
        }
    }

    function sign(body: string) {
        let signStr = `${SIGN_KEY}.${body}.${SIGN_KEY}`
        return md5(signStr)
    }

    async function post(api: string, body: any, headers?: Record<string, string>) {
        let sbody = typeof body == "string" ? body : JSON.stringify(body)
        try {
            let resp = await fetch(api, {
                method: "POST",
                body: sbody,
                headers: {
                    ...(headers || {}),
                    Sign: sign(sbody),
                    Token: _token,
                }
            })
            return await resp.json()
        } catch (e) {
            return null
        }
    }

    export function setToken(token: string) {
        _token = token
    }

    export function getUserName() {
        if (!_token) {
            return null
        }
        return _token.split(".")[0]
    }

    export function reg(userName: string, password: string): Promise<RegResult> {
        let param: LoginParam = {
            userName: userName,
            password: password,
        }
        return post("/api/reg", param)
    }

    export function login(userName: string, password: string): Promise<LoginResult> {
        let param: LoginParam = {
            userName: userName,
            password: password,
        }
        return post("/api/login", param)
    }

    export function changePassword(oldPassword: string, newPassword: string): Promise<BaseResult> {
        let param: ChangePasswordParam = {
            oldPassword: oldPassword,
            newPassword: newPassword
        }
        return post("/api/changePassword", param)
    }

    export function save(content: string): Promise<SaveResult> {
        return post("/api/save", content)
    }

    export function load(): Promise<LoadResult> {
        return post("/api/load", "")
    }
}