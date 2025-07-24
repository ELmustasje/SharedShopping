import { createContext, useContext } from "react";
import type { User } from "../models/User"; // adjust path as needed

export const UserContext = createContext<User | null>(null);

export function useUser() {
  return useContext(UserContext);
}
