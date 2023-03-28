import axios from "axios";
import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import ReplyIcon from "@mui/icons-material/Reply";
import RepeatIcon from "@mui/icons-material/Repeat";
import RepeatOnIcon from "@mui/icons-material/RepeatOn";
import FavoriteIcon from "@mui/icons-material/Favorite";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";

type Tweet = {
  tweetID: number;
  userID: number;
  username: string;
  displayName: string;
  content: string;
  reply: { Int64: number; Valid: boolean };
  quote: { Int64: number; Valid: boolean };
  replyCount: number;
  retweetCount: number;
  quoteCount: number;
  likeCount: number;
  isRetweeted: boolean;
  isLiked: boolean;
};

function showTweet(props: Tweet) {
  const navigate = useNavigate();
  const [isRetweeted, setIsRetweeted] = useState(props.isRetweeted);
  const [retweetCount, setRetweetCount] = useState(props.retweetCount);
  const [isLiked, setIsLiked] = useState(props.isLiked);
  const [likeCount, setLikeCount] = useState(props.likeCount);
  return (
    <div className="tweet">
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
      <div>{props.content}</div>
      <div>
        <button
          type="button"
          onClick={() => {
            navigate(`/tweets/${props.tweetID}`);
          }}
          key={props.tweetID}
        >
          <ReplyIcon />
          {props.replyCount}
        </button>
        <button
          type="button"
          onClick={() => {
            if (isRetweeted) {
              axios.delete(`/api/tweets/${props.tweetID}/retweets`).then(() => {
                setIsRetweeted(!isRetweeted);
                setRetweetCount(retweetCount - 1);
              });
            } else {
              axios.post(`/api/tweets/${props.tweetID}/retweets`).then(() => {
                setIsRetweeted(!isRetweeted);
                setRetweetCount(retweetCount + 1);
              });
            }
          }}
        >
          {isRetweeted ? <RepeatOnIcon /> : <RepeatIcon />}
          {retweetCount + props.quoteCount}
        </button>
        <button
          type="button"
          onClick={() => {
            if (isLiked) {
              axios.delete(`/api/tweets/${props.tweetID}/likes`).then(() => {
                setIsLiked(!isLiked);
                setLikeCount(likeCount - 1);
              });
            } else {
              axios.post(`/api/tweets/${props.tweetID}/likes`).then(() => {
                setIsLiked(!isLiked);
                setLikeCount(likeCount + 1);
              });
            }
          }}
        >
          {isLiked ? <FavoriteIcon /> : <FavoriteBorderIcon />}
          {likeCount}
        </button>
      </div>
    </div>
  );
}

export default showTweet;
