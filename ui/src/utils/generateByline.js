function GenerateByline(vehicle) {
  // the byline should be 'Year Make Model Trim (Package)', Trim and Package could be empty
  
  let trimValue = vehicle.trim ? ` ${vehicle.trim}` : "";
  let packageValue = vehicle.package ? ` (${vehicle.package} Package)` : "";

  const byline = `${vehicle.year} ${vehicle.make} ${vehicle.model} ${trimValue}${packageValue}`;
  
  return byline;
}

export default GenerateByline;
