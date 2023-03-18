import axios from "axios";
import { useNavigate, Link } from "react-router-dom";

function showTweet(props: {
  tweetID: number;
  userID: number;
  username: string;
  displayName: string;
  content: string;
  reply: { Int64: number; Valid: boolean };
  quote: { Int64: number; Valid: boolean };
  replyCount: number;
  retweetCount: number;
  quoteCount: number;
  likeCount: number;
}) {
  const navigate = useNavigate();
  return (
    <div>
      <div>名前：{props.displayName}</div>
      <div>
        ユーザー名：
        <Link to={`/users/${props.username}`} key={props.username}>
          @{props.username}
        </Link>
      </div>
      <div>ツイート：{props.content}</div>
      <div>リプライ数：{props.replyCount}</div>
      <div>リツイート数：{props.retweetCount}</div>
      <div>引用数：{props.quoteCount}</div>
      <div>いいね数：{props.likeCount}</div>
      <div>
        <button
          type="button"
          onClick={() => {
            navigate(`/tweets/${props.tweetID}`);
          }}
        >
          返信
        </button>
        <button
          type="button"
          onClick={() => {
            axios.post(`/api/tweets/${props.tweetID}/retweets`);
          }}
        >
          リツイート
        </button>
        <button
          type="button"
          onClick={() => {
            navigate(`/tweets/${props.tweetID}`);
          }}
        >
          引用リツイート
        </button>
        <button
          type="button"
          onClick={() => {
            axios.delete(`/api/tweets/${props.tweetID}/retweets`);
          }}
        >
          リツイート解除
        </button>
        <button
          type="button"
          onClick={() => {
            axios.post(`/api/tweets/${props.tweetID}/likes`);
          }}
        >
          いいね
        </button>
        <button
          type="button"
          onClick={() => {
            axios.delete(`/api/tweets/${props.tweetID}/likes`);
          }}
        >
          いいね解除
        </button>
        <button
          type="button"
          onClick={() => {
            axios.delete(`/api/tweets/${props.tweetID}`);
          }}
        >
          消す
        </button>
      </div>
    </div>
  );
}

export default showTweet;
