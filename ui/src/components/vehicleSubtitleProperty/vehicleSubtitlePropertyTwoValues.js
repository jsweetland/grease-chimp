import React from "react";
import VehicleSubtitlePropertyLabel from "./vehicleSubtitlePropertyLabel";
import VehicleSubtitlePropertyValue from "./vehicleSubtitlePropertyValue";
import VehicleSubtitlePropertyWrapper from "./vehicleSubtitlePropertyWrapper";

function VehicleSubtitlePropertyTwoValues(props) {
  let content = (
    <VehicleSubtitlePropertyWrapper>
      <VehicleSubtitlePropertyLabel text={props.label} />
      <VehicleSubtitlePropertyValue value={
        props.primary
          ? (props.secondary
            ? `${props.primary} (${props.secondary})`
            : "")
          : props.primary}
      />
    </VehicleSubtitlePropertyWrapper>
  );

  return content;
}

export default VehicleSubtitlePropertyTwoValues;
