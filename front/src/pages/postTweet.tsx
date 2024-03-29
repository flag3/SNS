import axios from "axios";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";

function Tweet() {
  const navigate = useNavigate();
  const [content, setContent] = useState("");

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios.post("/api/tweets", { content: content }).then(() => {
        navigate("/home");
      });
    },
    [content]
  );

  return (
    <div>
      <form>
        <div>
          <label htmlFor="content">ツイート内容: </label>
          <input
            type="text"
            id="content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
          ></input>
        </div>
        <div>
          <button type="submit" onClick={onClickHandler}>
            ツイートする
          </button>
        </div>
      </form>
    </div>
  );
}

export default Tweet;
