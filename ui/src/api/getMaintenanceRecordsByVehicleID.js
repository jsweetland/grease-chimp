import APIClient from "./apiClient"

const GetMaintenanceRecordsByVehicleID = async (vehicleID) => {
  const maintenanceRecordsResponse = await APIClient.get(`/maintenance/records/${vehicleID}`)
  const maintenanceRecords = await maintenanceRecordsResponse.data;

  return maintenanceRecords;
}

export default GetMaintenanceRecordsByVehicleID;