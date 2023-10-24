import React, { useState, useEffect} from "react";
import VehicleCardNickname from "./vehicleCardNickname";
import VehicleCardByline from "./vehicleCardByline";
import VehicleCardDetailLink from "./vehicleCardDetailLink";

function VehicleCard(props) {
  // set the byline
  // the byline should be 'Year Make Model Trim (Package)', Trim and Package could be empty
  let trimValue = props.vehicle.trim ? ` ${props.vehicle.trim}` : ""
  let packageValue = props.vehicle.package ? ` (${props.vehicle.package} Package)` : ""
  const byline = `${props.vehicle.year} ${props.vehicle.make} ${props.vehicle.model} ${trimValue}${packageValue}`

  // set the nickname
  // if the nickname is not provided, set it to 'Year Model'
  const nickname = props.vehicle.nickname ? props.vehicle.nickname : `${props.vehicle.year} ${props.vehicle.model}`

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
    </div>
  );

  return content;
}

export default VehicleCard;
