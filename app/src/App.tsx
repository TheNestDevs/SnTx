import "./App.css";
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
      </div>
    </div>
  );
}

export default App;
