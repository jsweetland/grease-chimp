import React from "react"
import PageHeader from "../components/pageHeader";

class HomePage extends React.Component {
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
          <a href="/vehicles">Vehicle List</a>
        </div>
      </>
    );
  
    return content;
  }
}

export default HomePage;
