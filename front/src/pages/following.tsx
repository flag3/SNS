import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import ShowUser from "./showUser";

function Following() {
  const { username } = useParams();
  const [userList, setUserList] = useState([]);

  useEffect(() => {
    axios.get(`/api/users/${username}/following`).then((res) => {
      setUserList(res.data);
    });
  }, [username]);

  useEffect(() => {
    console.log(userList);
  }, [userList]);

  return (
    <div>
      <h2>{username}がフォロー中</h2>
      {userList.map((user) => {
        return (
          <ShowUser
            key={user.userID}
            userID={user.userID}
            username={user.username}
            displayName={user.displayName}
            bio={user.bio}
            isFollowed={user.isFollowed}
            isFollowing={user.isFollowing}
          />
        );
      })}
    </div>
  );
}

export default Following;
