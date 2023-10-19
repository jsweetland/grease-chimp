import React from "react"

class VehicleCardDetailLink extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const style = {
      display: "float",
      float: "left",
      clear: "both",
      padding: "5px 0 0 0",
      fontSize: "10pt",
      fontWeight: "normal",
    }

    let vehicleDetailRoute = `/vehicleDetail/${this.props.vin}`

    let content = (
      <div style={style}>
        <a href={vehicleDetailRoute}>View Vehicle Detail</a>
      </div>
    )

    return content;
  }
}

export default VehicleCardDetailLink;
