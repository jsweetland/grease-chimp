import React, {useState, useEffect, useCallback} from "react"
import GetVehicles from "../api/getVehicles";
import VehicleCard from "./vehicleCard";

function VehicleCardList() {
  const [vehicles, updateVehicles] = useState([])

  const fetchVehicleData = useCallback(async () => {
    try {
      let getVehiclesResponse = await GetVehicles()
      updateVehicles(getVehiclesResponse)
    } catch (err) {
      console.log(err);
    }
  }, [])

  useEffect(() => {
    fetchVehicleData()
  }, [fetchVehicleData])

  const style = {};

  const vehicleList = vehicles?.map((vehicle) => (
    <VehicleCard vehicle={vehicle} />
  ));

  let content = (
    <div style={style}>
      {vehicleList}
    </div>
  );

  return content;
}

export default VehicleCardList;
