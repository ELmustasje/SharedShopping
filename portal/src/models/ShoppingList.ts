import type { Item } from "./Item";
import type { User } from "./User";
import { buildItem } from "./Item"
import { buildUser } from "./User"

export type ShoppingList = {
  id: number;
  name: string;
  ownerID: number;
  owner: User;
  users: User[];
  items: Item[];
  customItems: Item[];
};


export function buildShoppingList(raw: any): ShoppingList {
  return {
    id: raw.ID,
    name: raw.name,
    ownerID: raw.OwnerID,
    owner: buildUser(raw.Owner),
    users: (raw.Users || []).map(buildUser),
    items: (raw.Items || []).map(buildItem),
    customItems: (raw.CustomItems || []).map(buildItem),
  }
}
