import React from "react";
import VehicleCardNickname from "./vehicleCardNickname";
import VehicleCardByline from "./vehicleCardByline";
import VehicleCardDetailLink from "./vehicleCardDetailLink";
import GenerateByline from "../../utils/generateByline";
import GenerateNickname from "../../utils/generateNickname";
import VehicleCardMaintenanceRecordsLink from "./vehicleCardMaintenanceRecordsLink";

function VehicleCard(props) {
  // set the byline
  const byline = GenerateByline(props.vehicle);

  // set the nickname
  const nickname = GenerateNickname(props.vehicle);

  const style = {
    display: "inline-block",
    width: "300px",
    height: "100px",
    borderRadius: "10px",
    padding: "10px",
    margin: "20px 0 20px 20px",
    overflow: "auto",
    backgroundColor: "lightgray",
  };

  let content = (
    <div style={style} key={props.vehicle.id}>
      <VehicleCardNickname nickname={nickname} />
      <VehicleCardByline byline={byline} />
      <VehicleCardDetailLink id={props.vehicle.id} />
      <VehicleCardMaintenanceRecordsLink id={props.vehicle.id} />
    </div>
  );

  return content;
}

export default VehicleCard;
