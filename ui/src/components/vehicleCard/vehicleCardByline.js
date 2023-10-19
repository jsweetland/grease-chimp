import React from "react"

function VehicleCardByline(props) {
  const style = {
    display: "float",
    float: "left",
    clear: "both",
    padding: "5px 0 0 0",
    fontSize: "11pt",
    fontWeight: "normal",
  }

  let content = (
    <div style={style}>
      {props.byline}
    </div>
  )

  return content;
}

export default VehicleCardByline;
