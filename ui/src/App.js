import { useEffect } from "react";
import { Route, Routes } from "react-router-dom";
import HomePage from "./pages/home";
import VehiclesPage from "./pages/vehicles";
import VehicleDetailPage from "./pages/vehicleDetail";

const App = () => {
  useEffect(() => {
    document.title = "Grease Chimp";
  }, []);

  let content = (
    <>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/vehicles" element={<VehiclesPage />} />
        <Route path="/vehicleDetail/:vin" Component={VehicleDetailPage} />
      </Routes>
    </>
  );

  return content;
}

export default App;
