import axios from "axios";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import ShowTweetDetail from "./showTweetDetail";

function Quote() {
  const { tweetID } = useParams();
  const [tweetList, setTweetList] = useState([]);

  useEffect(() => {
    axios.get(`/api/tweets/${tweetID}/quote`).then((res) => {
      setTweetList(res.data);
    });
  }, [tweetID]);

  useEffect(() => {
    console.log(tweetList);
  }, [tweetList]);

  return (
    <div>
      <h2>引用ツイート</h2>
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
    </div>
  );
}

export default Quote;
