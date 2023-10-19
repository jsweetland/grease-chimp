import React from "react"
import PageHeader from "../components/pageHeader";

class VehicleDetailPage extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    console.log(JSON.stringify(this.props))

    const style = {
      background: "white",
    };

    console.log("this.props = ", JSON.stringify(this.props))
    // console.log("VIN=", this.props.match.params.vin)
  
    let content = (
      <>
        <PageHeader />
        <div style={style}>
          {/* <div>Detailed Information for VIN {this.props.match.params.vin}</div> */}
        </div>
      </>
    );
  
    return content;
  }
}

export default VehicleDetailPage;
