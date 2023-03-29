import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import ShowUserDetail from "./showUserDetail";
import ShowTweet from "./showTweet";

function Tweet() {
  const { username } = useParams();
  const [userList, setUserList] = useState([]);
  const [tweetList, setTweetList] = useState([]);

  useEffect(() => {
    axios.get(`/api/users/${username}`).then((res) => {
      setUserList(res.data);
    });
  }, [username]);

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
      <h2>@{username}</h2>
      {userList.map((user) => {
        return (
          <ShowUserDetail
            key={user.userID}
            userID={user.userID}
            username={user.username}
            displayName={user.displayName}
            bio={user.bio}
            location={user.location}
            website={user.website}
            followingCount={user.followingCount}
            followerCount={user.followerCount}
            isFollowed={user.isFollowed}
            isFollowing={user.isFollowing}
          />
        );
      })}
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
            isRetweeted={tweet.isRetweeted}
            isLiked={tweet.isLiked}
          />
        );
      })}
    </div>
  );
}

export default Tweet;
