import React from "react";
import VehicleSubtitleProperty from "../vehicleSubtitleProperty";
import VehicleSubtitlePropertyTwoValues from "../vehicleSubtitleProperty/vehicleSubtitlePropertyTwoValues";

function VehicleDetailsSubtitle(props) {
  const style = {
    margin: "0 0 15px 0",
  };

  let content = (
    <div style={style}>
      <VehicleSubtitleProperty label="ID" value={props.id} />
      <VehicleSubtitleProperty label="VIN" value={props.vin} />
      <VehicleSubtitlePropertyTwoValues
        label="License Plate"
        primary={props.plate ? props.plate.value : ""}
        secondary={props.plate ? props.plate.issuer : ""}
      />
    </div>
  );

  return content;
}

export default VehicleDetailsSubtitle;
