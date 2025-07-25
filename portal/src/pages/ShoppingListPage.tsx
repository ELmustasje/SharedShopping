import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import ItemRow from "../components/ItemRow";
import { buildShoppingList, type ShoppingList } from "../models/ShoppingList";
import api from "../api/axios";
import UserBadge from "../components/UserBadge";
import DropdownMenu from "../components/DropdownMenu";
import AddItemForm from "../components/AddItemForm";
import { master_food_list } from "../assets/master_food_list";

export default function ShoppingListPage() {
  const { id } = useParams();
  const [list, setList] = useState<ShoppingList | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchList = async () => {
    try {
      const response = await api.get("/api/getList", {
        params: { id },
      });
      const listData = buildShoppingList(response.data);
      setList(listData);
    } catch (err) {
      setError("Could not load shopping list.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (id) fetchList();
  }, [id]);

  const handleAddItem = () => {
    fetchList()
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <span className="text-gray-500 text-lg animate-pulse">Loading list...</span>
      </div>
    );
  }

  if (error || !list) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div className="text-red-600 text-lg font-medium">{error ?? "List not found."}</div>
      </div>
    );
  }

  return (
    <div className="max-w-3xl mx-auto p-6">
      <div className="bg-white rounded-2xl shadow-md p-6">
        <UserBadge />

        <div className="flex justify-between items-center mb-4">
          <h1 className="text-2xl font-bold text-gray-800">🛒 {list.name}</h1>

          <div className="flex items-center space-x-2">
            <AddItemForm listId={list.id} onAdd={handleAddItem} masterFoodList={master_food_list} />

            <DropdownMenu
              options={[
                {
                  label: "🧹 Delete Bought Items",
                  onClick: async () => {
                    await api.delete("/api/deleteBoughtItems", {
                      params: { listId: list.id },
                    });
                    fetchList();
                  },
                },
                {
                  label: "✏️ Rename List",
                  onClick: () => {
                    const newName = prompt("Enter new list name:");
                    if (newName) {
                      api.put("/api/renameList", { id: list.id, name: newName }).then(fetchList);
                    }
                  },
                },
                {
                  label: "Add Users",
                  onClick: () => {
                    const newName = prompt("Enter new list name:");
                    if (newName) {
                      api.put("/api/renameList", { id: list.id, name: newName }).then(fetchList);
                    }
                  },
                },
              ]}
            />
          </div>
        </div>

        <ul className="mt-6 divide-y divide-gray-200">
          {list.items.length === 0 ? (
            <li className="text-gray-500 text-center py-4">No items yet. Add something!</li>
          ) : (
            list.items.map(item => (
              <li key={item.id} className="py-2">
                <ItemRow item={item} />
              </li>
            ))
          )}
        </ul>
      </div>
    </div>
  );
}
