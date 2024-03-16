import { ArrowDown, ArrowUp } from "lucide-react";
import "./App.css";
import ActionButton from "./components/atoms/actionButton";
import Blur from "./components/atoms/blur";
import Header from "./components/molecules/header";
import Hero from "./components/molecules/hero";

function App() {
  return (
    <div className="w-[100vw]">
      <Blur />
      <div className="sm:mx-20 mx-6 sm:mt-14 mt-6 font-poppins">
        <Header />
        <Hero className="sm:mt-24 mt-20" />
        <div className="flex sm:flex-row flex-col sm:bg-transparent bg-card rounded-2xl sm:py-0 py-24 sm:px-0 px-11 justify-center sm:mt-24 mt-7 gap-[20vw]">
          <ActionButton>
            <div className="flex gap-2">
              Send
              <ArrowUp className="sm:group-hover:-translate-y-1 transition-transform ease-in-out" />
            </div>
          </ActionButton>
          <ActionButton>
            <div className="flex gap-2">
              Receive
              <ArrowDown className="sm:group-hover:translate-y-1 transition-transform ease-in-out" />
            </div>
          </ActionButton>
        </div>
      </div>
    </div>
  );
}

export default App;
