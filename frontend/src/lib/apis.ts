import axios, { AxiosRequestConfig } from "axios";
import {data} from "autoprefixer";

const baseUrl = 'http://127.0.0.1:8000'

class Apis {
    async postTodo(req: postTodoRequest, axiosconfig?: AxiosRequestConfig) {
        const data = new FormData()
        data.append('title', req.title)
        await axios.post<void>(`${baseUrl}/api/todo`, data, {
            headers: {'content-type': 'multipart/form-data'},
            ...axiosconfig
        })
    }

    async getTodo(axiosconfig?: AxiosRequestConfig) {
        const { data } = await axios.get(`${baseUrl}/api/todos`, axiosconfig)

        return data
    }
}

const apis = new Apis()
export default apis

export interface postTodoRequest {
    title: string
}

export interface GetTodoResponse {
    data: string[]
}
