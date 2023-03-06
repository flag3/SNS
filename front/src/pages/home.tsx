import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Favorite() {
  const { userID } = useParams();
  const [tweetInfo, setTweetInfo] = useState([]);

  useEffect(() => {
    axios.get(`/api/home`).then((res) => {
      setTweetInfo(res.data);
    });
  }, [userID]);

  useEffect(() => {
    console.log(tweetInfo);
  }, [tweetInfo]);

  return (
    <div>
      <h2>ホーム</h2>
      {tweetInfo.map(
        (tweet: { tweetid: number; userid: string; body: string }) => {
          return (
            <div key={tweet.tweetid}>
              <br></br>
              <div>番号：{tweet.tweetid}</div>
              <div>ID：{tweet.userid}</div>
              <div>ツイート：{tweet.body}</div>
            </div>
          );
        }
      )}
    </div>
  );
}

export default Favorite;
