import axios from "axios";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";

function Login() {
  const [userID, setUserID] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .post("/api/login", {
          userID: userID,
          password: password,
        })
        .then(() => {
          navigate("/home");
        })
        .catch((error) => {
          setErrorMessage(error.response.data);
        });
    },
    [userID, password]
  );

  return (
    <div>
      <h1>ログイン</h1>
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
          {
            errorMessage && <p>{errorMessage}</p>
            // errorMessage が空でない場合 <p>{errorMessage}</p> を返す
          }
        </div>
      </form>
    </div>
  );
}

export default Login;
