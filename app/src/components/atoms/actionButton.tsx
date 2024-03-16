import { FCProps } from "../../types";

const ActionButton: FCProps = ({ children }) => {
  return (
    <div className="group cursor-pointer transition-colors ease-in-out sm:w-[25vw] w-full sm:h-[20vw] h-20 text-2xl flex flex-col items-center justify-center ring-2 ring-ring sm:rounded-[3vw] rounded-full">
      {children}
    </div>
  );
};

export default ActionButton;
