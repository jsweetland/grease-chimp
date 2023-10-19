import React from "react"
import AppLogo from "./appLogo";

class PageHeader extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const style = {
      display: "block",
      background: "green",
    }

    return (
      <div style={style}>
        <AppLogo />
      </div>
    )
  }
}

export default PageHeader;
