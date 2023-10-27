import React from "react";
import VehicleSubtitlePropertyValueEmpty from "./vehicleSubtitlePropertyValueEmpty";

function VehicleSubtitlePropertyValue(props) {
  const style = {
    display: "inline-block",
    padding: "0 0 0 5px",
  };

  var content;
  
  if (props.value && props.value != "") {
    content = (
      <div style={style}>
        {props.value}
      </div>
    );
  } else {
    content = (
      <div style={style}>
        <VehicleSubtitlePropertyValueEmpty />
      </div>
    );
  }

  return content;
}

export default VehicleSubtitlePropertyValue;
