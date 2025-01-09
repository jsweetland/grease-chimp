import React from "react"

function VehicleCardMaintenanceRecordsLink(props) {
  const style = {
    display: "float",
    float: "left",
    clear: "both",
    padding: "5px 0 0 0",
    fontSize: "10pt",
    fontWeight: "normal",
  }

  let vehicleMaintenanceRecordsRoute = `/vehicle/${props.id}/maintenance/records`

  let content = (
    <div style={style}>
      <a href={vehicleMaintenanceRecordsRoute}>Maintenance Records</a>
    </div>
  )

  return content;
}

export default VehicleCardMaintenanceRecordsLink;
