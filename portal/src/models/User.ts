import type { ShoppingList } from "./ShoppingList";

export type User = {
  id: number;
  username: string;
  ownedLists: ShoppingList[];
  sharedLists: ShoppingList[];
};


export function buildUser(raw: any): User {
  return {
    id: raw.ID,
    username: raw.username,
    ownedLists: [], // not implemented 
    sharedLists: [],// not implemented
  }
}
