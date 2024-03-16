import { FCProps } from "../../types";
import { cn } from "../../utils/utils.styles";

const Hero: FCProps = ({ className }) => {
  return (
    <div className={cn(className, "w-full text-center sm:text-3xl text-2xl")}>
      <span className="text-accent">Sync</span> your content,
      <span className="sm:hidden">
        <br />
      </span>{" "}
      seamlessly across devices.
    </div>
  );
};

export default Hero;
