import axios from "axios";
import { Link } from "react-router-dom";

function showTweet(props: {
  userID: number;
  username: string;
  displayName: string;
  bio: { String: string; Valid: boolean };
  location: { String: string; Valid: boolean };
  website: { String: string; Valid: boolean };
  followingCount: number;
  followerCount: number;
  isFollowed: boolean;
  isFollowing: boolean;
}) {
  return (
    <div className="user">
      <div>名前：{props.displayName}</div>
      <div>
        ユーザー名：
        <Link to={`/users/${props.username}`} key={props.username}>
          @{props.username}
        </Link>
      </div>
      <div>フォローされてるか：{props.isFollowed ? 1 : 0}</div>
      <div>フォローしているか：{props.isFollowing ? 1 : 0}</div>
      {props.bio.Valid && <div>自己紹介：{props.bio.String}</div>}
      <div>場所：{props.location.String}</div>
      <div>Web：{props.website.String}</div>
      <div>{props.followerCount} フォロワー</div>
      <div>{props.followingCount} フォロー中</div>
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
