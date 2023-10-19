import React from "react"

class AppLogo extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const style = {
      display: "block",
      fontSize: "24pt",
      fontWeight: "bold",
      margin: "20px",
    }

    return (
      <div style={style}>
        Grease Chimp
      </div>
    )
  }
}

export default AppLogo;
