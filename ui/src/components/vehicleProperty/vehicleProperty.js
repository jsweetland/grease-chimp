import React from "react";
import VehiclePropertyLabel from "./vehiclePropertyLabel";
import VehiclePropertyValue from "./vehiclePropertyValue";
import VehiclePropertyWrapper from "./vehiclePropertyWrapper";

function VehicleProperty(props) {
  let content = (
    <VehiclePropertyWrapper>
      <VehiclePropertyLabel text={props.label} />
      <VehiclePropertyValue value={props.value} />
    </VehiclePropertyWrapper>
  );

  return content;
}

export default VehicleProperty;
