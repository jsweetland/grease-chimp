import React from "react";
import VehiclePropertyValueEmpty from "./vehiclePropertyValueEmpty";

function VehiclePropertyValue(props) {
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
        <VehiclePropertyValueEmpty />
      </div>
    );
  }

  return content;
}

export default VehiclePropertyValue;
