import axios from "axios";
import { useState } from "react";
import { Link } from "react-router-dom";

function showUser(props: {
  userID: number;
  username: string;
  displayName: string;
  bio: { String: string; Valid: boolean };
  isFollowed: boolean;
  isFollowing: boolean;
}) {
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
              axios.delete(`/api/users/${props.username}/follows`);
              setIsFollowing(!isFollowing);
            } else {
              axios.post(`/api/users/${props.username}/follows`);
              setIsFollowing(!isFollowing);
            }
          }}
        >
          {isFollowing ? <div>フォロー中</div> : <div>フォローする</div>}
        </button>
      </div>
      {props.bio.Valid && <div>{props.bio.String}</div>}
    </div>
  );
}

export default showUser;
