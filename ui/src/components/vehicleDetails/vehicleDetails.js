import GenerateNickname from "../../utils/generateNickname";
import PageTitle from "../pageTitle";
import VehicleDetailsSubtitle from "./vehicleDetailsSubtitle";
import VehicleProperty from "../vehicleProperty/vehicleProperty";
import VehiclePropertyTwoValues from "../vehicleProperty/vehiclePropertyTwoValues";

function VehicleDetails(props) {
  const style = {
    margin: "20px",
  };

  let content = (
    <div style={style}>
      <PageTitle text={GenerateNickname(props.vehicle)} />
      <VehicleDetailsSubtitle
        id={props.vehicle ? props.vehicle.id : ""}
        vin={props.vehicle ? props.vehicle.vin : ""}
        plate={props.vehicle ? props.vehicle.plate : ""}
      />
      <div>
        <VehicleProperty label="Year" value={props.vehicle ? props.vehicle.year : "-"} />
        <VehicleProperty label="Make" value={props.vehicle ? props.vehicle.make : "-"} />
        <VehicleProperty label="Model" value={props.vehicle ? props.vehicle.model : "-"} />
        <VehicleProperty label="Trim" value={props.vehicle ? props.vehicle.trim : "-"} />
        <VehicleProperty label="Package" value={props.vehicle ? props.vehicle.package : "-"} />
        <VehicleProperty label="Nickname" value={props.vehicle ? props.vehicle.nickname : "-"} />
        <VehiclePropertyTwoValues
          label="Color"
          primary={props.vehicle ? props.vehicle.color.name : ""}
          secondary={props.vehicle ? props.vehicle.color.hex : ""}
        />
      </div>
    </div>
  );

  return content;
}

export default VehicleDetails;
