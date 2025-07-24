export type Item = {
  id: number;
  name: string;
  bought?: boolean;
  amount: number;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string | null;
};

export function buildItem(raw: any): Item {
  return {
    id: raw.ID,
    name: raw.name,
    bought: raw.bought,
    amount: raw.amount,
    createdAt: raw.CreatedAt,
    updatedAt: raw.UpdatedAt,
    deletedAt: raw.DeletedAt,
  };
}

