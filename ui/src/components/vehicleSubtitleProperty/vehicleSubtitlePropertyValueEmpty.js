import React from "react";

function VehicleSubtitlePropertyValueEmpty() {
  const style = {
    display: "inline-block",
    color: "#808080",
    fontStyle: "italic",
  };

  let content = (
    <div style={style}>Not specified</div>
  );

  return content;
}

export default VehicleSubtitlePropertyValueEmpty;
