import { useEffect, useState } from "react";
import { ClipboardList, PlusCircle } from "lucide-react";
import ShoppingListCard from "../components/ShoppingListCard";
import AddShoppinglist from "../components/AddShoppinglist"

import { buildShoppingList, type ShoppingList } from "../models/ShoppingList";
import api from "../api/axios";
import UserBadge from "../components/UserBadge";
import { PopupBoxWithButton } from "../components/PopupBoxWithButton";

export default function HomePage() {
  const [lists, setLists] = useState<ShoppingList[]>([]);
  const [loading, setLoading] = useState(true);
  const [showAddList, setShowAddList] = useState(false)
  const [inviteToken, setInviteToken] = useState("wadaw")


  useEffect(() => {
    fetchLists();
  }, []);

  const fetchLists = async () => {
    try {
      const response = await api.get(`/api/getUserLists`);
      const rawLists = response.data;
      const builtLists = rawLists.map((raw: any) => buildShoppingList(raw));
      setLists(builtLists);
    } catch (error) {
      console.error("Failed to fetch shopping lists:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreated = async () => {
    await fetchLists()
  }

  if (loading) {
    return (
      <main className="max-w-5xl mx-auto p-6 flex justify-center items-center min-h-[60vh]">
        <div className="animate-spin rounded-full h-16 w-16 border-b-2 border-green-600"></div>
      </main>
    );
  }

  return (
    <main className="max-w-5xl mx-auto p-6">
      <UserBadge />
      <header className="flex justify-between items-center mb-8">
        <h1 className="text-4xl font-bold text-gray-900 flex items-center gap-3 mb-6">
          <ClipboardList size={36} className="text-green-600 animate-pulse" />
          Your Shopping Lists
        </h1>

        <button
          className="flex items-center gap-2 bg-green-600 text-white px-5 py-3 rounded-lg shadow hover:bg-green-700 transition"
          aria-label="Create new shopping list"
          onClick={() => setShowAddList(true)}
        >
          <PlusCircle size={20} />
          New List
        </button>
      </header>

      {lists.length === 0 ? (
        <section className="text-center text-gray-500 mt-32">
          <img
            src="/empty-state-illustration.svg"
            alt="No shopping lists"
            className="mx-auto mb-6 w-48 opacity-70"
          />
          <p className="text-xl font-semibold mb-2">No shopping lists found</p>
          <p className="text-md mb-4">
            Create your first list to get started managing your groceries easily.
          </p>
          <button
            className="inline-flex items-center gap-2 bg-green-600 text-white px-6 py-3 rounded-lg shadow hover:bg-green-700 transition"
            onClick={() => alert("Open modal or navigate to create new list")}
          >
            <PlusCircle size={20} />
            Create List
          </button>
        </section>
      ) : (
        <section className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          {lists.map((list) => (
            <ShoppingListCard key={list.id} list={list} />
          ))}
        </section>

      )}
      <div className="relative inline-block">
        {showAddList && (
          <AddShoppinglist onClose={() => setShowAddList(false)} onCreated={handleCreated} />
        )}
      </div>
      <div className="relative inline-block">
        {inviteToken && (
          <PopupBoxWithButton content="test" buttonText="test2" onClose={() => setInviteToken("")} onClick={() => setInviteToken("")} />
        )}
      </div>

    </main>
  );
}
