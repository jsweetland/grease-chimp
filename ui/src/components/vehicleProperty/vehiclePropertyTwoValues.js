import React from "react";
import VehiclePropertyLabel from "./vehiclePropertyLabel";
import VehiclePropertyValue from "./vehiclePropertyValue";

function VehiclePropertyTwoValues(props) {
  const style = {
    display: "block",
    fontFace: "Arial,sans",
    fontSize: "12pt",
  };

  let content = (
    <div style={style}>
      <VehiclePropertyLabel text={props.label} />
      <VehiclePropertyValue value={
        props.primary
          ? (props.secondary
            ? `${props.primary} (${props.secondary})`
            : "")
          : props.primary}
      />
    </div>
  );

  return content;
}

export default VehiclePropertyTwoValues;
