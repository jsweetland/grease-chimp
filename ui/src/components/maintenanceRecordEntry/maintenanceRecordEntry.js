import React from "react";
import MaintenanceRecordEntryDate from "./maintenanceRecordEntryDate";
import MaintenanceRecordEntryActivity from "./maintenanceRecordEntryActivity";
import MaintenanceRecordEntryMileage from "./maintenanceRecordEntryMileage";
// import VehicleCardNickname from "./vehicleCardNickname";
// import VehicleCardByline from "./vehicleCardByline";
// import VehicleCardDetailLink from "./vehicleCardDetailLink";
// import GenerateByline from "../../utils/generateByline";
// import GenerateNickname from "../../utils/generateNickname";
// import VehicleCardMaintenanceRecordsLink from "./vehicleCardMaintenanceRecordsLink";

function MaintenanceRecordEntry(props) {
  let style = {}

  if (props.isHeader) {
    style = {
      display: "table-row",
      overflow: "auto",
      background: "lightgray",
      fontWeight: "bold",
      fontSize: "11pt",
    };
  } else {
    style = {
      display: "table-row",
      overflow: "auto",
      fontSize: "11pt",
    };
  }

  let content = (
    <div style={style} key={props.record.id}>
      <MaintenanceRecordEntryDate date={props.record.dateperformed} />
      <MaintenanceRecordEntryActivity activity={props.record.activity} />
      <MaintenanceRecordEntryMileage mileage={props.record.mileage} />
    </div>
  );

  return content;
}

export default MaintenanceRecordEntry;
