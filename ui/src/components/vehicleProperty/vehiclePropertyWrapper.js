import React from "react";

function VehiclePropertyWrapper(props) {
  const style = {
    display: "block",
    fontFace: "Arial,sans",
    fontSize: "12pt",
  };

  let content = (
    <div style={style}>{props.children}</div>
  );

  return content;
}

export default VehiclePropertyWrapper;
