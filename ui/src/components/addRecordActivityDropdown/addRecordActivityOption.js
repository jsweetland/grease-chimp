import React from "react";

function AddRecordActivityOption(props) {
  let content = (
    <option value={props.id}>{props.activity}</option>
  );

  return content;
}

export default AddRecordActivityOption;
