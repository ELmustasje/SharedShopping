import { useState, useRef, useEffect } from "react";
import { Bars3Icon } from "@heroicons/react/24/outline";

interface DropdownMenuProps {
  options: { label: string; onClick: () => void }[];
}

export default function DropdownMenu({ options }: DropdownMenuProps) {
  const [open, setOpen] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handler = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handler);
    return () => document.removeEventListener("mousedown", handler);
  }, []);

  return (
    <div className="relative inline-block text-left" ref={ref}>
      <button
        onClick={() => setOpen((prev) => !prev)}
        className="p-2 rounded-full hover:bg-gray-100 transition"
        aria-label="Menu"
      >
        <Bars3Icon className="w-6 h-6 text-gray-600" />
      </button>

      {open && (
        <div className="absolute right-0 z-10 mt-2 w-48 bg-white rounded-md shadow-lg ring-1 ring-black ring-opacity-5">
          {options.map((option, idx) => (
            <button
              key={idx}
              onClick={() => {
                option.onClick();
                setOpen(false);
              }}
              className="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
            >
              {option.label}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
