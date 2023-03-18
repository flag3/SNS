import axios from "axios";
import { Link } from "react-router-dom";

function showTweet(props: {
  userID: number;
  username: string;
  displayName: string;
  bio: { String: string; Valid: boolean };
}) {
  return (
    <div>
      <div>名前：{props.displayName}</div>
      <div>
        ユーザー名：
        <Link to={`/users/${props.username}`} key={props.username}>
          @{props.username}
        </Link>
      </div>
      {props.bio.Valid && <div>自己紹介：{props.bio.String}</div>}
      <div>
        <button
          type="submit"
          onClick={() => {
            axios.post(`/api/users/${props.username}/follows`);
          }}
        >
          フォローする
        </button>
        <button
          type="submit"
          onClick={() => {
            axios.delete(`/api/users/${props.username}/follows`);
          }}
        >
          フォロー解除する
        </button>
      </div>
    </div>
  );
}

export default showTweet;
