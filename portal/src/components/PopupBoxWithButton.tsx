interface PopupBoxWithButtonProps {
  content: string;
  buttonText: string;
  onClick: () => void;
  onClose: () => void;
}

export function PopupBoxWithButton({ content, buttonText, onClick, onClose }: PopupBoxWithButtonProps) {
  return (
    <div className="fixed inset-0 z-40 flex items-start justify-center bg-black/40">
      <div
        className="
          absolute top-[33vh] -translate-y-1/2
          bg-white rounded-2xl shadow-2xl border border-gray-200
          p-8 w-full max-w-md animate-fade-in
          flex flex-col items-center justify-center text-center
          relative
        "
      >
        <button
          onClick={onClose}
          aria-label="Close popup"
          className="
            absolute top-4 right-4 text-gray-500 hover:text-gray-700 transition
            w-6 h-6 flex items-center justify-center"
        >
          <svg viewBox="0 0 20 20" fill="currentColor" className="w-full h-full">
            <path
              fillRule="evenodd"
              d="M6.293 6.293a1 1 0 011.414 0L10 8.586l2.293-2.293a1 1 0 
                 111.414 1.414L11.414 10l2.293 2.293a1 1 0 
                 01-1.414 1.414L10 11.414l-2.293 2.293a1 1 0 
                 01-1.414-1.414L8.586 10 6.293 7.707a1 1 0 010-1.414z"
              clipRule="evenodd"
            />
          </svg>
        </button>

        <div className="mb-4 text-gray-800">{content}</div>

        <button
          onClick={onClick}
          className="
            mt-2 px-4 py-2 bg-green-600 text-white rounded-lg
            hover:bg-green-700 transition"
        >
          {buttonText}
        </button>
      </div>
    </div>
  );
}
