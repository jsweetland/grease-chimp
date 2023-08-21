import VehicleCardList from "./components/vehicleCardList";

const App = () => {
  const style = {
    background: "white",
  };

  let content = (
    <div style={style}>
      <VehicleCardList />
    </div>
  );

  return content;
}

export default App;
