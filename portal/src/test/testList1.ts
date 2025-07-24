import type { Item } from "../models/Item"
import type { ShoppingList } from "../models/ShoppingList"
import type { User } from "../models/User"

const testOwner: User = {

  id: 1,
  userName: "alice",
  ownedLists: [], // can be filled later
  sharedLists: [],
}

const testUser2: User = {
  id: 2,
  userName: "bob",
  ownedLists: [],
  sharedLists: [],
}

const regularItems: Item[] = [
  {
    id: 101,
    name: "Milk",
    bought: false,
    createdAt: "2025-07-01T10:00:00Z",
    updatedAt: "2025-07-01T10:00:00Z",
    deletedAt: null,
  },
  {
    id: 102,
    name: "Bread",
    bought: true,
    createdAt: "2025-07-01T10:00:00Z",
    updatedAt: "2025-07-02T09:00:00Z",
    deletedAt: null,
  },
]

const customItems: Item[] = [
  {
    id: 201,
    name: "Local Honey",
    bought: false,
    createdAt: "2025-07-02T12:00:00Z",
    updatedAt: "2025-07-02T12:00:00Z",
    deletedAt: null,
  },
]

export const testList: ShoppingList = {
  id: 1,
  name: "Weekly Groceries",
  ownerID: testOwner.id,
  owner: testOwner,
  users: [testUser2], // shared with Bob
  items: regularItems,
  customItems: customItems,
}

// Optionally populate owned/shared lists:
testOwner.ownedLists = [testList]
testUser2.sharedLists = [testList]
