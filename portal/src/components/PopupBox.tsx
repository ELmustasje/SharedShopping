import { useEffect, useRef } from "react";

interface PopupBoxProps {
  content: string;
  onClose: () => void;
}

export function PopupBox({ content, onClose }: PopupBoxProps) {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClick = (e: MouseEvent | TouchEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        onClose();
      }
    };

    document.addEventListener("mousedown", handleClick, true);
    document.addEventListener("touchstart", handleClick, true);
    return () => {
      document.removeEventListener("mousedown", handleClick, true);
      document.removeEventListener("touchstart", handleClick, true);
    };
  }, [onClose]);

  return (
    <div className="fixed inset-0 z-40 flex items-start justify-center bg-black/40">
      <div
        ref={ref}
        className="absolute top-[33vh] -translate-y-1/2 bg-white rounded-2xl shadow-2xl p-8 w-full max-w-md animate-fade-in border border-gray-200
                  flex flex-col items-center justify-center text-center"
      >
        {content}
      </div>
    </div>
  );
}
