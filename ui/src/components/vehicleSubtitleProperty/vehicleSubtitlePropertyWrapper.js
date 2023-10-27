import React from "react";

function VehicleSubtitlePropertyWrapper(props) {
  const style = {
    display: "inline-block",
    fontFace: "Arial,sans",
    fontSize: "9pt",
    padding: "2px 5px",
    margin: "0 10px 0 0",
    border: "1px solid #000",
    borderRadius: "5px",
  };

  let content = (
    <div style={style}>{props.children}</div>
  );

  return content;
}

export default VehicleSubtitlePropertyWrapper;
