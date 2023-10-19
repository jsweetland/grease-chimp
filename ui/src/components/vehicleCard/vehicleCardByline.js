import React from "react"

class VehicleCardByline extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const style = {
      display: "float",
      float: "left",
      clear: "both",
      padding: "5px 0 0 0",
      fontSize: "11pt",
      fontWeight: "normal",
    }

    return (
      <div style={style}>
        {this.props.byline}
      </div>
    )
  }
}

export default VehicleCardByline;
