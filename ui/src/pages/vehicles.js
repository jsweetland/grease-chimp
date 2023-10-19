import React from "react"
import VehicleCardList from "../components/vehicleCardList";
import PageHeader from "../components/pageHeader";

function VehiclesPage() {
  const style = {
    background: "white",
  };

  let content = (
    <>
      <PageHeader />
      <div style={style}>
        <VehicleCardList />
      </div>
    </>
  );

  return content;
}

export default VehiclesPage;
