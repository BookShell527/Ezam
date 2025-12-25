import { Context, createContext } from "react";

interface User {
    id?: string,
    status?: string,
}

const userContext: Context<User> = createContext({});
