import React from "react"
import AppLogo from "./appLogo";

function PageHeader() {
  const style = {
    display: "block",
    background: "green",
  }

  let content = (
    <div style={style}>
      <AppLogo />
    </div>
  )

  return content;
}

export default PageHeader;
