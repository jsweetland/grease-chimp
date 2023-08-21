// This import will be used for when the values are coming in from state
// import { useState } from 'react';

export function VehicleCard(props) {
  let trimValue = props.vehicle.trim ? ` ${props.vehicle.trim}` : ""
  let packageValue = props.vehicle.package ? ` (${props.vehicle.package} Package)` : ""

  // the byline should be 'Year Make Model Trim (Package)', Trim and Package could be empty
  let byline = `${props.vehicle.year} ${props.vehicle.make} ${props.vehicle.model} ${trimValue}${packageValue}`

  // if the nickname is not provided, set it to 'Year Model'
  let nickname = props.vehicle.nickname ? props.vehicle.nickname : `${props.vehicle.year} ${props.vehicle.model}`

  return (
    <div class="vehicle-card" key={props.vehicle.vin}>
      <div class="vehicle-card-nickname">{nickname}</div>
      <div class="vehicle-card-byline">{byline}</div>
    </div>
  )
}

export default VehicleCard;
