import React from "react"

function VehicleCardDetailLink(props) {
  const style = {
    display: "float",
    float: "left",
    clear: "both",
    padding: "5px 0 0 0",
    fontSize: "10pt",
    fontWeight: "normal",
  }

  let vehicleDetailRoute = `/vehicleDetail/${props.vin}`

  let content = (
    <div style={style}>
      <a href={vehicleDetailRoute}>View Vehicle Detail</a>
    </div>
  )

  return content;
}

export default VehicleCardDetailLink;
