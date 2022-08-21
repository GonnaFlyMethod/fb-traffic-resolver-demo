import {fetch} from "services";
import {TUserWithID, TUserWithoutID} from "../../shared/types";

export const create = (user: TUserWithoutID) =>
    fetch.post<TUserWithoutID>("/api/users", user);

export const update = (userID: string, user: TUserWithoutID) => fetch.put(`/api/users/${userID}`, user);

export const remove = (userID: string) => fetch.delete(`/api/users/${userID}`);

export const list = () =>
    fetch.get<{ users: TUserWithID[]; count: number }>("/api/users");
