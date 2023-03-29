import axios from "axios";
import { useState } from "react";
import { Link } from "react-router-dom";

type UserDetail = {
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
};

function showUserDetail(props: UserDetail) {
  const [isFollowing, setIsFollowing] = useState(props.isFollowing);
  return (
    <div className="user">
      <div>
        <Link to={`/users/${props.username}`} key={props.username}>
          {props.displayName}
        </Link>
      </div>
      <div>
        <Link to={`/users/${props.username}`} key={props.username}>
          @{props.username}
        </Link>
      </div>
      {props.isFollowed && <div>フォローされています</div>}
      <div>
        <button
          type="submit"
          onClick={() => {
            if (isFollowing) {
              axios.delete(`/api/users/${props.username}/follows`).then(() => {
                setIsFollowing(!isFollowing);
              });
            } else {
              axios.post(`/api/users/${props.username}/follows`).then(() => {
                setIsFollowing(!isFollowing);
              });
            }
          }}
        >
          {isFollowing ? <div>フォロー中</div> : <div>フォローする</div>}
        </button>
      </div>
      {props.bio.Valid && <div>{props.bio.String}</div>}
      <div>場所：{props.location.String}</div>
      <div>
        Web：<a href={props.website.String}>{props.website.String}</a>
      </div>
      <div>
        <Link to={`/users/${props.username}/followers`} key="followers">
          {props.followerCount} フォロワー
        </Link>
      </div>
      <div>
        <Link to={`/users/${props.username}/following`} key="following">
          {props.followingCount} フォロー中
        </Link>
      </div>
      <div>
        <Link to={`/users/${props.username}/likes`} key="likes">
          @{props.username}のいいねを表示
        </Link>
      </div>
    </div>
  );
}

export default showUserDetail;
