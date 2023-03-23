import axios from "axios";
import { useCallback, useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import ShowTweetDetail from "./showTweetDetail";

function TweetID() {
  const { tweetID } = useParams();
  const navigate = useNavigate();
  const [content, setContent] = useState("");
  const [tweetList, setTweetList] = useState([]);

  console.log(tweetID);
  console.log(typeof tweetID);
  console.log(Number(tweetID));
  console.log(typeof Number(tweetID));

  const onClickHandler1 = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .post(`/api/tweets/${tweetID}/reply`, { content: content })
        .then(() => {
          navigate("/home");
        });
    },
    [content]
  );

  const onClickHandler2 = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .post(`/api/tweets/${tweetID}/quote`, { content: content })
        .then(() => {
          navigate("/home");
        });
    },
    [content]
  );

  useEffect(() => {
    axios.get(`/api/tweets/${tweetID}`).then((res) => {
      setTweetList(res.data);
    });
  }, [tweetID]);

  useEffect(() => {
    console.log(tweetList);
  }, [tweetList]);

  return (
    <div>
      <h2>@{tweetID}さんのツイート</h2>
      {tweetList.map((tweet) => {
        return (
          <ShowTweetDetail
            key={tweet.tweetID}
            tweetID={tweet.tweetID}
            userID={tweet.userID}
            username={tweet.username}
            displayName={tweet.displayName}
            content={tweet.content}
            reply={tweet.reply}
            quote={tweet.quote}
            replyCount={tweet.replyCount}
            retweetCount={tweet.retweetCount}
            quoteCount={tweet.quoteCount}
            likeCount={tweet.likeCount}
            isRetweeted={tweet.isRetweeted}
            isLiked={tweet.isLiked}
          />
        );
      })}
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
          <button type="submit" onClick={onClickHandler1}>
            リプライ
          </button>
          <button type="submit" onClick={onClickHandler2}>
            引用リツイート
          </button>
        </div>
      </form>
    </div>
  );
}

export default TweetID;
