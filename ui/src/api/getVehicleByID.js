import APIClient from "./apiClient"

const GetVehicleByID = async (id) => {
  const vehicleResponse = await APIClient.get(`/vehicle/${id}`)
  const vehicle = await vehicleResponse.data;

  return vehicle;
}

export default GetVehicleByID;