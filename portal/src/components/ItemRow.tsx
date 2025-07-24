import { useState } from "react";
import type { Item } from "../models/Item";
import { CheckCircle, Circle } from "lucide-react"; // Icons (if using lucide)
import api from "../api/axios";

interface Props {
  item: Item;
}

export default function ItemRow({ item }: Props) {
  const [bought, setBought] = useState(item.bought ?? false);

  const toggleBought = async () => {
    setBought(prev => !prev);
    try {
      await api.put("/api/putItem", { ...item, bought: !bought })
    } catch (err) {
      console.error("Failed to put item:", err);
    }

  };

  return (
    <li
      onClick={toggleBought}
      className={`flex items-center justify-between p-3 rounded-lg border shadow-sm transition-all cursor-pointer 
        ${bought ? "bg-green-100 text-gray-500 line-through" : "bg-white hover:bg-gray-50"}`}
    >
      <div className="flex items-center gap-3">
        {bought ? (
          <CheckCircle className="text-green-500" size={20} />
        ) : (
          <Circle className="text-gray-400" size={20} />
        )}
        <div>
          <div className="font-medium">{item.name}</div>
          <div className="text-sm text-gray-500">{item.amount}</div>
        </div>
      </div>
    </li>
  );
}
