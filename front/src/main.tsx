import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import App from "./App";
import PrivateRoute from "./lib/PrivateRoute";
import Signup from "./pages/signup";
import Login from "./pages/login";
import Logout from "./pages/logout";
import Tweet from "./pages/tweet";
import Favorite from "./pages/favorite";
import Following from "./pages/following";
import Followers from "./pages/followers";
import UserID from "./pages/userid";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/logout" element={<Logout />} />
        <Route
          path="/tweet"
          element={
            <PrivateRoute>
              <Tweet />
            </PrivateRoute>
          }
        />
        <Route
          path="/:userID"
          element={
            <PrivateRoute>
              <UserID />
            </PrivateRoute>
          }
        />
        <Route
          path="/:userID/likes"
          element={
            <PrivateRoute>
              <Favorite />
            </PrivateRoute>
          }
        />
        <Route
          path="/:userID/following"
          element={
            <PrivateRoute>
              <Following />
            </PrivateRoute>
          }
        />
        <Route
          path="/:userID/followers"
          element={
            <PrivateRoute>
              <Followers />
            </PrivateRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
