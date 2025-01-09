import React from "react"

function MaintenanceRecordEntryMileage(props) {
  const style = {
    display: "table-cell",
    padding: "5px",
  }

  let content = (
    <div style={style}>
      {props.mileage}
    </div>
  )

  return content;
}

export default MaintenanceRecordEntryMileage;
