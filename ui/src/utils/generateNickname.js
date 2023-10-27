function GenerateNickname(vehicle) {
  // if the nickname is not provided, set it to 'Year Model'

  if (vehicle) {
    if (vehicle.nickname) {
      return vehicle.nickname;
    } else {
      return `${vehicle.year} ${vehicle.model}`;
    }
  }

  return "";
}

export default GenerateNickname;
