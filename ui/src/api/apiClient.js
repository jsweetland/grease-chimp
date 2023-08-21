import axios from "axios";

export const APIClient = axios.create({
  baseURL: "http://localhost:10000",
});

export default APIClient;