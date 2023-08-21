import './App.css';
import VehicleCard from "./components/VehicleCard"
import Vehicles from './testdata/vehicles'

function App() {
  return (
    <div class="vehicle-card-list">
      {Vehicles.map(v => (
        <VehicleCard vehicle={v} />
      ))} 
    </div>
  );
}

export default App;
