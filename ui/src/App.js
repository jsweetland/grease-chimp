import { useEffect } from "react";
import { Route, Routes } from "react-router-dom";
import HomePage from "./pages/home";
import VehiclesPage from "./pages/vehicles";
import VehicleDetailPage from "./pages/vehicleDetail";
import VehicleMaintenanceRecordsPage from "./pages/vehicleMaintenanceRecords";
import VehicleMaintenanceAddRecordPage from "./pages/vehicleMaintenanceAddRecord";

const App = () => {
  useEffect(() => {
    document.title = "Grease Chimp";
  }, []);

  let content = (
    <>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/vehicles" element={<VehiclesPage />} />
        <Route path="/vehicleDetail/:id" Component={VehicleDetailPage} />
        <Route path="/vehicle/:vehicleID/maintenance/records" Component={VehicleMaintenanceRecordsPage} />
        <Route path="/vehicle/:vehicleID/maintenance/addrecord" Component={VehicleMaintenanceAddRecordPage} />
      </Routes>
    </>
  );

  return content;
}

export default App;
