import { useState, useRef, useEffect } from "react";
import api from "../api/axios";
import { PlusIcon } from "@heroicons/react/24/outline";

interface Props {
  listId: number;
  onAdd: () => void;
  masterFoodList: string[];
}

export default function AddItemForm({ listId, onAdd, masterFoodList }: Props) {
  const [name, setName] = useState("");
  const [amount, setAmount] = useState("");
  const [loading, setLoading] = useState(false);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [open, setOpen] = useState(false);

  const filteredFoods = masterFoodList.filter(food =>
    food.toLowerCase().includes(name.toLowerCase()) && name.trim() !== ""
  );

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;

    setLoading(true);

    try {
      await api.post("/api/postItem", {
        name: name.trim(),
        amount: amount.trim(),
        shoppingListId: listId,
      });
      onAdd();
      setName("");
      setAmount("");
      setShowSuggestions(false);
      setOpen(false)
    } catch (err) {
      console.error("Failed to add item:", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="relative">
      {/* The plus icon */}
      <button
        onClick={() => setOpen(prev => !prev)}
        className="p-2 rounded-full hover:bg-gray-100 transition"
        aria-label="Add Item"
        type="button"
      >
        <PlusIcon className="w-6 h-6 text-green-600" />
      </button>

      {/* The form itself */}
      {open && (
        <form
          onSubmit={handleSubmit}
          className="absolute right-0 mt-2 z-20 bg-white p-4 rounded-xl shadow-xl w-96"
        >
          {/* Item name */}
          <div className="mb-4 relative">
            <label className="block text-sm font-medium text-gray-700 mb-1">Item name</label>
            <input
              value={name}
              onChange={e => {
                setName(e.target.value);
                setShowSuggestions(true);
              }}
              placeholder="e.g. Bananas"
              className="w-full border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500"
              onFocus={() => setShowSuggestions(true)}
              onBlur={() => setTimeout(() => setShowSuggestions(false), 150)}
            />
            {showSuggestions && filteredFoods.length > 0 && (
              <ul className="absolute z-30 mt-1 max-h-60 w-full overflow-y-auto bg-white border border-gray-200 rounded-lg shadow-md">
                {filteredFoods.map((food, idx) => (
                  <li
                    key={idx}
                    onClick={() => {
                      setName(food);
                      setShowSuggestions(false);
                    }}
                    className="px-3 py-2 hover:bg-green-100 cursor-pointer"
                  >
                    {food}
                  </li>
                ))}
              </ul>
            )}
            {showSuggestions && filteredFoods.length === 0 && (
              <div className="absolute z-30 mt-1 w-full bg-white border border-gray-200 rounded-lg shadow-md px-3 py-2 text-sm text-gray-500">
                No matches found
              </div>
            )}
          </div>

          {/* Amount */}
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700 mb-1">Amount</label>
            <input
              value={amount}
              onChange={e => setAmount(e.target.value)}
              placeholder="e.g. 6 pcs"
              className="w-full border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          {/* Submit */}
          <button
            type="submit"
            disabled={loading}
            className="bg-green-500 hover:bg-green-600 text-white font-medium py-2 px-4 rounded-lg transition-all disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {loading ? "Adding..." : "Add"}
          </button>
        </form>
      )}
    </div>
  );
}
