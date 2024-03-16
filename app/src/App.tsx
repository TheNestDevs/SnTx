import "./App.css";
import Blur from "./components/atoms/blur";
import Header from "./components/molecules/header";

function App() {
  return (
    <div className="w-[100vw]">
      <Blur />
      <div className="sm:mx-20 mx-6 sm:mt-14 mt-6 font-poppins">
        <Header />
      </div>
    </div>
  );
}

export default App;
