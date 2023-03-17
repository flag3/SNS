import axios from "axios";
import { useNavigate } from "react-router-dom";

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
  likeCount: number;
}) {
  const navigate = useNavigate();
  return (
    <div>
      <div>番号：{props.tweetID}</div>
      <div>ID：{props.userID}</div>
      <div>名前：{props.displayName}</div>
      <div>ユーザー名：@{props.username}</div>
      <div>ツイート：{props.content}</div>
      <div>リプライ数：{props.replyCount}</div>
      <div>リツイート数：{props.retweetCount}</div>
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
            axios.post(`/api/tweets/${props.tweetID}/retweet`);
          }}
        >
          リツイート
        </button>
        <button
          type="button"
          onClick={() => {
            axios.delete(`/api/tweets/${props.tweetID}/retweet`);
          }}
        >
          リツイート解除
        </button>
        <button
          type="button"
          onClick={() => {
            axios.post(`/api/tweets/${props.tweetID}/like`);
          }}
        >
          いいね
        </button>
        <button
          type="button"
          onClick={() => {
            axios.delete(`/api/tweets/${props.tweetID}/like`);
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
