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
      <button
        type="submit"
        onClick={() => {
          axios.post("/api/follow", {
            followeeUserID: userID,
          });
        }}
      >
        フォローする
      </button>
      <button
        type="submit"
        onClick={() => {
          let url = "/api/follow/";
          url += userID;
          axios.delete(url);
        }}
      >
        フォロー解除する
      </button>
      {tweetInfo.map(
        (tweet: { tweetID: number; userID: string; body: string }) => {
          return (
            <div key={tweet.tweetID}>
              <br></br>
              <div>番号：{tweet.tweetID}</div>
              <div>ID：{tweet.userID}</div>
              <div>
                ツイート：{tweet.body}
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

export default Tweet;
