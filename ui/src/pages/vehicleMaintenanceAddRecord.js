import React, {useState, useEffect, useCallback} from "react"
import PageHeader from "../components/pageHeader";
import GetMaintenanceActivities from "../api/getMaintenanceActivities"; 
import {useParams} from "react-router-dom";
import AddRecordActivityDropdown from "../components/addRecordActivityDropdown";

function VehicleMaintenanceAddRecordPage() {
  let {vehicleID} = useParams();
  let [maintenanceActivities, updateMaintenanceActivities] = useState();

  const fetchMaintenanceActivityData = useCallback(async () => {
    try {
      let getMaintenanceActivitiesResponse = await GetMaintenanceActivities();
      updateMaintenanceActivities(getMaintenanceActivitiesResponse);
    } catch (err) {
      console.log(err);
    }
  }, [])

  useEffect(() => {
    fetchMaintenanceActivityData();
  }, [fetchMaintenanceActivityData])

  const style = {};

  let content = (
    <>
      <PageHeader />
      <form>
        <AddRecordActivityDropdown />
      </form>
    </>
  );

  return content;
}

export default VehicleMaintenanceAddRecordPage;
