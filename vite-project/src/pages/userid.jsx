import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Tweet() {
  const { userID } = useParams();
  const [tweetInfo, setTweetInfo] = useState([]);

  useEffect(() => {
    axios.get(`/api/${userID}`).then((res) => {
      setTweetInfo(res.data);
    });
  }, [userID]);

  useEffect(() => {
    console.log(tweetInfo);
  }, [tweetInfo]);

  return (
    <div>
      <h2>{userID}のツイート</h2>
      {tweetInfo.map((tweet) => {
        return (
          <div key={tweet.tweetid}>
            <br></br>
            <div>番号：{tweet.id}</div>
            <div>ID：{tweet.userid}</div>
            <div>ツイート：{tweet.body}</div>
          </div>
        );
      })}
    </div>
  );
}

export default Tweet;
