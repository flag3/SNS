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
  isRetweeted: boolean;
  isLiked: boolean;
}) {
  const navigate = useNavigate();
  return (
    <div className="tweet">
      <div>ツイートID：{props.tweetID}</div>
      <div>名前：{props.displayName}</div>
      <div>
        ユーザー名：
        <Link to={`/users/${props.username}`} key={props.username}>
          @{props.username}
        </Link>
      </div>
      <div>ツイート：{props.content}</div>
      {props.reply.Valid && <div>リプライ{props.reply.Int64}</div>}
      {props.quote.Valid && <div>引用{props.quote.Int64}</div>}
      <div>リプライ数：{props.replyCount}</div>
      <div>リツイート数：{props.retweetCount}</div>
      <div>引用数：{props.quoteCount}</div>
      <div>いいね数：{props.likeCount}</div>
      <div>リツイート済み：{props.isRetweeted ? 1 : 0}</div>
      <div>いいね済み：{props.isLiked ? 1 : 0}</div>
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
