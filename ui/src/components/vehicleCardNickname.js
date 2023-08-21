import React from "react"

class VehicleCardNickname extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const style = {
      display: "float",
      float: "left",
      clear: "both",
      fontSize: "16pt",
      fontWeight: "bold",
    }

    return (
      <div style={style}>
        {this.props.nickname}
      </div>
    )
  }
}

export default VehicleCardNickname;
