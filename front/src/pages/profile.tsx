import axios from "axios";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";

function Profile() {
  const navigate = useNavigate();
  const [displayName, setDisplayName] = useState("");
  const [bio, setBio] = useState("");
  const [location, setLocation] = useState("");
  const [website, setWebsite] = useState("");

  const onClickHandler1 = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .put("/api/profile/userDisplayName", { displayName: displayName })
        .then(() => {
          navigate("/home");
        });
    },
    [displayName]
  );

  const onClickHandler2 = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .put("/api/profile/userBio", { bio: { String: bio, Valid: true } })
        .then(() => {
          navigate("/home");
        });
    },
    [bio]
  );

  const onClickHandler3 = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .put("/api/profile/userLocation", {
          location: { String: location, Valid: true },
        })
        .then(() => {
          navigate("/home");
        });
    },
    [location]
  );

  const onClickHandler4 = useCallback(
    (e) => {
      e.preventDefault();
      axios
        .put("/api/profile/userWebsite", {
          website: { String: website, Valid: true },
        })
        .then(() => {
          navigate("/home");
        });
    },
    [website]
  );

  return (
    <div>
      <form>
        <div>
          <label htmlFor="content">名前の変更</label>
          <input
            type="text"
            id="content"
            value={displayName}
            onChange={(e) => setDisplayName(e.target.value)}
          ></input>
        </div>
        <div>
          <button type="submit" onClick={onClickHandler1}>
            変更する
          </button>
        </div>
      </form>
      <form>
        <div>
          <label htmlFor="content">Bioの変更</label>
          <input
            type="text"
            id="content"
            value={bio}
            onChange={(e) => setBio(e.target.value)}
          ></input>
        </div>
        <div>
          <button type="submit" onClick={onClickHandler2}>
            変更する
          </button>
        </div>
      </form>
      <form>
        <div>
          <label htmlFor="content">場所の変更</label>
          <input
            type="text"
            id="content"
            value={location}
            onChange={(e) => setLocation(e.target.value)}
          ></input>
        </div>
        <div>
          <button type="submit" onClick={onClickHandler3}>
            変更する
          </button>
        </div>
      </form>
      <form>
        <div>
          <label htmlFor="content">Webの変更</label>
          <input
            type="text"
            id="content"
            value={website}
            onChange={(e) => setWebsite(e.target.value)}
          ></input>
        </div>
        <div>
          <button type="submit" onClick={onClickHandler4}>
            変更する
          </button>
        </div>
      </form>
    </div>
  );
}

export default Profile;
