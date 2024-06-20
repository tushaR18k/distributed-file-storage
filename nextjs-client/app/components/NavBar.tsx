"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import Cookies from "js-cookie"; // Ensure js-cookie is installed
import { useRouter, usePathname } from "next/navigation";


export default function NavBar() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const router = useRouter(); // Initialize useRouter hook
  const pathName = usePathname();

  useEffect(() => {
    const token = Cookies.get("token");
    if (token) {
      setIsLoggedIn(true);
    } else {
      setIsLoggedIn(false);
    }
  }, [pathName]);

  const handleLogout = () => {
    // Remove token from cookies on logout
    Cookies.remove("token");
    setIsLoggedIn(false);
    
    // Redirect to home page after logout
    router.push("/");
  };

  return (
    <nav className="flex items-center justify-between flex-wrap bg-blue-500 p-6">
      <div className="flex items-center flex-shrink-0 text-white mr-6">
        <span className="font-semibold text-xl tracking-tight">My App</span>
      </div>
      <div className="flex">
        <div className="text-sm mr-4">
          <Link href="/home" className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-gray-200">
              Home
          </Link>
        </div>
        <div className="text-sm mr-4">
          <Link href="/files" className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-gray-200">
              Files
          </Link>
        </div>
        <div className="text-sm mr-4">
          <Link href="/upload" className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-gray-200">
              Upload
          </Link>
        </div>
        <div>
          {!isLoggedIn ? (
            <Link href="/login" className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-gray-200">
                Login
            </Link>
          ) : (
            <button
              onClick={handleLogout}
              className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-gray-200"
            >
              Logout
            </button>
          )}
        </div>
      </div>
    </nav>
  );
}
