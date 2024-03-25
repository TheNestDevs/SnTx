import { MenuIcon } from "lucide-react";
import { GithubLogo } from "../atoms/githubLogo";

const Header = () => {
  return (
    <div className="w-full h-10 flex justify-between font-poppins ">
      <div className="text-3xl flex flex-col justify-center font-semibold">
        SnTx
      </div>
      <div className="flex flex-col justify-center sm:hidden">
        <MenuIcon size={24} />
      </div>
      <div className="h-full hidden sm:flex flex-col justify-center">
        <div className="flex text-xl gap-16 font-medium">
          <a
            href="/about"
            className="cursor-pointer hover:text-accent transition-colors ease-in-out flex flex-col justify-center"
          >
            About
          </a>
          <a
            href="https://github.com/TheNestDevs/SnTx"
            className="flex flex-col justify-center w-12 h-12 hover:bg-[#ffffff22] transition-colors ease-in-out rounded-full"
            target="_blank"
          >
            <GithubLogo className="mx-auto" />
          </a>
        </div>
      </div>
    </div>
  );
};

export default Header;
