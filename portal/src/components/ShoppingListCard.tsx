import { Link } from "react-router-dom";
import type { ShoppingList } from "../models/ShoppingList";
import { Users, User2, ListTodo } from "lucide-react";

export default function ShoppingListCard({ list }: { list: ShoppingList }) {
  return (
    <Link to={`/list/${list.id}`}>
      <div className="p-5 bg-white rounded-xl shadow-sm hover:shadow-md hover:bg-gray-50 transition-all border border-gray-200">
        <div className="flex items-center justify-between mb-3">
          <h2 className="text-lg font-semibold text-gray-800 truncate">{list.name}</h2>
          <div className="flex items-center gap-1 text-sm text-gray-500">
            <ListTodo size={16} />
            <span>{list.items?.length ?? 0}</span>
          </div>
        </div>

        <div className="flex justify-between text-sm text-gray-500">
          <div className="flex items-center gap-1">
            <User2 size={14} />
            <span>{list.owner?.username || "Unknown"}</span>
          </div>

          <div className="flex items-center gap-1">
            <Users size={14} />
            <span>{list.users?.length ?? 0}</span>
          </div>
        </div>
      </div>
    </Link>
  );
}
