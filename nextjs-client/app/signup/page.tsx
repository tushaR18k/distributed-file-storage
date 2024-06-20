"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import Head from "next/head";
import Link from "next/link";


export default function Signup (){
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const router = useRouter();

    const handleSignup = async (e: React.FormEvent) => {
        e.preventDefault()
        try{
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/signup`,{
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({username,password,email})
            });

            if (res.ok){
                router.push("/login")
            }else{
                const data = await res.json()
                setError(data.error || "An error occurred during signup");
            }
        }catch(err){
            console.log(err)
            setError("An error occurred: ")
        }
    }

     return (
    <div className="flex min-h-screen flex-col items-center justify-center py-2">
      <Head>
        <title>Sign Up</title>
        <meta name="description" content="Sign up for an account" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">Sign Up</h1>
        <p className="mt-3 text-2xl">Create a new account below.</p>

        <form onSubmit={handleSignup} className="mt-6 flex w-full max-w-4xl flex-col items-center justify-around sm:w-full">
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="mt-3 w-full max-w-md p-2 border border-gray-300 rounded text-black"
            required
          />
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="mt-3 w-full max-w-md p-2 border border-gray-300 rounded text-black"
            required
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="mt-3 w-full max-w-md p-2 border border-gray-300 rounded text-black"
            required
          />
          {error && <p className="mt-3 text-red-500">{error}</p>}
          <button type="submit" className="mt-6 w-32 rounded-xl bg-blue-600 p-2 text-white hover:bg-blue-700">
            Sign Up
          </button>
        </form>
      </main>

      <footer className="flex h-24 w-full items-center justify-center border-t">
        <Link href="/" className="flex items-center justify-center gap-2">
          Powered by{' '}
          <img src="/vercel.svg" alt="Vercel Logo" className="h-4" />
        </Link>
      </footer>
    </div>
  );

}