import { useUser } from "../contexts/userContext.tsx"
import { User } from "lucide-react";


export default function UserBadge() {
  const user = useUser();

  if (!user) return null;

  return (
    <div
      className="fixed top-4 right-4 flex items-center gap-2 bg-white border border-gray-200 rounded-full px-4 py-2 shadow-md hover:shadow-lg transition-shadow text-sm text-gray-700"
    >
      <div className="bg-green-100 text-green-600 p-1 rounded-full">
        <User size={16} />
      </div>
      <span className="font-medium truncate max-w-[120px]">
        {user.username}
      </span>
    </div>
  );
}
