"use client";
import { useEffect, useState } from "react";
import { useRouter, usePathname } from "next/navigation";
import Cookies from "js-cookie";

interface AuthWrapperProps{
    children: React.ReactNode
}

export const AuthWrapper: React.FC<AuthWrapperProps> = ({children}) => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const router = useRouter();
    const pathName = usePathname();

    useEffect(() => {
        const token = Cookies.get("token");
        if (token) {
          setIsLoggedIn(true);
        } else {
          setIsLoggedIn(false);
          router.push("/login"); // Redirect to login if not logged in
        }
      }, [pathName]);
    
      return isLoggedIn ? <>{children}</> : null;
};