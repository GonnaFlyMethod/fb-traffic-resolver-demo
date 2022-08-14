import { makeAutoObservable } from "mobx";

import { API } from "services";
import LoadingModel from "models/Loading";
import {TUser} from "../../../shared/types";

class UsersModel {
  private _users: TUser[] = [];

  loading: LoadingModel;

  constructor() {
    makeAutoObservable(this);

    this.loading = new LoadingModel();
  }

  set users(data: TUser[]) {
    this._users = data;
  }

  get users() {
    return this._users;
  }

  async fetch() {
    this.loading.begin();

    const data = await API.user.list();
    this.users = data.users;

    this.loading.end();
  }

  async create(user: TUser){
    this.loading.begin();

    await API.user.create(user);
    await this.fetch();

    this.loading.end();
  }

  async remove(id: string) {
    this.loading.begin();

    await API.user.remove(id);
    await this.fetch();

    this.loading.end();
  }

  async update(user: TUser) {
    this.loading.begin();

    await API.user.update(user);
    await this.fetch();

    this.loading.end();
  }
}

export default new UsersModel();
