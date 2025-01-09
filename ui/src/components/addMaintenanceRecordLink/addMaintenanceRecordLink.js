import React from "react";

function AddMaintenanceRecordLink(props) {
  const style = {
    fontSize: "10pt",
    margin: "20px",
  };

  let addRecordRoute = `/vehicle/${props.vehicleID}/maintenance/addrecord`

  let content = (
    <div style={style}>
      <a href={addRecordRoute}>+ Record Maintenance</a>
    </div>
  );

  return content;
}

export default AddMaintenanceRecordLink;
