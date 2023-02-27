import axios from "axios";
import { useCallback, useState } from "react";

function Login() {
  const [userid, setUserid] = useState("");
  const [password, setPassword] = useState("");

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios.post("/api/login", {
        userid: userid,
        password: password,
      });
    },
    [userid, password]
  );

  return (
    <div>
      <form>
        <div>
          <label htmlFor="userid">Username: </label>
          <input
            type="text"
            id="userid"
            value={userid}
            onChange={(e) => setUserid(e.target.value)}
          ></input>
        </div>
        <div>
          <label htmlFor="password">Password: </label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          ></input>
        </div>
        <div>
          <button type="submit" onClick={onClickHandler}>
            Login
          </button>
        </div>
      </form>
    </div>
  );
}

export default Login;
