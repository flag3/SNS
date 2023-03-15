import { useNavigate } from "react-router-dom";
import "./App.css";

function App() {
  const navigate = useNavigate();

  return (
    <div className="App">
      <h1>Twitter</h1>
      <div className="card">
        <button onClick={() => navigate("/login")}>ログイン</button>
        <button onClick={() => navigate("/signup")}>サインアップ</button>
      </div>
    </div>
  );
}

export default App;
