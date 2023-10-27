import React from "react";
import VehicleSubtitlePropertyLabel from "./vehicleSubtitlePropertyLabel";
import VehicleSubtitlePropertyValue from "./vehicleSubtitlePropertyValue";
import VehicleSubtitlePropertyWrapper from "./vehicleSubtitlePropertyWrapper";

function VehicleSubtitleProperty(props) {
  let content = (
    <VehicleSubtitlePropertyWrapper>
      <VehicleSubtitlePropertyLabel text={props.label} />
      <VehicleSubtitlePropertyValue value={props.value} />
    </VehicleSubtitlePropertyWrapper>
  );

  return content;
}

export default VehicleSubtitleProperty;
