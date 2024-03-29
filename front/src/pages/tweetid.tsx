import axios from "axios";
import { useCallback, useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import ShowTweetDetail from "./showTweetDetail";

function TweetID() {
  const { tweetID } = useParams();
  const navigate = useNavigate();
  const [content, setContent] = useState("");
  const [tweetList, setTweetList] = useState([]);
  const [replyList, setReplyList] = useState([]);

  const onClickReplyHandler = useCallback(
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

  const onClickQuoteHandler = useCallback(
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

  useEffect(() => {
    axios.get(`/api/tweets/${tweetID}/reply`).then((res) => {
      setReplyList(res.data);
    });
  }, [tweetID]);

  useEffect(() => {
    console.log(replyList);
  }, [replyList]);

  return (
    <div>
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
          <button type="submit" onClick={onClickReplyHandler}>
            リプライ
          </button>
          <button type="submit" onClick={onClickQuoteHandler}>
            引用リツイート
          </button>
        </div>
      </form>
      {replyList.map((tweet) => {
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
    </div>
  );
}

export default TweetID;
