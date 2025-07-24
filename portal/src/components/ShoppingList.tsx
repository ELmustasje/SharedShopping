// import { type Item } from "../models/Item.ts"
//
// type ShoppingListProps = {
//   items: Item[];
//   onSelect?: (item: Item) => void;
// };
//
//
// export default function ShoppingList({ items, onSelect }: ShoppingListProps) {
//   const sortedItems = items.slice().sort((a, b) => Number(a.bought) - Number(b.bought));
//   return (
//     <div className="w-full max-w-md mx-auto border rounded-xl shadow p-4 bg-white space-y-2">
//       {sortedItems.map((item) => (
//         < div
//           onClick={() => onSelect?.(item)}
//           className={`p-3 rounded cursor-pointer transition-colors ${item.bought
//             ? "bg-red-100 text-gray-500 line-through"
//             : "bg-gray-100 hover:bg-gray-200"
//             }`}
//         >
//           {item.name}
//         </div>
//       ))
//       }
//     </div >
//   );
// }
