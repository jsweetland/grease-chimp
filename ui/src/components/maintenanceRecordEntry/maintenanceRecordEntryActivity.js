import React from "react"

function MaintenanceRecordEntryActivity(props) {
  const style = {
    display: "table-cell",
    padding: "5px",
  }

  let content = (
    <div style={style}>
      {props.activity}
    </div>
  )

  return content;
}

export default MaintenanceRecordEntryActivity;
