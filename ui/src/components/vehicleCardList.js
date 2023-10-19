import React from "react"
import GetVehicles from "../api/getVehicles";
import VehicleCard from "./vehicleCard/vehicleCard";

class VehicleCardList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      vehicles: [],
    }
  }

  componentDidMount() {
    this.renderVehicleCardList();
  }

  renderVehicleCardList = async () => {
    try {
      const vehicles = await GetVehicles();

      this.setState({
        vehicles: vehicles,
      });
    } catch (err) {
      console.log(err);
    }
  }

  render() {
    const style = {};

    console.log("Vehicles:")
    console.log(this.state.vehicles)

    const vehicles = this.state.vehicles?.map((vehicle) => (
      <VehicleCard vehicle={vehicle} />
    ));

    return (
      <div style={style}>
        {vehicles}
      </div>
    );
  }
}

export default VehicleCardList;
