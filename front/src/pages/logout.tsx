import axios from "axios";
import { Link } from "react-router-dom";

function Logout() {
  axios.get("/api/logout");

  return (
    <div>
      <h2>ログアウトしました</h2>
      <Link to="/">
        <button type="button">Twitter</button>
      </Link>
    </div>
  );
}

export default Logout;
