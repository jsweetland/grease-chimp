import React from "react"

function PageTitle(props) {
  const style = {
    display: "block",
    fontSize: "16pt",
    fontWeight: "bold",
    margin: "20px 0 10px 0",
  }

  let content = (
    <div style={style}>
      {props.text}
    </div>
  )
  
  return content;
}

export default PageTitle;
