import React from "react";

function VehicleSubtitlePropertyLabel(props) {
  const style = {
    display: "inline-block",
  };

  let content = (
    <div style={style}>
      {props.text}:
    </div>
  );

  return content;
}

export default VehicleSubtitlePropertyLabel;
