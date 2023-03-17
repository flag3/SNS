import axios from "axios";
import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

function Favorite() {
  const { userID } = useParams();
  const [tweetInfo, setTweetInfo] = useState([]);
  const navigate = useNavigate();

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
        (tweet: { tweetID: number; userID: number; content: string }) => {
          return (
            <div key={tweet.tweetID}>
              <br></br>
              <div>番号：{tweet.tweetID}</div>
              <div>ID：{tweet.userID}</div>
              <div>ツイート：{tweet.content}</div>
              <div>
                <button
                  type="button"
                  onClick={() => {
                    navigate(`/tweets/${tweet.tweetID}`);
                  }}
                >
                  リプライする
                </button>
                <button
                  type="button"
                  onClick={() => {
                    axios.post(`/api/tweets/${tweet.tweetID}/retweets`);
                  }}
                >
                  リツイート
                </button>
                <button
                  type="button"
                  onClick={() => {
                    axios.delete(`/api/tweets/${tweet.tweetID}/retweets`);
                  }}
                >
                  リツイート解除
                </button>
                <button
                  type="button"
                  onClick={() => {
                    axios.post(`/api/tweets/${tweet.tweetID}/likes`);
                  }}
                >
                  いいね
                </button>
                <button
                  type="button"
                  onClick={() => {
                    axios.delete(`/api/tweets/${tweet.tweetID}/likes`);
                  }}
                >
                  いいね解除
                </button>
                <button
                  type="button"
                  onClick={() => {
                    let url = "/api/tweet/";
                    url += tweet.tweetID;
                    axios.delete(url);
                  }}
                >
                  消す
                </button>
              </div>
            </div>
          );
        }
      )}
    </div>
  );
}

export default Favorite;
