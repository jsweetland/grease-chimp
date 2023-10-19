import React from "react"
import PageHeader from "../components/pageHeader";

function HomePage() {
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

export default HomePage;
