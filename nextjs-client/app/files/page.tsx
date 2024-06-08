// app/pages/files.tsx
"use client";
import Head from 'next/head'
import Link from 'next/link'
import { useState, useEffect } from 'react';
import { fetchFiles, File } from '../lib/api';



export default function Files() {
  const [files, setFiles]  = useState<File[]>([]);

  useEffect(()=>{
    async function getFiles(){
      try{
        const filesData = await fetchFiles();
        setFiles(filesData);
      }catch(error){
        console.log('Error fetching files: ', error);
      }
    }
    getFiles();
  },[]);

  return (
    <div className="flex min-h-screen flex-col items-center justify-center py-2">
      <Head>
        <title>Manage Files</title>
        <meta name="description" content="View and manage your uploaded files" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">Manage Files</h1>
        <p className="mt-3 text-2xl">View and manage your uploaded files below.</p>

        {files.length > 0 ?(
          <ul className="mt-6 w-full max-w-4xl list-disc text-left sm:w-full">
            {files.map((file:File,index:number)=> (
              <li key={index} className='mt-2 text-xl'>
                <Link href={`${process.env.NEXT_PUBLIC_API_URL}/api/download/${file.name}`}
                className='text-blue-500 hover:underline'>
                  {file.name}
                </Link>
              </li>
            ))}
          </ul>
        ):(
          <p className="mt-6 text-xl text-red-500">You haven't uploaded any files, please upload some.</p>
        )}

        <Link href="/upload" className="mt-6 inline-block rounded bg-blue-500 px-6 py-3 text-white font-bold text-lg transition duration-300 ease-in-out hover:bg-blue-700 hover:shadow-lg">
          Upload More files!
        </Link>
      </main>

      <footer className="flex h-24 w-full items-center justify-center border-t">
        <Link href="/" className="flex items-center justify-center gap-2">
            Powered by{' '}
            <img src="/vercel.svg" alt="Vercel Logo" className="h-4" />
        </Link>
      </footer>
    </div>
  )
}
