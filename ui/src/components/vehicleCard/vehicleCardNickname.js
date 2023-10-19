import React from "react"

function VehicleCardNickname(props) {
  const style = {
    display: "float",
    float: "left",
    clear: "both",
    fontSize: "16pt",
    fontWeight: "bold",
  }

  let content = (
    <div style={style}>
      {props.nickname}
    </div>
  )

  return content;
}

export default VehicleCardNickname;
