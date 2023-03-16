import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Following() {
  const { id } = useParams();
  const [accountInfo, setAccountInfo] = useState([]);

  useEffect(() => {
    axios.get(`/api/users`).then((res) => {
      setAccountInfo(res.data);
    });
  }, [id]);

  useEffect(() => {
    console.log(accountInfo);
  }, [accountInfo]);

  return (
    <div>
      <h2>ユーザーリスト</h2>
      {accountInfo.map(
        (account: {
          id: number;
          username: string;
          displayName: string;
          bio: { String: string; Valid: boolean };
        }) => {
          return (
            <div key={account.id}>
              <br></br>
              <div>ID：{account.id}</div>
              <div>ユーザー名：{account.username}</div>
              <div>表示名：{account.displayName}</div>
              {account.bio.Valid && <div>自己紹介：{account.bio.String}</div>}
            </div>
          );
        }
      )}
    </div>
  );
}

export default Following;
