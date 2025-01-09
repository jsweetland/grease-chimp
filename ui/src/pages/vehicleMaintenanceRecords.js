import React, {useState, useEffect, useCallback} from "react"
import PageHeader from "../components/pageHeader";
import GetMaintenanceRecordsByVehicleID from "../api/getMaintenanceRecordsByVehicleID";
import {useParams} from "react-router-dom";
import MaintenanceRecordEntry from "../components/maintenanceRecordEntry";
import AddMaintenanceRecordLink from "../components/addMaintenanceRecordLink";

function VehicleMaintenanceRecordsPage() {
  let {vehicleID} = useParams();
  let [maintenanceRecords, updateMaintenanceRecords] = useState();

  const fetchMaintenanceRecordData = useCallback(async (vehicleID) => {
    try {
      let getMaintenanceRecordsResponse = await GetMaintenanceRecordsByVehicleID(vehicleID);
      updateMaintenanceRecords(getMaintenanceRecordsResponse);
    } catch (err) {
      console.log(err);
    }
  }, [])

  useEffect(() => {
    fetchMaintenanceRecordData(vehicleID);
  }, [fetchMaintenanceRecordData, vehicleID])

  const tableStyle = {
    display: "table",
    background: "white",
    margin: "20px",
  };

  const tableHeaderData = {
    dateperformed: "Date",
    activity: "Activity",
    mileage: "Mileage",
  }

  const noRecordsFoundStyle = {
    display: "block",
    fontSize: "11pt",
    fontStyle: "italic",
    padding: "5px",
  }

  const maintenanceRecordList = maintenanceRecords?.map((record) => (
    <MaintenanceRecordEntry record={record} />
  ));

  const listElement = maintenanceRecordList ? 
      maintenanceRecordList : 
      (<div style={noRecordsFoundStyle}>No records found.</div>);

  let content = (
    <>
      <PageHeader />
      <AddMaintenanceRecordLink vehicleID={vehicleID} />
      <div style={tableStyle}>
        <MaintenanceRecordEntry record={tableHeaderData} isHeader="true" />
        {listElement}
      </div>
    </>
  );

  return content;
}

export default VehicleMaintenanceRecordsPage;
