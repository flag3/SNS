import React from "react";
import { useNavigate, NavLink } from "react-router-dom";
import "./sidebar.css";

const SidebarData = [];

const Sidebar = () => {
  const navigate = useNavigate();
  return (
    <div className="Sidebar">
      <ul className="SidebarList">
        {SidebarData.map((value, key) => {
          return (
            <li
              key={key}
              id={window.location.pathname === value.link ? "active" : ""}
              className="row"
              onClick={() => {
                navigate(value.link);
              }}
            >
              <div id="title">{value.title}</div>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default Sidebar;
