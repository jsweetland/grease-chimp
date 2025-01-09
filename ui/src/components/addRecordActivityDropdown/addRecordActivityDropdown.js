import React, {useState, useEffect, useCallback} from "react"
import GetMaintenanceActivities from "../../api/getMaintenanceActivities";
import AddRecordActivityOption from "./addRecordActivityOption";

function AddRecordActivityDropdown(props) {
  const [activities, updateActivities] = useState([])

  const fetchActivityData = useCallback(async () => {
    try {
      let getActivitiesResponse = await GetMaintenanceActivities()
      updateActivities(getActivitiesResponse)
    } catch (err) {
      console.log(err);
    }
  }, [])

  useEffect(() => {
    fetchActivityData()
  }, [fetchActivityData])

  const style = {};

  const activityOptions = activities?.map((activity) => (
    <AddRecordActivityOption id={activity.id} activity={activity.activity} />
  ));

  let content = (
    <select id="activity" name="activity">
      {activityOptions}
    </select>
  );

  return content;
};

export default AddRecordActivityDropdown;
