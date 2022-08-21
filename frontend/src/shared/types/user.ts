export interface TUserWithoutID {
    name: string;
    surname: string;
    email: string;
}

export interface TUserWithID extends TUserWithoutID {
    id: string;
}


