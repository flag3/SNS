import axios from "axios";
import { useEffect, useState } from "react";

function Users() {
  const [userList, setUserList] = useState([]);

  useEffect(() => {
    axios.get(`/api/users`).then((res) => {
      setUserList(res.data);
    });
  }, []);

  useEffect(() => {
    console.log(userList);
  }, [userList]);

  return (
    <div>
      <h2>ユーザーリスト</h2>
      {userList.map(
        (user: {
          id: number;
          username: string;
          displayName: string;
          bio: { String: string; Valid: boolean };
        }) => {
          return (
            <div key={user.id}>
              <br></br>
              <div>名前：{user.displayName}</div>
              <div>ユーザー名：@{user.username}</div>
              {user.bio.Valid && <div>自己紹介：{user.bio.String}</div>}
              <button
                type="submit"
                onClick={() => {
                  axios.post(`/api/users/${user.username}/follows`);
                }}
              >
                フォローする
              </button>
              <button
                type="submit"
                onClick={() => {
                  axios.delete(`/api/users/${user.username}/follows`);
                }}
              >
                フォロー解除する
              </button>
            </div>
          );
        }
      )}
    </div>
  );
}

export default Users;
