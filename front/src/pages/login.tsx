import axios from "axios";
import { useCallback, useState } from "react";

function Login() {
  const [userID, setUserID] = useState("");
  const [password, setPassword] = useState("");

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios.post("/api/login", {
        userID: userID,
        password: password,
      });
    },
    [userID, password]
  );

  return (
    <div>
      <form>
        <div>
          <label htmlFor="userID">Username: </label>
          <input
            type="text"
            id="userID"
            value={userID}
            onChange={(e) => setUserID(e.target.value)}
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
