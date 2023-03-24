import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import App from "./App";
import PrivateRoute from "./lib/PrivateRoute";
import Signup from "./pages/signup";
import Login from "./pages/login";
import Logout from "./pages/logout";
import PostTweet from "./pages/postTweet";
import GetTweet from "./pages/getTweet";
import UserLike from "./pages/userLikes";
import LikeUser from "./pages/likeUsers";
import Following from "./pages/following";
import Followers from "./pages/followers";
import UserID from "./pages/userid";
import TweetID from "./pages/tweetid";
import Home from "./pages/home";
import Users from "./pages/users";
import Profile from "./pages/profile";
import Quote from "./pages/quote";
import Retweet from "./pages/retweet";
import Sidebar from "./components/sidebar";
import NonSidebar from "./components/nonsidebar";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
      <NonSidebar />
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/logout" element={<Logout />} />
        <Route
          path="/tweet"
          element={
            <PrivateRoute>
              <Sidebar />
              <PostTweet />
            </PrivateRoute>
          }
        />
        <Route
          path="/tweets"
          element={
            <PrivateRoute>
              <Sidebar />
              <GetTweet />
            </PrivateRoute>
          }
        />
        <Route
          path="/users/:username"
          element={
            <PrivateRoute>
              <Sidebar />
              <UserID />
            </PrivateRoute>
          }
        />
        <Route
          path="/users/:username/likes"
          element={
            <PrivateRoute>
              <Sidebar />
              <UserLike />
            </PrivateRoute>
          }
        />
        <Route
          path="/users/:username/following"
          element={
            <PrivateRoute>
              <Sidebar />
              <Following />
            </PrivateRoute>
          }
        />
        <Route
          path="/users/:username/followers"
          element={
            <PrivateRoute>
              <Sidebar />
              <Followers />
            </PrivateRoute>
          }
        />
        <Route
          path="/home"
          element={
            <PrivateRoute>
              <Sidebar />
              <Home />
            </PrivateRoute>
          }
        />
        <Route
          path="/users"
          element={
            <PrivateRoute>
              <Sidebar />
              <Users />
            </PrivateRoute>
          }
        />
        <Route
          path="/tweets/:tweetID"
          element={
            <PrivateRoute>
              <Sidebar />
              <TweetID />
            </PrivateRoute>
          }
        />
        <Route
          path="/tweets/:tweetID/quotes"
          element={
            <PrivateRoute>
              <Sidebar />
              <Quote />
            </PrivateRoute>
          }
        />
        <Route
          path="/tweets/:tweetID/retweets"
          element={
            <PrivateRoute>
              <Sidebar />
              <Retweet />
            </PrivateRoute>
          }
        />
        <Route
          path="/tweets/:tweetID/likes"
          element={
            <PrivateRoute>
              <Sidebar />
              <LikeUser />
            </PrivateRoute>
          }
        />
        <Route
          path="/profile"
          element={
            <PrivateRoute>
              <Sidebar />
              <Profile />
            </PrivateRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
