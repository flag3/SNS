import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

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
      <h2>{username}のツイート</h2>
      <button
        type="submit"
        onClick={() => {
          axios.post(`/api/users/${username}/follows`);
        }}
      >
        フォローする
      </button>
      <button
        type="submit"
        onClick={() => {
          axios.delete(`/api/users/${username}/follows`);
        }}
      >
        フォロー解除する
      </button>
      {tweetList.map(
        (tweet: {
          tweetID: number;
          userID: number;
          content: string;
          replay: number;
          quote: number;
        }) => {
          return (
            <div key={tweet.tweetID}>
              <br></br>
              <div>番号：{tweet.tweetID}</div>
              <div>ID：{tweet.userID}</div>
              <div>
                ツイート：{tweet.content}
                <button
                  type="submit"
                  onClick={() => {
                    axios.post("/api/tweets/" + tweet.tweetID + "/likes");
                  }}
                >
                  いいね
                </button>
                <button
                  type="submit"
                  onClick={() => {
                    axios.delete("/api/tweets/" + tweet.tweetID + "/likes");
                  }}
                >
                  いいね解除
                </button>
                <button
                  type="submit"
                  onClick={() => {
                    axios.delete("/api/tweets/" + tweet.tweetID);
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

export default Tweet;
