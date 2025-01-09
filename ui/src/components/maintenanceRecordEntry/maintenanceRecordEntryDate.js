import React from "react"

function MaintenanceRecordEntryDate(props) {
  const style = {
    display: "table-cell",
    padding: "5px",
  }

  let content = (
    <div style={style}>
      {props.date}
    </div>
  )

  return content;
}

export default MaintenanceRecordEntryDate;
