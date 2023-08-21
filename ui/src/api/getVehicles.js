import APIClient from "./apiClient"

const GetVehicles = async () => {
  const vehiclesResponse = await APIClient.get("/vehicles")
  const vehicles = await vehiclesResponse.data;

  return vehicles;
}

export default GetVehicles;