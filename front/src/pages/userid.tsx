import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import ShowTweet from "./showTweet";

function Tweet() {
  const { username } = useParams();
  const [tweetList, setTweetList] = useState([]);

  useEffect(() => {
    axios.get(`/api/users/${username}/tweets`).then((res) => {
      setTweetList(res.data);
    });
  }, [username]);

  useEffect(() => {
    console.log(tweetList);
  }, [tweetList]);

  return (
    <div>
      <h2>@{username}さんのツイート</h2>
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
          />
        );
      })}
    </div>
  );
}

export default Tweet;
