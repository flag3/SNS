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
            followeeuserid: userID,
          });
        }}
      >
        フォローする
      </button>
      <button
        type="submit"
        onClick={() => {
          let url = "/api/delete/follow/";
          url += userID;
          axios.delete(url);
        }}
      >
        フォロー解除する
      </button>
      {tweetInfo.map(
        (tweet: { tweetid: number; userid: string; body: string }) => {
          return (
            <div key={tweet.tweetid}>
              <br></br>
              <div>番号：{tweet.tweetid}</div>
              <div>ID：{tweet.userid}</div>
              <div>
                ツイート：{tweet.body}
                <button
                  type="submit"
                  onClick={() => {
                    axios.post("/api/like", {
                      tweetid: tweet.tweetid,
                    });
                  }}
                >
                  いいね
                </button>
                <button
                  type="submit"
                  onClick={() => {
                    let url = "/api/delete/like/";
                    url += tweet.tweetid;
                    axios.delete(url);
                  }}
                >
                  いいね解除
                </button>
                <button
                  type="submit"
                  onClick={() => {
                    let url = "/api/delete/tweet/";
                    url += tweet.tweetid;
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
