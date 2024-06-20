"use client";
import { ReactNode } from "react";
import NavBar from "./NavBar";

interface LayoutProps{
    children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div>
      <NavBar/>
      <div>{children}</div>
    </div>
  );
};

export default Layout;
