import React from "react";
import { useNavigate } from "react-router-dom";
import "./sidebar.css";

const SidebarData = [];

const Sidebar = () => {
  return (
    <div className="Sidebar">
      <ul className="SidebarList">
        {SidebarData.map((value, key) => {
          const navigate = useNavigate();
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
