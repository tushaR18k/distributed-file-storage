// app/pages/upload.tsx
"use client";
import Head from 'next/head'
import {useState} from 'react';
import { useRouter } from 'next/navigation';
import { AuthWrapper } from '../components/AuthWrapper';
import Cookies from 'js-cookie';

export default function Upload() {

  const [selectedFiles, setSelectedFiles] = useState<FileList| null>(null);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState('');
  const router = useRouter();

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
      setSelectedFiles(event.target.files);
  };

  const hanndleUpload = async (event: React.FormEvent) => {
    event.preventDefault();

    if(!selectedFiles || selectedFiles.length === 0){
      setMessage("Please select atleast one file to upload!");
      return;
    }

    setUploading(true);
    setMessage('');
    const token = Cookies.get("token");

    const formData = new FormData();
    Array.from(selectedFiles).forEach(file => {
      formData.append('file',file);
    });

    try{
      const response = await fetch(`http://${process.env.NEXT_PUBLIC_AUTH_API_HOST}:${process.env.NEXT_PUBLIC_AUTH_API_PORT}/api/upload`,{
        method: 'POST',
        headers:{
          'Authorization':`Bearer ${token}`
        },
        body: formData,
      });
      if(response.ok){
        setMessage("Files uploaded Successfully!");
        router.push('/files');
      }else if(response.status === 401){
        router.push("/login")
      }else{
        const errorData = await response.json();
        setMessage(`Error: ${errorData.message}`);
        
      }
    }catch(error){
      console.log(error);
      setMessage('An error occurred while uploading the files.');
    }finally{
      setUploading(false);
      setSelectedFiles(null);
    }
  }



  return (
    <AuthWrapper>
      <div className="flex min-h-screen flex-col items-center justify-center py-2">
      <Head>
        <title>Upload Files</title>
        <meta name="description" content="Upload your files to the distributed storage system" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">Upload Files</h1>
        <p className="mt-3 text-2xl">Use the form below to upload your files.</p>

        <form className="mt-6 flex w-full max-w-4xl flex-col items-center justify-around sm:w-full" onSubmit={hanndleUpload}>
          <input type="file" multiple  onChange = {handleFileChange} className="block w-full text-sm text-slate-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-violet-50 file:text-violet-700 hover:file:bg-violet-100"/>
          <button
            type="submit"
            disabled={uploading}
            className={`mt-6 w-32 rounded-xl p-2 text-white ${uploading ? 'bg-gray-500 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700'}`}
          >
            {uploading ? 'Uploading...':'Upload'}
          </button>
        </form>

        {message && (
          <p className='mt-4 text-xl text-red-500'>{message}</p>
        )}

      </main>

      <footer className="flex h-24 w-full items-center justify-center border-t">
        <a
          className="flex items-center justify-center gap-2"
          href="https://github.com/"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          <img src="/vercel.svg" alt="Vercel Logo" className="h-4" />
        </a>
      </footer>
    </div>
    </AuthWrapper>
    
  )
}
