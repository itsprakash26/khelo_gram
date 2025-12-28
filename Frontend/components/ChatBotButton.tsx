const ChatBotButton = () => {
  return (
    <button
      className="
        fixed 
        bottom-6 
        right-6 
        z-50 
        h-14 
        w-14 
        rounded-full 
        bg-primary 
        flex 
        items-center 
        justify-center 
        shadow-xl 
        transition-all
        duration-300
        hover:scale-110
        hover:shadow-2xl
        active:scale-95
      "
      aria-label="Chat Bot"
    >
      <img
        className="
          h-full 
          w-full 
          object-cover 
          rounded-full
          transition-transform
          duration-300
          hover:rotate-8
        "
        src="/chatbot-icon.png"
        alt="Chat Bot"
      />
    </button>
  );
};

export default ChatBotButton;
