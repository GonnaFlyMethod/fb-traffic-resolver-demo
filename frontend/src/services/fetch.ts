import axios, {AxiosInstance, AxiosRequestConfig} from "axios";

import {extractDataInterceptor} from "./interceptors";

export class FetchService {
    private instance: AxiosInstance;

    constructor() {
        this.instance = axios.create();

        this.instance.interceptors.response.use(extractDataInterceptor);
    }

    public post<T>(
        url: string,
        data?: any,
        config?: AxiosRequestConfig
    ): Promise<T> {
        return this.instance.post(url, data, config);
    }

    public put<T>(
        url: string,
        data?: any,
        config?: AxiosRequestConfig
    ): Promise<T> {
        return this.instance.put(url, data, config);
    }

    public patch<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return this.instance.patch(url, config);
    }

    public get<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return this.instance.get(url);
    }

    public delete<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return this.instance.delete(url, config);
    }
}

export default new FetchService();
