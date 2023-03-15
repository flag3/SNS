import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Favorite() {
  const { userID } = useParams();
  const [tweetInfo, setTweetInfo] = useState([]);

  useEffect(() => {
    axios.get(`/api/${userID}/likes`).then((res) => {
      setTweetInfo(res.data);
    });
  }, [userID]);

  useEffect(() => {
    console.log(tweetInfo);
  }, [tweetInfo]);

  return (
    <div>
      <h2>{userID}さんがいいねしました</h2>
      {tweetInfo.map(
        (tweet: { tweetID: number; userID: string; content: string }) => {
          return (
            <div key={tweet.tweetID}>
              <br></br>
              <div>番号：{tweet.tweetID}</div>
              <div>ID：{tweet.userID}</div>
              <div>ツイート：{tweet.content}</div>
            </div>
          );
        }
      )}
    </div>
  );
}

export default Favorite;
