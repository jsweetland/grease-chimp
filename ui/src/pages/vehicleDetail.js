import React, {useState, useEffect, useCallback} from "react"
import PageHeader from "../components/pageHeader";
import GetVehicleByID from "../api/getVehicleByID";
import VehicleDetails from "../components/vehicleDetails";
import {useParams} from "react-router-dom";

function VehicleDetailPage() {
  let {id} = useParams();
  let [vehicle, updateVehicle] = useState();

  const fetchVehicleData = useCallback(async (id) => {
    try {
      let getVehicleResponse = await GetVehicleByID(id);
      updateVehicle(getVehicleResponse);
    } catch (err) {
      console.log(err);
    }
  }, [])

  useEffect(() => {
    fetchVehicleData(id);
  }, [fetchVehicleData, id])

  const style = {
    background: "white",
  };
  
  let content = (
    <>
      <PageHeader />
      <div style={style}>
        <VehicleDetails vehicle={vehicle} />
      </div>
    </>
  );

  return content;
}

export default VehicleDetailPage;
