import axios from "axios";
import { useEffect, useState } from "react";
import ShowTweet from "./showTweet";

function Tweets() {
  const [tweetList, setTweetList] = useState([]);

  useEffect(() => {
    axios.get(`/api/tweets`).then((res) => {
      setTweetList(res.data);
    });
  }, []);

  useEffect(() => {
    console.log(tweetList);
  }, [tweetList]);

  return (
    <div>
      <h2>ツイート一覧</h2>
      {tweetList.map((tweet) => {
        return (
          <ShowTweet
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

export default Tweets;
