import { Link } from "react-router-dom";
import "./App.css";

function App() {
  return (
    <div className="App">
      <h1>Twitter</h1>
      <div className="card">
        <Link to={"/login"}>
          <button>ログイン</button>
        </Link>
        <Link to={"/signup"}>
          <button>サインアップ</button>
        </Link>
      </div>
    </div>
  );
}

export default App;
