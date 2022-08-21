import {makeAutoObservable} from "mobx";

import {API} from "services";
import LoadingModel from "models/Loading";
import {TUserWithID, TUserWithoutID} from "../../../shared/types";

class UsersModel {
    private _users: TUserWithID[] = [];

    loading: LoadingModel;

    constructor() {
        makeAutoObservable(this);

        this.loading = new LoadingModel();
    }

    set users(data: TUserWithID[]) {
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

    async create(user: TUserWithoutID) {
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

    async update(userID: string, user: TUserWithoutID) {
        this.loading.begin();

        await API.user.update(userID, user);
        await this.fetch();

        this.loading.end();
    }
}

export default new UsersModel();
