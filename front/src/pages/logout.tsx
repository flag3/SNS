import axios from "axios";
import { useNavigate } from "react-router-dom";

function Logout() {
  axios.get("/api/logout");
  const navigate = useNavigate();

  return (
    <div>
      <h2>ログアウトしました</h2>
      <button
        type="button"
        onClick={() => {
          navigate("/login");
        }}
      >
        ログインする
      </button>
    </div>
  );
}

export default Logout;
