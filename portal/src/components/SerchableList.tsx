import { useState } from "react"

interface Props {
  master_food_list: string[]
}
export default function SearchableList({ master_food_list }: Props) {
  const [search, setSearch] = useState("")

  const filtered = master_food_list.filter((item) => item.toLowerCase().includes(search.toLowerCase()))

  return (
    <div className="max-h-80 overflow-y-auto p-2 border rounded">
      <input
        type="text"
        className="w-full mb-2 p-2 border rounded"
        placeholder="Search..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />
      <ul className="space-y-1">
        {filtered.slice(0, 10).map((item, idx) => (
          <li key={idx} className="p-2 bg-gray-100 rounded">{item}</li>
        ))}
      </ul>
    </div>
  )
}
