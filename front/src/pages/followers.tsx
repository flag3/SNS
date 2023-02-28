import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Following() {
  const { userID } = useParams();
  const [accountInfo, setAccountInfo] = useState([]);

  useEffect(() => {
    axios.get(`/api/${userID}/followers`).then((res) => {
      setAccountInfo(res.data);
    });
  }, [userID]);

  useEffect(() => {
    console.log(accountInfo);
  }, [accountInfo]);

  return (
    <div>
      <h2>{userID}さんのフォロワー</h2>
      {accountInfo.map((account: { userid: string; username: string }) => {
        return (
          <div key={account.userid}>
            <br></br>
            <div>ユーザーID：{account.userid}</div>
            <div>表示名：{account.username}</div>
          </div>
        );
      })}
    </div>
  );
}

export default Following;
