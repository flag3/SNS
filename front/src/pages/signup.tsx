import axios from "axios";
import { useCallback, useState } from "react";

function Signup() {
  const [userid, setUserid] = useState("");
  const [password, setPassword] = useState("");

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios.post("/api/signup", {
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
            Signup
          </button>
        </div>
      </form>
    </div>
  );
}

export default Signup;
