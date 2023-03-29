import React from "react";
import { NavLink } from "react-router-dom";
import "./sidebar.css";

const SidebarData = [];

const Sidebar = () => {
  return (
    <div className="Sidebar">
      <ul className="SidebarList">
        {SidebarData.map((value, key) => {
          return (
            <NavLink
              to={value.link}
              key={key}
              className={({ isActive, isPending }) =>
                isPending ? "pending" : isActive ? "active" : "row"
              }
            >
              <div id="title">{value.title}</div>
            </NavLink>
          );
        })}
      </ul>
    </div>
  );
};

export default Sidebar;
