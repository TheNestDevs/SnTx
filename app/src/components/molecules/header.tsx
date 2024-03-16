import { Github } from "lucide-react";

const Header = () => {
  return (
    <div className="w-full h-10 flex justify-between">
      <div className="text-4xl font-bold">SnTx</div>
      <div className="h-full hidden sm:flex flex-col justify-center">
        <div className="flex text-2xl gap-16 font-medium">
          <a href="/about" className="cursor-pointer">
            About
          </a>
          <a href="/register" className="flex flex-col justify-center">
            <Github size={24} />
          </a>
        </div>
      </div>
    </div>
  );
};

export default Header;
