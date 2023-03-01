import axios from "axios";

function Logout() {
  axios.get("/api/logout");
  return <h2>ログアウトしました</h2>;
}

export default Logout;
