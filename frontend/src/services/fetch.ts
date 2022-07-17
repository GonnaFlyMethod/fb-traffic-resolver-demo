
import axios, { AxiosInstance, AxiosRequestConfig } from "axios";

import { API_URL } from "shared/consts";

import { extractDataInterceptor } from "./interceptors";

export type FetchServiceConfig = {
  apiUrl: string;
  photoUrl: string;
};

export class FetchService {
  private instance: AxiosInstance;

  constructor() {
    this.instance = axios.create();

    this.instance.interceptors.response.use(extractDataInterceptor);
  }

  init() {
    this.instance.defaults.baseURL = API_URL;
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
