import { fetch } from "services";
import { TUserMeta } from "shared/types";

export const create = (user: TUserMeta) =>
  fetch.post<TUserMeta>("/users", user);

export const update = (user: TUserMeta) => fetch.put(`/users/${user.id}`, user);

export const remove = (id: number) => fetch.delete(`/users/${id}`);

export const list = () =>
  fetch.get<{ users: TUserMeta[]; count: number }>("/users");
