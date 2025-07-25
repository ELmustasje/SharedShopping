import { useState } from "react";
import api from "../api/axios";

interface AddShoppinglistProps {
  onClose: () => void;
  onCreated: () => void;
}

export default function NewListModal({ onClose, onCreated }: AddShoppinglistProps) {
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async () => {
    if (!name.trim()) {
      setError("List name is required");
      return;
    }

    try {
      setLoading(true);
      await api.post("/api/postList", { name: name.trim() });
      onCreated();
      onClose();
    } catch {
      setError("Failed to create list.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 z-40 flex items-center justify-center bg-black/40">
      <div className="bg-white rounded-2xl shadow-2xl p-8 w-full max-w-md animate-fade-in border border-gray-200">
        <h2 className="text-2xl font-bold text-gray-800 mb-4 text-center">
          Create a New Shopping List
        </h2>

        <input
          type="text"
          placeholder="Enter list name..."
          className="w-full px-4 py-3 border border-gray-300 rounded-lg mb-4 focus:outline-none focus:ring-2 focus:ring-green-500"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />

        {error && <p className="text-red-600 text-sm mb-4">{error}</p>}

        <div className="flex justify-end space-x-3">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-gray-100 hover:bg-gray-200 text-sm rounded-md"
            disabled={loading}
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            className="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-md hover:bg-green-700"
            disabled={loading}
          >
            {loading ? "Creating..." : "Create"}
          </button>
        </div>
      </div>
    </div>
  );
}
