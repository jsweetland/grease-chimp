import APIClient from "./apiClient"

const GetMaintenanceActivities = async () => {
  const maintenanceActivitiesResponse = await APIClient.get(`/maintenance/activities`)
  const maintenanceActivities = await maintenanceActivitiesResponse.data;

  return maintenanceActivities;
}

export default GetMaintenanceActivities;