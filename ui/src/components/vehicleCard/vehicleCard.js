import React from "react"
import VehicleCardNickname from "./vehicleCardNickname";
import VehicleCardByline from "./vehicleCardByline";
import VehicleCardDetailLink from "./vehicleCardDetailLink";

class VehicleCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      byline: "",
      nickname: "",
    }
  }

  componentDidMount() {
    this.setByline();
    this.setNickname();
  }

  setByline = () => {
    // the byline should be 'Year Make Model Trim (Package)', Trim and Package could be empty

    let trimValue = this.props.vehicle.trim ? ` ${this.props.vehicle.trim}` : ""
    let packageValue = this.props.vehicle.package ? ` (${this.props.vehicle.package} Package)` : ""

    this.setState({
      byline: `${this.props.vehicle.year} ${this.props.vehicle.make} ${this.props.vehicle.model} ${trimValue}${packageValue}`
    })
  }

  setNickname = () => {
    // if the nickname is not provided, set it to 'Year Model'

    this.setState({
      nickname: this.props.vehicle.nickname ? this.props.vehicle.nickname : `${this.props.vehicle.year} ${this.props.vehicle.model}`
    })
  }

  render() {
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

    return (
      <div style={style} key={this.props.vehicle.vin}>
        <VehicleCardNickname nickname={this.state.nickname} />
        <VehicleCardByline byline={this.state.byline} />
        <VehicleCardDetailLink vin={this.props.vehicle.vin} />
      </div>
    );
  }
}

export default VehicleCard;
