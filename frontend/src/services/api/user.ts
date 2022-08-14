import { fetch } from "services";
import {TUser} from "../../shared/types";

export const create = (user: TUser) =>
  fetch.post<TUser>("/users", user);

export const update = (user: TUser) => fetch.put(`/users/${user.id}`, user);

export const remove = (id: string) => fetch.delete(`/users/${id}`);

export const list = () =>
  fetch.get<{ users: TUser[]; count: number }>("/users");
