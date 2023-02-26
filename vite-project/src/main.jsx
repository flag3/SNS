import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import App from "./App";
import PrivateRoute from "./lib/PrivateRoute";
import Signup from "./pages/signup";
import Login from "./pages/login";
import Tweet from "./pages/tweet";
import UserID from "./pages/userid";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
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
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
