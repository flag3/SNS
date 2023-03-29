import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import ShowUser from "./showUser";

function User() {
  const { tweetID } = useParams();
  const [userList, setUserList] = useState([]);

  useEffect(() => {
    axios.get(`/api/tweets/${tweetID}/likes`).then((res) => {
      setUserList(res.data);
    });
  }, []);

  useEffect(() => {
    console.log(userList);
  }, [userList]);

  console.log(userList);

  return (
    <div>
      <h2>いいねしたユーザー</h2>
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

export default User;
