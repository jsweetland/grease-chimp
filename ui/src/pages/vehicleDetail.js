import React from "react"
import PageHeader from "../components/pageHeader";
import {useParams} from "react-router-dom";

function VehicleDetailPage() {
  const style = {
    background: "white",
  };

  const {id} = useParams()
  console.log("id = ", JSON.stringify(id))
  
  let content = (
    <>
      <PageHeader />
      <div style={style}>
        <div>Detailed Information for ID {id}</div>
      </div>
    </>
  );

  return content;
}

export default VehicleDetailPage;
