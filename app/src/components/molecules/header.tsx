import { Github, MenuIcon } from "lucide-react";

const Header = () => {
  return (
    <div className="w-full h-10 flex justify-between">
      <div className="text-4xl font-bold">SnTx</div>
      <div className="flex flex-col justify-center sm:hidden">
        <MenuIcon size={24} />
      </div>
      <div className="h-full hidden sm:flex flex-col justify-center">
        <div className="flex text-2xl gap-16 font-medium">
          <a href="/about" className="cursor-pointer">
            About
          </a>
          <a
            href="https://github.com/TheNestDevs/SnTx"
            className="flex flex-col justify-center"
            target="_blank"
          >
            <Github size={24} />
          </a>
        </div>
      </div>
    </div>
  );
};

export default Header;
