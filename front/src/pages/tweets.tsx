import axios from "axios";
import { useEffect, useState } from "react";

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
      {tweetList.map(
        (tweet: { tweetID: number; userID: string; content: string }) => {
          return (
            <div key={tweet.tweetID}>
              <br></br>
              <div>番号：{tweet.tweetID}</div>
              <div>ID：{tweet.userID}</div>
              <div>ツイート：{tweet.content}</div>
              <div>
                <button
                  type="submit"
                  onClick={() => {
                    axios.post("/api/like", {
                      tweetID: tweet.tweetID,
                    });
                  }}
                >
                  いいね
                </button>
                <button
                  type="submit"
                  onClick={() => {
                    let url = "/api/like/";
                    url += tweet.tweetID;
                    axios.delete(url);
                  }}
                >
                  いいね解除
                </button>
                <button
                  type="submit"
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

export default Tweets;
