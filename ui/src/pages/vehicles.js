import React from "react"
import VehicleCardList from "../components/vehicleCardList";
import PageHeader from "../components/pageHeader";

class VehiclesPage extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const style = {
      background: "white",
    };
  
    let content = (
      <>
        <PageHeader />
        <div style={style}>
          <VehicleCardList />
        </div>
      </>
    );
  
    return content;
  }
}

export default VehiclesPage;
