import axios from "axios";
import { useCallback, useState } from "react";

function Tweet() {
  //const [userid, setUserid] = useState("");
  const [body, setBody] = useState("");

  const onClickHandler = useCallback(
    (e) => {
      e.preventDefault();
      axios.post("/api/tweet", {
        body: body,
      });
    },
    [body]
  );

  return (
    <div>
      <form>
        <div>
          <label htmlFor="body">ツイート内容: </label>
          <input
            type="text"
            id="body"
            value={body}
            onChange={(e) => setBody(e.target.value)}
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
