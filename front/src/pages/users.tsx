import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Following() {
  const { userID } = useParams();
  const [accountInfo, setAccountInfo] = useState([]);

  useEffect(() => {
    axios.get(`/api/users`).then((res) => {
      setAccountInfo(res.data);
    });
  }, [userID]);

  useEffect(() => {
    console.log(accountInfo);
  }, [accountInfo]);

  return (
    <div>
      <h2>ユーザーリスト</h2>
      {accountInfo.map((account: { userID: string; username: string }) => {
        return (
          <div key={account.userID}>
            <br></br>
            <div>ユーザーID：{account.userID}</div>
            <div>表示名：{account.username}</div>
          </div>
        );
      })}
    </div>
  );
}

export default Following;
