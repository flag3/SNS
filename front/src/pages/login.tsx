import axios from "axios";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .post("/api/login", {
          username: username,
          password: password,
        })
        .then(() => {
          navigate("/home");
        })
        .catch((error) => {
          setErrorMessage(error.response.data);
        });
    },
    [username, password]
  );

  return (
    <div>
      <h1>ログイン</h1>
      <form>
        <div>
          <label htmlFor="username">Username: </label>
          <input
            type="text"
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
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
          {errorMessage && <p>{errorMessage}</p>}
        </div>
      </form>
    </div>
  );
}

export default Login;
