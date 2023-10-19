import React from "react"
import PageHeader from "../components/pageHeader";
import {useParams} from "react-router-dom";

function VehicleDetailPage() {
  const style = {
    background: "white",
  };

  const {vin} = useParams()
  console.log("vin = ", JSON.stringify(vin))
  // console.log("VIN=", this.props.match.params.vin)
  
  let content = (
    <>
      <PageHeader />
      <div style={style}>
        <div>Detailed Information for VIN {vin}</div>
      </div>
    </>
  );

  return content;
}

export default VehicleDetailPage;
